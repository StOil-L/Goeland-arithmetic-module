/******************/
/* simplex_rules.go */
/******************/
/**
* This file provides rules to manage simplex rules
**/

//TODO
//remainders

package ari

import (
	"errors"
	"fmt"
	"math/big"

	typing "goeland-arith/ArithmTests/polymorphism"
	treetypes "goeland-arith/ArithmTests/tree-types"
	"goeland-arith/ArithmTests/types"
)

type pair_coef_var struct {
	coef     *big.Rat
	variable types.Meta
}

/* Take a list of formulas, return the fomulas which are a predicate with at least one metavariable */
func keepRelevantPred(fl types.FormList) []types.Pred {
	relevant_pred_list := []types.Pred{}
	for _, f := range fl {
		if pred, ok := f.(types.Pred); ok {
			if checkBinaryArithmeticPredicate(pred.GetID()) {
				if checkPredContainsMetavariable(pred) {
					relevant_pred_list = append(relevant_pred_list, pred)
				}
			}
		}
	}
	return relevant_pred_list
}

/* check if a binary predicate contain at least on metavariable (recursively) */
func checkPredContainsMetavariable(p types.Pred) bool {
	arg1 := p.GetArgs()[0]
	arg2 := p.GetArgs()[1]
	return len(arg1.GetMetas()) > 0 || len(arg2.GetMetas()) > 0
}

/* create a substition from the simplex's result */
/**
* TODO : que renvoie votre simplexe ?
* On va dire que le simplex me renvoie un map [string]*Big.rat.
* Dans l'idéal il aurait fallu avoir une structure commune int et rat
**/
func buildSubstitutionFromSimplexResult(res_simplex map[string]*big.Rat, map_mv_sv map[string]types.Meta) treetypes.Substitutions {
	subst_res := treetypes.MakeEmptySubstitution()
	for key, value := range res_simplex {
		if typing.IsInt(map_mv_sv[key].GetTypeHint()) && !value.IsInt() {
			fmt.Printf("Error in BuildSubstitutionFromSimplexResult : value of %v should be an int but is %v", map_mv_sv[key].ToString(), value.String())
			return treetypes.Failure()
		} else {
			subst_res[map_mv_sv[key]] = types.MakerFun(types.MakerId(value.String()), []types.Term{})
		}
	}
	return subst_res

}

/*** TODO : cette fonction est juste là pour que ça compile ***/
func simplex(s info_system) (info_system, bool) {
	return s, false
}

/* Normalize equation for simplex call */
/**
* TODO :
* Le type de retour doit être celui que prend votre simplexe en entrée
* Je suppose que vous allez parser les prédicat et les transforer en >= ou <=
* Je vous mets le schéma de base, à vous de l'adapter et de savoir quoi renvoyer !
*
* On va renvoyer 3 choses :
* la strcuture/liste pour votre simplexe
* une map de correspondence entre les metavariables que vous allez utiliser et celles que j'ai moi. On va supposer que vos variables sont des strings
* Une liste des variables qui doivent être entières
*
* En théorie dans la liste des prédicat il n'y aura que des <, >, <=, >=, mais dans le doute on va quand-même traiter l'égalité
* Ca va beaucoup ressembler aux autre fonctions de parse que vous avez, sauf que là au lieu que les deux côtés soient des ints ou des rats, vous allez avoir des métavariables à convertir
* Les deux listes qui concernent les metavariables sont à utiliser comme des pointeurs pour qu'elles puissent se construire tout au long du parsing
**/

func normalizeForSimplex(pl []types.Pred) (info_system, map[string]types.Meta) {
	map_variable_metavariables := make(map[string]types.Meta)
	int_variables := []string{}
	var tab_variable = make([]types.Meta, 0)
	ligne_matrice := 0
	var list_list_pcv = make([][]pair_coef_var, 0)
	for _, p := range pl {
		var list_pcv = make([]pair_coef_var, 0)

		t1, val1, err1 := termToSimplex(p.GetArgs()[0], &map_variable_metavariables, &int_variables, true, list_pcv)
		t2, val2, err2 := termToSimplex(p.GetArgs()[1], &map_variable_metavariables, &int_variables, false, list_pcv)
		if err1 != nil || err2 != nil {
			fmt.Printf("Error in normalizeForSimplex")
			var mauvais_system info_system
			return mauvais_system, nil
		}

		switch p.GetID().GetName() {

		case types.Id_eq.GetName():
			tab_variable = passe1(p, t1, t2, &ligne_matrice, tab_variable, true)
			list_list_pcv = passe2(list_list_pcv, list_pcv, 2, val1, val2)

		case types.Id_neq.GetName():

		//à réfléchir, si on a 2x != 3 alors on a 2x > 3  OU  2x < 3
		//il faudrait donc faire 2 systèmes, un avec chacune des équations..
		//et encore, ça ne marche que si on cherche une solution entière..

		case "less":

			tab_variable = passe1(p, t1, t2, &ligne_matrice, tab_variable, false)

		case "lesseq":
			tab_variable = passe1(p, t1, t2, &ligne_matrice, tab_variable, false)
			list_list_pcv = passe2(list_list_pcv, list_pcv, 1, val1, val2)

		case "great":
			tab_variable = passe1(p, t1, t2, &ligne_matrice, tab_variable, false)

		case "greateq":
			tab_variable = passe1(p, t1, t2, &ligne_matrice, tab_variable, false)
			list_list_pcv = passe2(list_list_pcv, list_pcv, 0, val1, val2)
		}

	}
	for _, tab := range list_list_pcv {
		for _, ma_struct := range tab {
			fmt.Print(ma_struct.coef, " ", ma_struct.variable.GetName(), "  ---  ")
		}
	}

	fmt.Println()
	fmt.Println()
	fmt.Println()
	list_list_pcv_sort := passe3(list_list_pcv, tab_variable)

	for _, tab := range list_list_pcv_sort {
		for _, ma_struct := range tab {
			fmt.Print(ma_struct.coef, " ", ma_struct.variable.GetName(), "  ---  ")
		}
	}

	fmt.Println()
	//ici le system
	system := constrSystem(list_list_pcv_sort, int_variables)
	return system, map_variable_metavariables
}

