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
* la strcuture/liste pour votre simplexe
* une map de correspondence entre les metavariables que vous allez utiliser et celles que j'ai moi. On va supposer que vos variables sont des strings
* Une liste des variables qui doivent être entières
*
* En théorie dans la liste des prédicat il n'y aura que des <, >, <=, >=, mais dans le doute on va quand-même traiter l'égalité
* Ca va beaucoup ressembler aux autre fonctions de parse que vous avez, sauf que là au lieu que les deux côtés soient des ints ou des rats, vous allez avoir des métavariables à convertir
* Les deux listes qui concernent les metavariables sont à utiliser comme des pointeurs pour qu'elles puissent se construire tout au long du parsing
**/

func normalizeForSimplex(pl []types.Pred) ([]string, map[string]types.Meta, []string) {
	res_for_simplex := []string{}
	map_variable_metavariables := make(map[string]types.Meta)
	int_variables := []string{}
	var tab_variable = make([]string, 0)
	cpt:=0
	//passe:=0
	passe:=2
	for passe<3{
		for _, p := range pl {
			switch p.GetID().GetName() {
		
			case types.Id_eq.GetName():
				t1, err1 := termToSimplex(p.GetArgs()[0], &map_variable_metavariables, &int_variables)
				t2, err2 := termToSimplex(p.GetArgs()[1], &map_variable_metavariables, &int_variables)
				if err1 != nil || err2 != nil {
					fmt.Printf("Error in normalizeForSimplex")
					return nil, nil, nil
				}
				tab_variable=passe1eg(p,t1,t2,&cpt,tab_variable)
				fmt.Println("ici", tab_variable)		
		
			case types.Id_neq.GetName():

			//à réfléchir, si on a 2x != 3 alors on a 2x > 3  OU  2x < 3
			//il faudrait donc faire 2 systèmes, un avec chacune des équations.. 
			//et encore, ça ne marche que si on cherche une solution entière..
			
			
			case "less":
				t1, err1 := termToSimplex(p.GetArgs()[0], &map_variable_metavariables, &int_variables)
				t2, err2 := termToSimplex(p.GetArgs()[1], &map_variable_metavariables, &int_variables)
				if err1 != nil || err2 != nil {
					fmt.Printf("Error in normalizeForSimplex")
					return nil, nil, nil
				}
				fmt.Println(t1,t2)
				res_for_simplex = append(res_for_simplex, " > ")
			
			case "lesseq":
				t1, err1 := termToSimplex(p.GetArgs()[0], &map_variable_metavariables, &int_variables)
				t2, err2 := termToSimplex(p.GetArgs()[1], &map_variable_metavariables, &int_variables)
				if err1 != nil || err2 != nil  {
					fmt.Printf("Error in normalizeForSimplex")
					return nil, nil, nil
				}
				
	
				tab_variable=passe1greateq_lesseq(p,t1,t2,&cpt,tab_variable)

			case "great":


			case "greateq":
				t1, err1 := termToSimplex(p.GetArgs()[0], &map_variable_metavariables, &int_variables)
				t2, err2 := termToSimplex(p.GetArgs()[1], &map_variable_metavariables, &int_variables)
				if err1 != nil || err2 != nil  {
					fmt.Printf("Error in normalizeForSimplex")
					return nil, nil, nil
				}

				tab_variable=passe1greateq_lesseq(p,t1,t2,&cpt,tab_variable)

			}
			fmt.Println("tab_var : ",tab_variable)
			fmt.Println("cpt = ",cpt)
		
			
		}
		passe+=1
	}

	return res_for_simplex, map_variable_metavariables, int_variables
}





func passe1eg(p types.Pred,t1 []string, t2 []string, cpt *int,tab_variable []string) ([]string){
	*cpt+=2
	//je code svp jugez pas :p
	present:=false
	if len(t1)==0 && t2!=nil{
		for i:=0;i<len(t2);i++{
			variable := p.GetArgs()[1].GetMetas()[i]
			present=false
			fmt.Println("variable2 =",variable.GetName())
			for i:=0;i<len(tab_variable);i++{
				if tab_variable[i]==variable.GetName(){
					present=true
				}
			}
			if !present{
				tab_variable=append(tab_variable,variable.GetName())
			}
		

		}
		
	} else if t1!=nil && len(t2)==0{
		for i:=0;i<len(t1);i++{
			variable := p.GetArgs()[0].GetMetas()[i]
			present=false
			fmt.Println("variable1 =",variable.GetName())
			for i:=0;i<len(tab_variable);i++{
				if tab_variable[i]==variable.GetName(){
					present=true
				}
			}
			if !present{
				tab_variable=append(tab_variable,variable.GetName())
			}
		

		}

	}

	fmt.Println("t1 ", t1)
	fmt.Println("t2 ", t2)
	fmt.Println("tableau_variable ", tab_variable)
return tab_variable
}


