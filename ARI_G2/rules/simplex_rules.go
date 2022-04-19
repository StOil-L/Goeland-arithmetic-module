/******************/
/* simplex_rules.go */
/******************/
/**
* This file provides rules to manage simplex rules
**/

package ari

import (
	typing "ARI/polymorphism"
	treetypes "ARI/tree-types"
	"ARI/types"
	"errors"
	"fmt"
	"math/big"
	"strconv"
)

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
func simplex([]string, []string) (bool, map[string]*big.Rat) {
	return false, nil
}

/* Normalize equation for simplex call */
/**
* TODO :
* Le type de retour doit être celui que prend votre simplexe en entrée
* Je suppose que vous allez parser les prédicat et les transforer en >= ou <=
* Je vous mets le schéma de base, à vous de l'adapter et de savoir quoi renvoyer !
*
* On va renvoyer 3 choses :
* la strcuture/liste pour votre simplexe (ci on dit que c'est une liste de string)
* une map de correspondence entre les metavariables que vous allez utiliser et celles que j'ai moi. On va supposer que vos variables sont des strings
* Une liste des variables qui doivent être entières
*
* En théorie dans la liste des prédicat il n'y aura que des <, >, <=, >=, mais dans le doute on va quand-même traiter l'égalité
* Ca va beaucoup ressembler aux autre fonctions de parse que vous avez, sauf que là au lieu que les deux côtés soient des ints ou des rats, vous allez avoir des métavariables à convertir
* Les deux listes qui concernent les metavariables sont à utiliser comme des pointeurs pour qu'elles puissent se construire tout au long du parsing
*
* Il faut compléter la fonction
**/
func normalizeForSimplex(pl []types.Pred) ([]string, map[string]types.Meta, []string) {
	res_for_simplex := []string{}
	map_variable_metavariables := make(map[string]types.Meta)
	int_variables := []string{}
	for _, p := range pl {
		switch p.GetID().GetName() {
		case types.Id_eq.GetName():
			// si j'ai une égalité
			// Ce qui suit est un exemple et dépend de votre format d'entrée
			// ici je dis que si j'ai par exemple x = 2 alors je transforme ça en x >= 2 et 2 >= x
			t1, err1 := termToSimplex(p.GetArgs()[0], &map_variable_metavariables, &int_variables)
			t2, err2 := termToSimplex(p.GetArgs()[1], &map_variable_metavariables, &int_variables)
			if err1 != nil || err2 != nil {
				fmt.Printf("Error in normalizeForSimplex")
				return nil, nil, nil
			}
			res_for_simplex = append(res_for_simplex, t1+" >= "+t2)
			res_for_simplex = append(res_for_simplex, "-"+t1+" >= "+"-"+t2)
		case types.Id_neq.GetName():
		case "less":
			// Si j'ai un <
			// Ce qui suit est un exemple et dépend de votre format d'entrée
			// cas x < 3 ou sum(x, y) < 4 par exemple, et vous voulez le transformer en -x > -3
			t1, err1 := termToSimplex(p.GetArgs()[0], &map_variable_metavariables, &int_variables)

			t2, err2 := termToSimplex(p.GetArgs()[1], &map_variable_metavariables, &int_variables)

			if err1 != nil || err2 != nil {
				fmt.Printf("Error in normalizeForSimplex")
				return nil, nil, nil
			}

			fmt.Println("t1, t2",t1,t2)
			res_for_simplex = append(res_for_simplex, "-"+t1+" > "+"-"+t2)
		case "lesseq":
		case "great":
		case "greateq":
		}
	}

	return res_for_simplex, map_variable_metavariables, int_variables
}