func constrSystem(list_list_pcv_sort [][]pair_coef_var, int_variables []string) info_system {
	var system info_system
	var matrice_coef = make([][]*big.Rat, 0)
	var tab_constr = make([]*big.Rat, 0)
	var tab_var = make([]string, 0)
	for i := 0; i < len(list_list_pcv_sort); i++ {
		var ligne_matrice = make([]*big.Rat, 0)
		for j := 0; j < len(list_list_pcv_sort[i]); j++ {
			if j == len(list_list_pcv_sort[i])-1 {
				tab_constr = append(tab_constr, list_list_pcv_sort[i][j].coef)
				matrice_coef = append(matrice_coef, ligne_matrice)
			} else {
				ligne_matrice = append(ligne_matrice, list_list_pcv_sort[i][j].coef)
				if i == 0 {
					tab_var = append(tab_var, list_list_pcv_sort[i][j].variable.ToString())
				}
			}
		}
	}

	system.tab_coef = matrice_coef
	system.tab_cont = tab_constr
	system.tab_nom_var = tab_var

	pos_var_tab := create_pos_var_tab(system.tab_coef, system.tab_nom_var)
	system.pos_var_tab = pos_var_tab
	bland := bland(system.pos_var_tab, system.tab_coef, system.tab_nom_var)
	system.bland = bland
	alpha_tab := create_alpha_tab(system.tab_coef, system.tab_nom_var)
	system.alpha_tab = alpha_tab
	var tab_int_bool = make([]bool, 0)
	for i := 0; i < len(system.tab_nom_var); i++ {
		present := false
		for j := 0; j < len(int_variables); j++ {
			if int_variables[j] == system.tab_nom_var[i] {
				present = true
			}
		}
		if present {
			tab_int_bool = append(tab_int_bool, true)
		} else {
			tab_int_bool = append(tab_int_bool, false)
		}
	}
	system.tab_int_bool = tab_int_bool

	return system
}

func passe3(list_list_pcv [][]pair_coef_var, tab_variable []types.Meta) [][]pair_coef_var {

	var list_list_pcv_sort = make([][]pair_coef_var, 0)
	var meta_const types.Meta
	for i := 0; i < len(list_list_pcv); i++ {
		var list_pcv_sort = make([]pair_coef_var, 0)
		cpt_var := 0
		var pair pair_coef_var
		pair.coef = newRat()
		for j := 0; j < len(list_list_pcv[i]); j++ {
			again := false
			if list_list_pcv[i][j].variable == tab_variable[cpt_var] {
				for k := 0; k < len(list_pcv_sort); k++ {
					if list_list_pcv[i][j].variable == list_pcv_sort[k].variable {
						list_pcv_sort[k].coef.Add(list_pcv_sort[k].coef, list_list_pcv[i][j].coef)
						again = true

					}
				}
				if !again {
					var pair pair_coef_var
					pair.coef = new(big.Rat).Set(list_list_pcv[i][j].coef)
					pair.variable = list_list_pcv[i][j].variable
					list_pcv_sort = append(list_pcv_sort, pair)
					if cpt_var < len(tab_variable)-1 {
						cpt_var += 1
					}
				} else {
					again = false
				}
			} else {
				for k := 0; k < len(list_pcv_sort); k++ {
					if list_list_pcv[i][j].variable == list_pcv_sort[k].variable {
						list_pcv_sort[k].coef.Add(list_pcv_sort[k].coef, list_list_pcv[i][j].coef)
						again = true
					}
				}
				if !again && list_list_pcv[i][j].variable != meta_const {
					//ajouter les variables précédentes
					for list_list_pcv[i][j].variable != tab_variable[cpt_var] {
						var pair2 pair_coef_var
						pair2.coef = newRat()
						pair2.variable = tab_variable[cpt_var]
						list_pcv_sort = append(list_pcv_sort, pair2)
						cpt_var += 1
					}
					var pair2 pair_coef_var
					pair2.coef = new(big.Rat).Set(list_list_pcv[i][j].coef)
					pair2.variable = list_list_pcv[i][j].variable
					list_pcv_sort = append(list_pcv_sort, pair2)
					if cpt_var < len(tab_variable)-1 {
						cpt_var += 1
					}
				} else if list_list_pcv[i][j].variable == meta_const {

					pair.coef.Add(pair.coef, list_list_pcv[i][j].coef)
				}
			}
			again = false
			if j == len(list_list_pcv[i])-1 {
				for n := 0; n < len(tab_variable); n++ {
					present := false
					for m := 0; m < len(list_pcv_sort); m++ {
						if tab_variable[n] == list_pcv_sort[m].variable {
							present = true
						}
					}
					if !present {
						var pair2 pair_coef_var
						pair2.coef = newRat()
						pair2.variable = tab_variable[n]
						list_pcv_sort = append(list_pcv_sort, pair2)
					} else {
						present = false
					}
				}
				list_pcv_sort = append(list_pcv_sort, pair)
			}
		}
		list_list_pcv_sort = append(list_list_pcv_sort, list_pcv_sort)
	}
	return list_list_pcv_sort
}

func passe1(p types.Pred, t1 []types.Meta, t2 []types.Meta, cpt *int, tab_variable []types.Meta, b bool) []types.Meta {
	if b {
		*cpt += 2
	} else {
		*cpt += 1
	}
	tab_variable = append(tab_variable, auxPasse1(t2, p, tab_variable, 1)...)
	tab_variable = append(tab_variable, auxPasse1(t1, p, tab_variable, 0)...)

	return tab_variable
}

func auxPasse1(t []types.Meta, p types.Pred, tab_variable []types.Meta, number int) []types.Meta {
	var tab_variable2 = make([]types.Meta, 0)
	for i := 0; i < len(t); i++ {
		variable := p.GetArgs()[number].GetMetas()[i]
		present := false
		for i := 0; i < len(tab_variable); i++ {
			if tab_variable[i] == variable {
				present = true
			}
		}
		for i := 0; i < len(tab_variable2); i++ {
			if tab_variable2[i] == variable {
				present = true
			}
		}

		if !present {
			tab_variable2 = append(tab_variable2, variable)
		}
	}
	return tab_variable2
}

func passe2(list_list_pcv [][]pair_coef_var, list_pcv []pair_coef_var, connector int, val1 []pair_coef_var, val2 []pair_coef_var) [][]pair_coef_var {

	if connector != 1 {
		for i := 0; i < len(val1); i++ {
			list_pcv = append(list_pcv, val1[i])
		}
		for i := 0; i < len(val2); i++ {
			list_pcv = append(list_pcv, val2[i])
		}
		list_list_pcv = append(list_list_pcv, list_pcv)
	}

	if connector == 2 || connector == 1 {
		list_list_pcv = append(list_list_pcv, egPasse2(val1, val2))
	}

	return list_list_pcv
}