func passe1greateq_lesseq(p types.Pred, t1 []string, t2 []string, cpt *int,tab_variable []string) ([]string){
			*cpt+=1
			present:=false
			if len(t1)==0 && t2!=nil{
				for i:=0;i<len(t2);i++{
					variable := p.GetArgs()[1].GetMetas()[i]
					present=false
					fmt.Println("variable2 =",variable.GetName())
					for i:=0;i<len(tab_variable);i++{
						if tab_variable[i]==variable.GetName(){
							present=true
						}
					}
					if !present{
						tab_variable=append(tab_variable,variable.GetName())
					}
				

				}
				
			} else if t1!=nil && len(t2)==0{
				for i:=0;i<len(t1);i++{
					variable := p.GetArgs()[0].GetMetas()[i]
					present=false
					fmt.Println("variable1 =",variable.GetName())
					for i:=0;i<len(tab_variable);i++{
						if tab_variable[i]==variable.GetName(){
							present=true
						}
					}
					if !present{
						tab_variable=append(tab_variable,variable.GetName())
					}
				

				}
			
			}

			fmt.Println("t1 ", t1)
			fmt.Println("t2 ", t2)
		return tab_variable
}


/** TODO
* C'est ici qu'on gère la conversion des variables
* Je fais beaucoup de disjonction de cas en fonction de int ou rat, mais selon votre format d'entrée ce ne sera peut-être pas nécessaire
**/
func termToSimplex(t types.Term, map_v_mv *map[string]types.Meta, iv *[]string) ([]string, error) {
	var tab_var = make([]string,0)
	switch ttype := t.(type) {
	case types.Meta:
		// C'est ici que je stock les metavairables, que je regarde si elles sont entière et que je fais la correspondence
		var_for_simplex := ttype.ToString()
		(*map_v_mv)[var_for_simplex] = ttype // Je stock la nouvelle variable de simplexe dans une map pour refaire le lien après
		if typing.IsInt(ttype.GetTypeHint()) {
			(*iv) = append((*iv), var_for_simplex) // Je stock aussi la variable dans la liste des variables entière si elle doit être entière
		}
		tab_var= append(tab_var,var_for_simplex)
		return tab_var, nil
	case types.Fun:
		
		
		switch t.GetName(){
		case "sum":
			var arg1, arg2 types.Term
			arg1 = ttype.GetArgs()[0]
			arg2 = ttype.GetArgs()[1]
			
			var1,_:=termToSimplex(arg1,map_v_mv,iv)
			var2,_:=termToSimplex(arg2,map_v_mv,iv)
			if var1!=nil{
				tab_var=append(tab_var,var1...)
			} 
			if var2!=nil{
				tab_var=append(tab_var,var2...)
			}

			return tab_var,nil
		case "product":

			var arg1, arg2 types.Term
			arg1 = ttype.GetArgs()[0]
			arg2 = ttype.GetArgs()[1]
			
			var1,_:=termToSimplex(arg1,map_v_mv,iv)
			var2,_:=termToSimplex(arg2,map_v_mv,iv)
			if var1!=nil{
				tab_var=append(tab_var,var1...)
			} 
			if var2!=nil{
				tab_var=append(tab_var,var2...)
			}

			return tab_var,nil
		
		case "difference":
			
			var arg1, arg2 types.Term
			arg1 = ttype.GetArgs()[0]
			arg2 = ttype.GetArgs()[1]
			
			var1,_:=termToSimplex(arg1,map_v_mv,iv)
			var2,_:=termToSimplex(arg2,map_v_mv,iv)
			if var1!=nil{
				tab_var=append(tab_var,var1...)
			} 
			if var2!=nil{
				tab_var=append(tab_var,var2...)
			}
			
			return tab_var,nil

		

		default:
			return tab_var,nil
		}
		//return tab_var, nil
	default:
		fmt.Printf("Unexpected type in termToSimplex\n")
		return tab_var, errors.New("Error")
	}
}

/**
* TODO
**/
func funToSimplex(f types.Fun, map_v_mv *map[string]types.Meta, iv *[]string) (string, error) {
	switch f.GetID().GetName() {
	case "sum":
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