/** TODO
* C'est ici qu'on gère la conversion des variables
* Je fais beaucoup de disjonction de cas en fonction de int ou rat, mais selon votre format d'entrée ce ne sera peut-être pas nécessaire
**/
func termToSimplex(t types.Term, map_v_mv *map[string]types.Meta, iv *[]string) (string, error) {
	switch ttype := t.(type) {
	case types.Meta:
		// C'est ici que je stock les metavairables, que je regarde si elles sont entière et que je fais la correspondence
		var_for_simplex := ttype.ToString()
		(*map_v_mv)[var_for_simplex] = ttype // Je stock la nouvelle variable de simplexe dans une map pour refaire le lien après
		if typing.IsInt(ttype.GetTypeHint()) {
			(*iv) = append((*iv), var_for_simplex) // Je stock aussi la variable dans la liste des variables entière si elle doit être entière
		}
		return var_for_simplex, nil
	case types.Fun:
		switch {
		// Si j'ai des variables dans l'équation
		case len(ttype.GetMetas()) > 0:
			return funToSimplex(ttype, map_v_mv, iv)

		// Si c'est un résultat entier
		case typing.IsInt(ttype.GetTypeHint()):
			return "", nil
			res, err := FunToInt(ttype)
			fmt.Println("test")

			if err != nil {
				fmt.Printf("Error in termToSimplex\n")
				return "", err
			}
			return strconv.Itoa(res), nil

		// Si c'est un résultat décimal     
		case typing.IsRat(ttype.GetTypeHint()):
			res, err := funToRat(ttype)
			if err != nil {
				fmt.Printf("Error in termToSimplex\n")
				return "", err
			}
			return res.String(), nil

		default:
			fmt.Printf("Unexpected type in termToSimplex\n")
			return "", errors.New("Error")
		}
	default:
		fmt.Printf("Unexpected type in termToSimplex\n")
		return "", errors.New("Error")
	}
}

/**
* TODO
* ici, j'ai une fonction avec une metavariable, comme sum(x,3)
* Je fais beaucoup de disjonction de cas en fonction de int ou rat, mais selon votre format d'entrée ce ne sera peut-être pas nécessaire
*
* Il faut compléter la fonction
**/
func funToSimplex(f types.Fun, map_v_mv *map[string]types.Meta, iv *[]string) (string, error) {
	switch f.GetID().GetName() {

	case "sum":
		arg1 := f.GetArgs()[0]
		arg2 := f.GetArgs()[1]

		switch {
		// Si la/les variables(s) sont dans t2
		case len(arg1.GetMetas()) == 0:
			// je cherck le type de t1
			switch f.GetTypeHint() {
			case tInt:
				res1, err1 := termToInt(arg1)
				if err1 != nil {
					fmt.Printf("Error in funToSimplex : sum\n")
					return "", err1
				}
				res2, err2 := termToSimplex(arg2, map_v_mv, iv)
				if err2 != nil {
					return "", err2
				}
				return strconv.Itoa(res1) + " + " + res2, nil
			case tRat:
				res1, err1 := termToRat(arg1)
				if err1 != nil {
					fmt.Printf("Error in funToSimplex : sum\n")
					return "", err1
				}
				res2, err2 := termToSimplex(arg2, map_v_mv, iv)
				if err2 != nil {
					return "", err2
				}
				return res1.String() + " + " + res2, nil
			default:
				fmt.Printf("Error funToSimplex : type not found\n")
				return "", errors.New("Error")
			}

		// Si la/les variables(s) sont dans t1
		case len(arg2.GetMetas()) == 0:
			// je cherck le type de t2
			switch f.GetTypeHint() {
			case tInt:
				res1, err1 := termToSimplex(arg1, map_v_mv, iv)
				if err1 != nil {
					fmt.Printf("Error in funToSimplex : sum\n")
					return "", err1
				}
				res2, err2 := termToInt(arg2)
				if err2 != nil {
					return "", err2
				}
				return res1 + " + " + strconv.Itoa(res2), nil
			case tRat:
				res1, err1 := termToSimplex(arg1, map_v_mv, iv)
				if err1 != nil {
					fmt.Printf("Error in funToSimplex : sum\n")
					return "", err1
				}
				res2, err2 := termToRat(arg2)
				if err2 != nil {
					return "", err2
				}
				return res1 + " + " + res2.String(), nil
			default:
				fmt.Printf("Error funToSimplex : type not found\n")
				return "", errors.New("Error")
			}
		default:
			fmt.Printf("Error funToSimplex : type not found\n")
			return "", errors.New("Error")
		}

	case "difference":
	case "product":
	case "quotient":
	case "quotient_e":
	case "quotient_t":
	case "quotient_f":
	case "remainder_e":
	case "remainder_t":
	case "remainder_f":
	case "uminus":
	case "floor":
	case "ceiling":
	case "truncate":
	case "round":
	default:
		return "", errors.New("Error")
	}
	return "", errors.New("Error")
}