func egPasse2(val1 []pair_coef_var, val2 []pair_coef_var) []pair_coef_var {
	var list_pcv_eg = make([]pair_coef_var, 0)
	for i := 0; i < len(val1); i++ {
		egneg := new(big.Rat).Mul(val1[i].coef, big.NewRat(-1, 1))
		var pair pair_coef_var
		pair.variable = val1[i].variable
		pair.coef = egneg
		list_pcv_eg = append(list_pcv_eg, pair)
	}

	for i := 0; i < len(val2); i++ {
		egneg2 := new(big.Rat).Mul(val2[i].coef, big.NewRat(-1, 1))
		var pair2 pair_coef_var
		pair2.coef = egneg2
		pair2.variable = val2[i].variable
		list_pcv_eg = append(list_pcv_eg, pair2)
	}

	return list_pcv_eg

}

/** TODO
* C'est ici qu'on gère la conversion des variables
* Je fais beaucoup de disjonction de cas en fonction de int ou rat, mais selon votre format d'entrée ce ne sera peut-être pas nécessaire
**/
func termToSimplex(t types.Term, map_v_mv *map[string]types.Meta, iv *[]string, left bool, pcv []pair_coef_var) ([]types.Meta, []pair_coef_var, error) {

	var tab_var = make([]types.Meta, 0)
	switch ttype := t.(type) {

	case types.Meta:
		// C'est ici que je stock les metavairables, que je regarde si elles sont entière et que je fais la correspondence
		var_for_simplex := ttype.ToString()
		(*map_v_mv)[var_for_simplex] = ttype // Je stock la nouvelle variable de simplexe dans une map pour refaire le lien après
		if typing.IsInt(ttype.GetTypeHint()) {
			is_in_tab_int := false
			for i := 0; i < len(*iv); i++ {
				if var_for_simplex == (*iv)[i] {
					is_in_tab_int = true
				}
			}
			if !is_in_tab_int {
				(*iv) = append((*iv), var_for_simplex) // Je stock aussi la variable dans la liste des variables entière si elle doit être entière
			}
		}

		var pair pair_coef_var
		pair.variable = ttype
		if left {
			pair.coef = big.NewRat(1, 1)
		} else {
			pair.coef = big.NewRat(-1, 1)
		}
		pcv = append(pcv, pair)
		tab_var = append(tab_var, ttype)
		return tab_var, pcv, nil

	case types.Fun:
		switch t.GetName() {
		case "sum":
			var arg1, arg2 types.Term
			arg1 = ttype.GetArgs()[0]
			arg2 = ttype.GetArgs()[1]

			var1, _, _ := termToSimplex(arg1, map_v_mv, iv, left, pcv)
			var2, _, _ := termToSimplex(arg2, map_v_mv, iv, left, pcv)
			if var1 != nil {
				tab_var = append(tab_var, var1...)
			}
			if var2 != nil {
				tab_var = append(tab_var, var2...)
			}
			pair, _, _ := funToSimplex(ttype, map_v_mv, iv, pcv, left)
			pcv = append(pcv, pair...)

		case "product":
			var arg1, arg2 types.Term
			arg1 = ttype.GetArgs()[0]
			arg2 = ttype.GetArgs()[1]
			var1, _, _ := termToSimplex(arg1, map_v_mv, iv, left, pcv)
			var2, _, _ := termToSimplex(arg2, map_v_mv, iv, left, pcv)
			if var1 != nil {
				tab_var = append(tab_var, var1...)
			}
			if var2 != nil {
				tab_var = append(tab_var, var2...)
			}
			pair, _, _ := funToSimplex(ttype, map_v_mv, iv, pcv, left)
			pcv = append(pcv, pair...)

		case "quotient":
			var arg1, arg2 types.Term
			arg1 = ttype.GetArgs()[0]
			arg2 = ttype.GetArgs()[1]
			var1, _, _ := termToSimplex(arg1, map_v_mv, iv, left, pcv)
			var2, _, _ := termToSimplex(arg2, map_v_mv, iv, left, pcv)
			if var1 != nil {
				tab_var = append(tab_var, var1...)
			}
			if var2 != nil {
				tab_var = append(tab_var, var2...)
			}
			pair, _, _ := funToSimplex(ttype, map_v_mv, iv, pcv, left)
			pcv = append(pcv, pair...)

		case "quotient_e":
			var arg1, arg2 types.Term
			arg1 = ttype.GetArgs()[0]
			arg2 = ttype.GetArgs()[1]
			var1, _, _ := termToSimplex(arg1, map_v_mv, iv, left, pcv)
			var2, _, _ := termToSimplex(arg2, map_v_mv, iv, left, pcv)
			if var1 != nil {
				tab_var = append(tab_var, var1...)
			}
			if var2 != nil {
				tab_var = append(tab_var, var2...)
			}
			pair, _, _ := funToSimplex(ttype, map_v_mv, iv, pcv, left)
			pcv = append(pcv, pair...)

		case "quotient_t":
			var arg1, arg2 types.Term
			arg1 = ttype.GetArgs()[0]
			arg2 = ttype.GetArgs()[1]
			var1, _, _ := termToSimplex(arg1, map_v_mv, iv, left, pcv)
			var2, _, _ := termToSimplex(arg2, map_v_mv, iv, left, pcv)
			if var1 != nil {
				tab_var = append(tab_var, var1...)
			}
			if var2 != nil {
				tab_var = append(tab_var, var2...)
			}
			pair, _, _ := funToSimplex(ttype, map_v_mv, iv, pcv, left)
			pcv = append(pcv, pair...)

		case "quotient_f":
			var arg1, arg2 types.Term
			arg1 = ttype.GetArgs()[0]
			arg2 = ttype.GetArgs()[1]
			var1, _, _ := termToSimplex(arg1, map_v_mv, iv, left, pcv)
			var2, _, _ := termToSimplex(arg2, map_v_mv, iv, left, pcv)
			if var1 != nil {
				tab_var = append(tab_var, var1...)
			}
			if var2 != nil {
				tab_var = append(tab_var, var2...)
			}
			pair, _, _ := funToSimplex(ttype, map_v_mv, iv, pcv, left)
			pcv = append(pcv, pair...)

		case "difference":
			var arg1, arg2 types.Term
			arg1 = ttype.GetArgs()[0]
			arg2 = ttype.GetArgs()[1]
			var1, _, _ := termToSimplex(arg1, map_v_mv, iv, left, pcv)
			var2, _, _ := termToSimplex(arg2, map_v_mv, iv, left, pcv)
			if var1 != nil {
				tab_var = append(tab_var, var1...)
			}
			if var2 != nil {
				tab_var = append(tab_var, var2...)
			}
			pair, _, _ := funToSimplex(ttype, map_v_mv, iv, pcv, left)
			pcv = append(pcv, pair...)

		default:
			var pair pair_coef_var
			var value *big.Rat
			if checkArithmeticFun(ttype.GetID()) {
				_, value, _ = funToSimplex(ttype, map_v_mv, iv, pcv, left)
			} else {
				value, _ = new(big.Rat).SetString(t.GetName())
			}
			if left {
				value.Mul(value, big.NewRat(-1, 1))
			}

			pair.coef = value
			pcv = append(pcv, pair)
		}
	default:
		fmt.Printf("Unexpected type in termToSimplex\n")
		return tab_var, pcv, errors.New("Error")
	}
	return tab_var, pcv, nil

}

/**
* TODO
**/

func funToSimplex(f types.Fun, map_v_mv *map[string]types.Meta, iv *[]string, tab_pcv []pair_coef_var, left bool) ([]pair_coef_var, *big.Rat, error) {

	switch f.GetID().GetName() {

	case "sum":
		arg1 := f.GetArgs()[0]
		arg2 := f.GetArgs()[1]

		if arg1.IsFun() && arg2.IsMeta() {
			var additive_value *big.Rat
			tab_pcv, additive_value, _ = BinaryAdditiveFuncSimplexFunAndMeta(arg1, arg2, left, tab_pcv, map_v_mv, iv, false)
			return tab_pcv, additive_value, nil
		}

		if arg1.IsMeta() && arg2.IsFun() {
			var additive_value *big.Rat
			tab_pcv, additive_value, _ = BinaryAdditiveFuncSimplexFunAndMeta(arg2, arg1, left, tab_pcv, map_v_mv, iv, false)
			return tab_pcv, additive_value, nil
		}

		if arg1.IsMeta() && arg2.IsMeta() {
			tab_pcv = BinaryAdditiveFuncSimplex2Metas(arg1, arg2, tab_pcv, left, false)
			return tab_pcv, newRat(), nil
		}

		if arg1.IsFun() && arg2.IsFun() {
			var additive_value *big.Rat
			tab_pcv, additive_value = BinaryAdditiveFuncSimplex2Fun(arg1, arg2, tab_pcv, left, false, f, map_v_mv, iv)
			return tab_pcv, additive_value, nil
		}

	case "difference":
		arg1 := f.GetArgs()[0]
		arg2 := f.GetArgs()[1]

		if arg1.IsFun() && arg2.IsMeta() {
			var additive_value *big.Rat
			tab_pcv, additive_value, _ = BinaryAdditiveFuncSimplexFunAndMeta(arg1, arg2, left, tab_pcv, map_v_mv, iv, true)
			return tab_pcv, additive_value, nil

		}

		if arg1.IsMeta() && arg2.IsFun() {
			var additive_value *big.Rat
			//le false est en paramètre car arg2 est de type Fun
			tab_pcv, additive_value, _ = BinaryAdditiveFuncSimplexFunAndMeta(arg2, arg1, left, tab_pcv, map_v_mv, iv, false)
			return tab_pcv, additive_value, nil
		}

		if arg1.IsMeta() && arg2.IsMeta() {
			tab_pcv = BinaryAdditiveFuncSimplex2Metas(arg1, arg2, tab_pcv, left, true)
			return tab_pcv, newRat(), nil
		}

		if arg1.IsFun() && arg2.IsFun() {
			var additive_value *big.Rat
			tab_pcv, additive_value = BinaryAdditiveFuncSimplex2Fun(arg1, arg2, tab_pcv, left, true, f, map_v_mv, iv)
			return tab_pcv, additive_value, nil
		}

	case "product":
		arg1 := f.GetArgs()[0]
		arg2 := f.GetArgs()[1]

		if arg1.IsFun() && arg2.IsMeta() {
			tab_pcv = binaryMultiplicativeFuncSimplexMetaFun(arg1, arg2, tab_pcv, map_v_mv, iv, left)
			return tab_pcv, newRat(), nil
		}

		if arg1.IsMeta() && arg2.IsFun() {
			tab_pcv = binaryMultiplicativeFuncSimplexMetaFun(arg2, arg1, tab_pcv, map_v_mv, iv, left)
			return tab_pcv, newRat(), nil
		}

		if arg1.IsFun() && arg2.IsFun() {
			var value *big.Rat
			tab_pcv, value = binaryMultiplicativeFuncSimplex2Fun(arg1, arg2, tab_pcv, f, map_v_mv, iv, left)
			return tab_pcv, value, nil
		}
	case "quotient":

		arg1 := f.GetArgs()[0]
		arg2 := f.GetArgs()[1]
		if arg1.IsMeta() && arg2.IsFun() {
			tab_pcv = quotientMetaFun(arg1, arg2, "quotient", tab_pcv, left, map_v_mv, iv)
			return tab_pcv, newRat(), nil
		}

		if arg1.IsFun() && arg2.IsFun() {
			var value *big.Rat
			tab_pcv, value = binaryMultiplicativeFuncSimplex2Fun(arg1, arg2, tab_pcv, f, map_v_mv, iv, left)
			return tab_pcv, value, nil
		}
	case "quotient_e":
		arg1 := f.GetArgs()[0]
		arg2 := f.GetArgs()[1]
		if arg1.IsMeta() && arg2.IsFun() {
			tab_pcv = quotientMetaFun(arg1, arg2, "quotient_e", tab_pcv, left, map_v_mv, iv)
			return tab_pcv, newRat(), nil
		}
		if arg1.IsFun() && arg2.IsFun() {
			var value *big.Rat
			tab_pcv, value = binaryMultiplicativeFuncSimplex2Fun(arg1, arg2, tab_pcv, f, map_v_mv, iv, left)
			return tab_pcv, value, nil
		}

	case "quotient_t":
		arg1 := f.GetArgs()[0]
		arg2 := f.GetArgs()[1]
		if arg1.IsMeta() && arg2.IsFun() {
			tab_pcv = quotientMetaFun(arg1, arg2, "quotient_t", tab_pcv, left, map_v_mv, iv)
			return tab_pcv, newRat(), nil
		}
		if arg1.IsFun() && arg2.IsFun() {
			var value *big.Rat
			tab_pcv, value = binaryMultiplicativeFuncSimplex2Fun(arg1, arg2, tab_pcv, f, map_v_mv, iv, left)
			return tab_pcv, value, nil
		}
	case "quotient_f":
		arg1 := f.GetArgs()[0]
		arg2 := f.GetArgs()[1]
		if arg1.IsMeta() && arg2.IsFun() {
			tab_pcv = quotientMetaFun(arg1, arg2, "quotient_f", tab_pcv, left, map_v_mv, iv)
			return tab_pcv, newRat(), nil
		}
		if arg1.IsFun() && arg2.IsFun() {
			var value *big.Rat
			tab_pcv, value = binaryMultiplicativeFuncSimplex2Fun(arg1, arg2, tab_pcv, f, map_v_mv, iv, left)
			return tab_pcv, value, nil
		}
	case "remainder_e":
		arg1 := f.GetArgs()[0]
		arg2 := f.GetArgs()[1]

		remainder := remainderFuncSimplex(arg1, arg2, map_v_mv, iv, tab_pcv, left, f)

		var pair pair_coef_var
		pair.coef = remainder
		tab_pcv = append(tab_pcv, pair)

		return tab_pcv, remainder, nil

	case "remainder_t":
		arg1 := f.GetArgs()[0]
		arg2 := f.GetArgs()[1]

		remainder := remainderFuncSimplex(arg1, arg2, map_v_mv, iv, tab_pcv, left, f)

		var pair pair_coef_var
		pair.coef = remainder
		tab_pcv = append(tab_pcv, pair)

		return tab_pcv, remainder, nil

	case "remainder_f":
		arg1 := f.GetArgs()[0]
		arg2 := f.GetArgs()[1]

		remainder := remainderFuncSimplex(arg1, arg2, map_v_mv, iv, tab_pcv, left, f)

		var pair pair_coef_var
		pair.coef = remainder
		tab_pcv = append(tab_pcv, pair)

		return tab_pcv, remainder, nil

	case "uminus":
		arg := f.GetArgs()[0]
		uminus := unaryFuncSimplex(arg, f, map_v_mv, iv, tab_pcv, left)
		var pair pair_coef_var
		pair.coef = uminus
		tab_pcv = append(tab_pcv, pair)
		return tab_pcv, uminus, nil

	case "floor":
		arg := f.GetArgs()[0]
		floor := unaryFuncSimplex(arg, f, map_v_mv, iv, tab_pcv, left)
		var pair pair_coef_var
		pair.coef = floor
		tab_pcv = append(tab_pcv, pair)
		return tab_pcv, floor, nil

	case "ceiling":
		arg := f.GetArgs()[0]
		ceiling := unaryFuncSimplex(arg, f, map_v_mv, iv, tab_pcv, left)
		var pair pair_coef_var
		pair.coef = ceiling
		tab_pcv = append(tab_pcv, pair)
		return tab_pcv, ceiling, nil

	case "truncate":
		arg := f.GetArgs()[0]
		truncate := unaryFuncSimplex(arg, f, map_v_mv, iv, tab_pcv, left)
		var pair pair_coef_var
		pair.coef = truncate
		tab_pcv = append(tab_pcv, pair)
		return tab_pcv, truncate, nil

	case "round":
		arg := f.GetArgs()[0]
		round := unaryFuncSimplex(arg, f, map_v_mv, iv, tab_pcv, left)
		var pair pair_coef_var
		pair.coef = round
		tab_pcv = append(tab_pcv, pair)
		return tab_pcv, round, nil

	default:
		return tab_pcv, newRat(), errors.New("Error")
	}
	return tab_pcv, newRat(), errors.New("Error")
}

func remainderFuncSimplex(arg1 types.Term, arg2 types.Term, map_v_mv *map[string]types.Meta, iv *[]string, tab_pcv []pair_coef_var, left bool, f types.Fun) *big.Rat {

	var value *big.Rat
	if arg_fun1, ok := arg1.(types.Fun); ok {
		if arg_fun2, ok := arg2.(types.Fun); ok {
			//reste entre 2 nombres
			if !checkArithmeticFun(arg_fun1.GetID()) && !checkArithmeticFun(arg_fun2.GetID()) {
				value, _ = EvaluateFun(f)
				//pour ne pas fausser les calculs
				value.Mul(value, big.NewRat(-1, 1))
				return value
				//reste entre 2 fonction
			} else if checkArithmeticFun(arg_fun1.GetID()) && checkArithmeticFun(arg_fun2.GetID()) {
				_, val, _ := funToSimplex(arg_fun1, map_v_mv, iv, tab_pcv, left)
				//pour ne pas fausser les calculs
				if checkBinaryArithmeticFun(arg_fun1.GetID()) {
					val.Mul(val, big.NewRat(-1, 1))
				}
				_, val_bis, _ := funToSimplex(arg_fun2, map_v_mv, iv, tab_pcv, left)
				//pour ne pas fausser les calculs
				if checkBinaryArithmeticFun(arg_fun2.GetID()) {
					val_bis.Mul(val_bis, big.NewRat(-1, 1))
				}
				val2 := types.MakerConst(types.MakerId(val.RatString()), tRat)
				val2_bis := types.MakerConst(types.MakerId(val_bis.RatString()), tRat)
				argId := types.MakerFun(types.MakerId(f.GetID().ToString()), []types.Term{val2, val2_bis}, typing.GetTypeScheme(f.GetID().ToString(), typing.MkTypeCross(tRat, tRat)))
				value, _ = EvaluateFun(argId)
				if checkBinaryArithmeticFun(arg_fun1.GetID()) || checkBinaryArithmeticFun(arg_fun2.GetID()) {
					value.Mul(value, big.NewRat(-1, 1))
				}
				return value
				// arg1 est un nombre mais pas arg2
			} else if !checkArithmeticFun(arg_fun1.GetID()) {
				_, val, _ := funToSimplex(arg_fun2, map_v_mv, iv, tab_pcv, left)
				//pour ne pas fausser les calculs
				if checkBinaryArithmeticFun(arg_fun2.GetID()) {
					val.Mul(val, big.NewRat(-1, 1))
				}
				val2 := types.MakerConst(types.MakerId(val.RatString()), tRat)
				argId := types.MakerFun(types.MakerId(f.GetID().ToString()), []types.Term{arg1, val2}, typing.GetTypeScheme(f.GetID().ToString(), typing.MkTypeCross(arg_fun1.GetTypeHint(), tRat)))
				value, _ = EvaluateFun(argId)
				value.Mul(value, big.NewRat(-1, 1))
				return value

			} else if !checkArithmeticFun(arg_fun2.GetID()) {
				_, val, _ := funToSimplex(arg_fun1, map_v_mv, iv, tab_pcv, left)
				//pour ne pas fausser les calculs
				if checkBinaryArithmeticFun(arg_fun1.GetID()) {
					val.Mul(val, big.NewRat(-1, 1))
				}
				val2 := types.MakerConst(types.MakerId(val.RatString()), tRat)
				argId := types.MakerFun(types.MakerId(f.GetID().ToString()), []types.Term{val2, arg2}, typing.GetTypeScheme(f.GetID().ToString(), typing.MkTypeCross(tRat, arg_fun2.GetTypeHint())))
				value, _ = EvaluateFun(argId)
				value.Mul(value, big.NewRat(-1, 1))
				return value

			}
		}
	}
	return value
}

func quotientMetaFun(arg1 types.Term, arg2 types.Term, operation string, tab_pcv []pair_coef_var, left bool, map_v_mv *map[string]types.Meta, iv *[]string) []pair_coef_var {

	if arg_fun2, ok2 := arg2.(types.Fun); ok2 {
		var unVal string
		if arg_fun2.GetTypeHint().ToString() == "rat" {
			unVal = "1/1"
		} else {
			unVal = "1"
		}
		if !checkArithmeticFun(arg_fun2.GetID()) {
			un := types.MakerConst(types.MakerId(unVal), arg_fun2.GetTypeHint())
			quo := types.MakerFun(types.MakerId(operation), []types.Term{un, arg2}, typing.GetTypeScheme(operation, typing.MkTypeCross(tRat, arg_fun2.GetTypeHint())))

			solution, _ := EvaluateFun(quo)

			var pair pair_coef_var
			if !left {
				solution.Mul(solution, big.NewRat(-1, 1))
			}
			pair.coef = solution
			if arg_meta1, ok1 := arg1.(types.Meta); ok1 {
				pair.variable = arg_meta1
			}
			tab_pcv = append(tab_pcv, pair)
		} else {
			_, value, _ := funToSimplex(arg_fun2, map_v_mv, iv, tab_pcv, left)
			var pair pair_coef_var
			value.Inv(value)
			if !left {
				value.Mul(value, big.NewRat(-1, 1))
			}
			pair.coef = value
			if arg_meta1, ok1 := arg1.(types.Meta); ok1 {
				pair.variable = arg_meta1
			}
			tab_pcv = append(tab_pcv, pair)
		}
	}
	return tab_pcv
}

func binaryMultiplicativeFuncSimplex2Fun(arg1 types.Term, arg2 types.Term, tab_pcv []pair_coef_var, f types.Fun, map_v_mv *map[string]types.Meta, iv *[]string, left bool) ([]pair_coef_var, *big.Rat) {

	var value *big.Rat
	if arg_fun2, ok2 := arg2.(types.Fun); ok2 {
		if arg_fun1, ok1 := arg1.(types.Fun); ok1 {
			//cas 2 nombres
			if !checkArithmeticFun(arg_fun1.GetID()) && !checkArithmeticFun(arg_fun2.GetID()) {
				multiplicative_value, _ := EvaluateFun(f)
				if left {
					multiplicative_value.Mul(multiplicative_value, big.NewRat(-1, 1))
				}
				//si somme des nombres non lié a une méta
				var pair pair_coef_var
				pair.coef = multiplicative_value
				tab_pcv = append(tab_pcv, pair)
				//multiplicative value est renvoyé pour le cas o`u la valeur est liée a une méta
				return tab_pcv, multiplicative_value
				//cas pas de fonction binaire
			} else if !checkBinaryArithmeticFun(arg_fun1.GetID()) && !checkBinaryArithmeticFun(arg_fun2.GetID()) {
				if checkArithmeticFun(arg_fun1.GetID()) && !checkArithmeticFun(arg_fun2.GetID()) {
					val, _ := new(big.Rat).SetString(arg_fun2.GetName())
					tab_pcv, _, _ := funToSimplex(arg_fun1, map_v_mv, iv, tab_pcv, left)
					val = tab_pcv[len(tab_pcv)-1].coef.Mul(tab_pcv[len(tab_pcv)-1].coef, val)
					return tab_pcv, val
				} else if checkArithmeticFun(arg_fun2.GetID()) && !checkArithmeticFun(arg_fun1.GetID()) {
					val, _ := new(big.Rat).SetString(arg_fun1.GetName())
					tab_pcv, _, _ := funToSimplex(arg_fun2, map_v_mv, iv, tab_pcv, left)
					val = tab_pcv[len(tab_pcv)-1].coef.Mul(tab_pcv[len(tab_pcv)-1].coef, val)
					return tab_pcv, val
					//cas 2 fonctions unaires
				} else {
					tab_pcv, _, _ := funToSimplex(arg_fun1, map_v_mv, iv, tab_pcv, left)
					val := tab_pcv[len(tab_pcv)-1].coef
					_, value, _ := funToSimplex(arg_fun2, map_v_mv, iv, tab_pcv, left)
					val.Mul(val, value)
					return tab_pcv, val

				}
				//arg2 est fonction binaire mais pas arg1
			} else if !checkBinaryArithmeticFun(arg_fun1.GetID()) {
				//arg1 est un nombre
				if !checkArithmeticFun(arg_fun1.GetID()) {
					tab_pcv, _, _ := funToSimplex(arg_fun2, map_v_mv, iv, tab_pcv, left)
					val := tab_pcv[len(tab_pcv)-1].coef
					res_binaire := types.MakerConst(types.MakerId(val.RatString()), tRat)
					multiplicative := types.MakerFun(types.MakerId(f.GetName()), []types.Term{arg1, res_binaire}, typing.GetTypeScheme(f.GetName(), typing.MkTypeCross(tRat, tRat)))
					solution, _ := EvaluateFun(multiplicative)
					tab_pcv[len(tab_pcv)-1].coef = solution
					return tab_pcv, solution
					//arg1 est une fonction unaire
				} else {
					tab_pcv, _, _ := funToSimplex(arg_fun2, map_v_mv, iv, tab_pcv, left)
					val := tab_pcv[len(tab_pcv)-1].coef
					_, value, _ := funToSimplex(arg_fun1, map_v_mv, iv, tab_pcv, left)
					res_binaire := types.MakerConst(types.MakerId(val.RatString()), tRat)
					res_unaire := types.MakerConst(types.MakerId(value.RatString()), tRat)
					multiplicative := types.MakerFun(types.MakerId(f.GetName()), []types.Term{res_unaire, res_binaire}, typing.GetTypeScheme(f.GetName(), typing.MkTypeCross(tRat, tRat)))
					solution, _ := EvaluateFun(multiplicative)
					tab_pcv[len(tab_pcv)-1].coef = solution
					return tab_pcv, solution
				}
			} else if !checkBinaryArithmeticFun(arg_fun2.GetID()) {
				if !checkArithmeticFun(arg_fun2.GetID()) {
					tab_pcv, _, _ := funToSimplex(arg_fun1, map_v_mv, iv, tab_pcv, left)
					val := tab_pcv[len(tab_pcv)-1].coef
					res_binaire := types.MakerConst(types.MakerId(val.RatString()), tRat)
					multiplicative := types.MakerFun(types.MakerId(f.GetName()), []types.Term{res_binaire, arg2}, typing.GetTypeScheme(f.GetName(), typing.MkTypeCross(tRat, tRat)))
					solution, _ := EvaluateFun(multiplicative)
					tab_pcv[len(tab_pcv)-1].coef = solution
					return tab_pcv, solution
				} else {
					tab_pcv, _, _ := funToSimplex(arg_fun1, map_v_mv, iv, tab_pcv, left)
					val := tab_pcv[len(tab_pcv)-1].coef
					_, value, _ := funToSimplex(arg_fun2, map_v_mv, iv, tab_pcv, left)
					res_binaire := types.MakerConst(types.MakerId(val.RatString()), tRat)
					res_unaire := types.MakerConst(types.MakerId(value.RatString()), tRat)
					multiplicative := types.MakerFun(types.MakerId(f.GetName()), []types.Term{res_binaire, res_unaire}, typing.GetTypeScheme(f.GetName(), typing.MkTypeCross(tRat, tRat)))
					solution, _ := EvaluateFun(multiplicative)
					tab_pcv[len(tab_pcv)-1].coef = solution
					return tab_pcv, solution
				}
				//2 fonctions binaires
			} else {
				tab_pcv, _, _ := funToSimplex(arg_fun1, map_v_mv, iv, tab_pcv, left)
				val := tab_pcv[len(tab_pcv)-1].coef
				_, value, _ := funToSimplex(arg_fun2, map_v_mv, iv, tab_pcv, left)
				if f.GetName() != "product" {
					value.Inv(value)
				}
				val.Mul(val, value)
				return tab_pcv, val
			}
		}
	}
	return tab_pcv, value
}

func binaryMultiplicativeFuncSimplexMetaFun(arg1 types.Term, arg2 types.Term, tab_pcv []pair_coef_var, map_v_mv *map[string]types.Meta, iv *[]string, left bool) []pair_coef_var {

	if arg_fun, ok := arg1.(types.Fun); ok {
		if arg_meta, ok2 := arg2.(types.Meta); ok2 {
			var pair pair_coef_var
			pair.variable = arg_meta
			if !checkArithmeticFun(arg_fun.GetID()) {
				coef, _ := new(big.Rat).SetString(arg1.GetName())
				pair.coef = coef
				if !left {
					pair.coef.Mul(pair.coef, big.NewRat(-1, 1))
				}
			} else if checkBinaryArithmeticFun(arg_fun.GetID()) {
				_, val, _ := funToSimplex(arg_fun, map_v_mv, iv, tab_pcv, left)
				//si j'ai une méta à gauche, je multiplie par -1 pour retrouver le bon signe du coef. Si la meta est à droite, je multiplie par -1 pour obtenir le bon signe
				val.Mul(val, big.NewRat(-1, 1))
				pair.coef = val
			} else {
				//arg_fun est une fonction unaire
				_, val, _ := funToSimplex(arg_fun, map_v_mv, iv, tab_pcv, left)
				if left {
					val.Mul(val, big.NewRat(-1, 1))
				}
				pair.coef = val

			}
			tab_pcv = append(tab_pcv, pair)
		}
	}
	return tab_pcv
}

func unaryFuncSimplex(arg types.Term, f types.Fun, map_v_mv *map[string]types.Meta, iv *[]string, tab_pcv []pair_coef_var, left bool) *big.Rat {
	var value *big.Rat
	if arg_fun, ok := arg.(types.Fun); ok {
		if !checkArithmeticFun(arg_fun.GetID()) {
			value, _ = EvaluateFun(f)
			return value
		} else {
			_, val, _ := funToSimplex(arg_fun, map_v_mv, iv, tab_pcv, left)
			//pour ne pas fausser les calculs
			if checkBinaryArithmeticFun(arg_fun.GetID()) {
				val.Mul(val, big.NewRat(-1, 1))
			}
			val2 := types.MakerConst(types.MakerId(val.RatString()), tRat)
			argId := types.MakerFun(types.MakerId(f.GetID().ToString()), []types.Term{val2}, tRat)
			value, _ = EvaluateFun(argId)
			return value.Mul(value, big.NewRat(-1, 1))
		}
	}
	return value
}

func BinaryAdditiveFuncSimplex2Metas(arg1 types.Term, arg2 types.Term, tab_pcv []pair_coef_var, left bool, diff bool) []pair_coef_var {
	if arg_meta1, ok1 := arg1.(types.Meta); ok1 {
		if arg_meta2, ok2 := arg2.(types.Meta); ok2 {
			var pair1, pair2 pair_coef_var
			pair1.variable = arg_meta1
			pair2.variable = arg_meta2
			if left {
				pair1.coef = big.NewRat(1, 1)
				if !diff {
					pair2.coef = big.NewRat(1, 1)
				} else {
					pair2.coef = big.NewRat(-1, 1)
				}
			} else {
				pair1.coef = big.NewRat(-1, 1)
				if !diff {
					pair2.coef = big.NewRat(-1, 1)
				} else {
					pair2.coef = big.NewRat(1, 1)
				}
			}
			tab_pcv = append(tab_pcv, pair1)
			tab_pcv = append(tab_pcv, pair2)
		}
	}
	return tab_pcv
}

func BinaryAdditiveFuncSimplex2Fun(arg1 types.Term, arg2 types.Term, tab_pcv []pair_coef_var, left bool, diff bool, f types.Fun, map_v_mv *map[string]types.Meta, iv *[]string) ([]pair_coef_var, *big.Rat) {

	if arg_fun2, ok2 := arg2.(types.Fun); ok2 {
		if arg_fun1, ok1 := arg1.(types.Fun); ok1 {
			//si j'ai 2 nombres
			if !checkArithmeticFun(arg_fun1.GetID()) && !checkArithmeticFun(arg_fun2.GetID()) {
				additive_value, _ := EvaluateFun(f)
				if !diff {
					if left {
						additive_value.Mul(additive_value, big.NewRat(-1, 1))
					}
				}
				//si somme des nombres non lié a une méta
				var pair pair_coef_var
				pair.coef = additive_value
				tab_pcv = append(tab_pcv, pair)
				//additive value est renvoyé pour le cas o`u la valeur est liée a une méta
				return tab_pcv, additive_value
			} else if !checkArithmeticFun(arg_fun1.GetID()) {
				var pair pair_coef_var
				number_value, _ := new(big.Rat).SetString(arg_fun1.GetName())
				if left {
					number_value.Mul(number_value, big.NewRat(-1, 1))
				}
				pair.coef = number_value
				if diff {
					vpcv2, _, _ := funToSimplex(arg_fun2, map_v_mv, iv, tab_pcv, !left)
					tab_pcv = append(tab_pcv, pair)
					tab_pcv = append(tab_pcv, vpcv2...)
				} else {
					vpcv2, _, _ := funToSimplex(arg_fun2, map_v_mv, iv, tab_pcv, left)
					tab_pcv = append(tab_pcv, pair)
					tab_pcv = append(tab_pcv, vpcv2...)
				}

			} else if !checkArithmeticFun(arg_fun2.GetID()) {
				var pair pair_coef_var
				number_value, _ := new(big.Rat).SetString(arg_fun2.GetName())
				pair.coef = number_value
				tab_pcv = append(tab_pcv, pair)
				if diff {
					vpcv1, _, _ := funToSimplex(arg_fun1, map_v_mv, iv, tab_pcv, left)
					tab_pcv = append(tab_pcv, vpcv1...)
				} else {
					vpcv1, _, _ := funToSimplex(arg_fun1, map_v_mv, iv, tab_pcv, !left)
					tab_pcv = append(tab_pcv, vpcv1...)
				}
			} else if checkArithmeticFun(arg_fun1.GetID()) && checkArithmeticFun(arg_fun2.GetID()) {
				if diff {
					if left {
						vpcv1, _, _ := funToSimplex(arg_fun1, map_v_mv, iv, tab_pcv, left)
						vpcv2, _, _ := funToSimplex(arg_fun2, map_v_mv, iv, tab_pcv, !left)
						tab_pcv = append(tab_pcv, vpcv1...)
						tab_pcv = append(tab_pcv, vpcv2...)
					} else {
						vpcv1, _, _ := funToSimplex(arg_fun1, map_v_mv, iv, tab_pcv, !left)
						vpcv2, _, _ := funToSimplex(arg_fun2, map_v_mv, iv, tab_pcv, left)
						tab_pcv = append(tab_pcv, vpcv1...)
						tab_pcv = append(tab_pcv, vpcv2...)
					}
				} else {
					if !left {
						vpcv1, _, _ := funToSimplex(arg_fun1, map_v_mv, iv, tab_pcv, !left)
						vpcv2, _, _ := funToSimplex(arg_fun2, map_v_mv, iv, tab_pcv, !left)
						tab_pcv = append(tab_pcv, vpcv1...)
						tab_pcv = append(tab_pcv, vpcv2...)
					} else {
						vpcv1, _, _ := funToSimplex(arg_fun1, map_v_mv, iv, tab_pcv, left)
						if !checkBinaryArithmeticFun(arg_fun1.GetID()) {
							vpcv1[0].coef.Mul(vpcv1[0].coef, big.NewRat(-1, 1))
						}
						vpcv2, _, _ := funToSimplex(arg_fun2, map_v_mv, iv, tab_pcv, left)
						if !checkBinaryArithmeticFun(arg_fun2.GetID()) {
							vpcv2[0].coef.Mul(vpcv2[0].coef, big.NewRat(-1, 1))
						}
						tab_pcv = append(tab_pcv, vpcv1...)
						tab_pcv = append(tab_pcv, vpcv2...)
					}
				}

			}
		}
	}
	return tab_pcv, newRat()
}

func BinaryAdditiveFuncSimplexFunAndMeta(arg1 types.Term, arg2 types.Term, left bool, tab_pcv []pair_coef_var, map_v_mv *map[string]types.Meta, iv *[]string, diff bool) ([]pair_coef_var, *big.Rat, error) {

	if arg_fun, ok := arg1.(types.Fun); ok {
		var pair pair_coef_var
		if arg_meta, ok2 := arg2.(types.Meta); ok2 {
			pair.variable = arg_meta
			if left {
				if !diff {
					pair.coef = big.NewRat(1, 1)
				} else {
					pair.coef = big.NewRat(-1, 1)
				}
			} else {
				if !diff {
					pair.coef = big.NewRat(-1, 1)
				} else {
					pair.coef = big.NewRat(1, 1)
				}
			}
			tab_pcv = append(tab_pcv, pair)
		}

		if !checkBinaryArithmeticFun(arg_fun.GetID()) {
			var pair2 pair_coef_var
			additive_value, _ := new(big.Rat).SetString(arg_fun.GetName())
			pair2.coef = additive_value
			if left {
				additive_value.Mul(additive_value, big.NewRat(-1, 1))
			}
			tab_pcv = append(tab_pcv, pair2)
		} else {
			return funToSimplex(arg_fun, map_v_mv, iv, tab_pcv, left)
		}
	}
	return tab_pcv, newRat(), nil
}

func checkArithmeticFun(id types.Id) bool {
	s := id.ToString()
	return s == "sum" || s == "product" || s == "difference" || s == "quotient" || s == "quotient_e" || s == "quotient_t" || s == "quotient_f" || s == "remainder_e" || s == "remainder_t" || s == "remainder_f" || s == "floor" || s == "ceiling" || s == "uminus" || s == "truncate" || s == "round"
}
