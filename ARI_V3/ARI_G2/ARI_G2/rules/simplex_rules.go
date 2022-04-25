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



type pair_coef_var struct{
	coef *big.Rat
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
	ligne_matrice:=0
	passe:=0
	var list_list_pcv = make([][]pair_coef_var,0)
	
	for passe<3{
			

		for number_of_predicate, p := range pl {
			var list_pcv = make([]pair_coef_var,0)
			
			t1, val1, err1 := termToSimplex(p.GetArgs()[0], &map_variable_metavariables, &int_variables,passe,ligne_matrice,list_pcv)

			t2, val2, err2 := termToSimplex(p.GetArgs()[1], &map_variable_metavariables, &int_variables,passe,ligne_matrice,list_pcv)
			if err1 != nil || err2 != nil  {
				fmt.Printf("Error in normalizeForSimplex")
				return nil, nil, nil
			}

		
			switch p.GetID().GetName() {
		
				case types.Id_eq.GetName():
					if passe==0{
						tab_variable=passe1(p,t1,t2,&ligne_matrice,tab_variable, true)
					}
					
					if passe==1{
						lenPl:=len(pl)
						list_list_pcv=passe2(list_list_pcv,number_of_predicate,list_pcv,true, val1,val2, lenPl)
					}

					


				case types.Id_neq.GetName():

				//à réfléchir, si on a 2x != 3 alors on a 2x > 3  OU  2x < 3
				//il faudrait donc faire 2 systèmes, un avec chacune des équations.. 
				//et encore, ça ne marche que si on cherche une solution entière..
				
				
				case "less":
					
					if passe==0{
						tab_variable=passe1(p,t1,t2,&ligne_matrice,tab_variable, false)
					}
				case "lesseq":
					
					if passe==0{
						tab_variable=passe1(p,t1,t2,&ligne_matrice,tab_variable, false)
					}
				case "great":
					
					if passe==0{
						tab_variable=passe1(p,t1,t2,&ligne_matrice,tab_variable, false)
					}

				case "greateq":
					if passe==0{
						tab_variable=passe1(p,t1,t2,&ligne_matrice,tab_variable, false)
					}
			}
			
		}
		passe+=1
	}

	return res_for_simplex, map_variable_metavariables, int_variables
}





func passe2(list_list_pcv [][]pair_coef_var,number_of_predicate int, list_pcv []pair_coef_var, eg bool,val1 []pair_coef_var,val2 []pair_coef_var, lenPl int )  [][]pair_coef_var{

		for i:=0;i<len(val1);i++{
			list_pcv=append(list_pcv,val1[i])
		}

		
		for i:=0; i<len(val2);i++{
			list_pcv=append(list_pcv,val2[i]) 
		}

		list_list_pcv=append(list_list_pcv,list_pcv)
	
	if eg{
		var list_pcv_eg = make([]pair_coef_var,0)
		for i:=0;i<len(val1);i++{
			egneg:=new(big.Rat).Mul(val1[i].coef,big.NewRat(-1,1))
			var pair pair_coef_var
			pair.variable=val1[i].variable
			pair.coef=egneg
			list_pcv_eg=append(list_pcv_eg,pair)
		}

		for i:=0; i<len(val2);i++{
			egneg2:=new(big.Rat).Mul(val2[i].coef,big.NewRat(-1,1))
			var pair2 pair_coef_var
			pair2.coef=egneg2
			pair2.variable=val2[i].variable
			list_pcv_eg=append(list_pcv_eg,pair2)
		}

		list_list_pcv=append(list_list_pcv,list_pcv_eg)
	
	}	
	if lenPl-1==number_of_predicate{		
		fmt.Println("list_list_pcv = ",list_list_pcv)
	}
	
	return list_list_pcv
}



func passe1(p types.Pred,t1 []string, t2 []string, cpt *int,tab_variable []string, b bool) ([]string){
	if b{
		*cpt+=2
	}else {
		*cpt+=1
	}
	present:=false
	if t2!=nil{
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
	}	
	if t1!=nil {
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

	fmt.Println("tab_var : ",tab_variable)
	fmt.Println("cpt = ",*cpt)

return tab_variable
}


/** TODO
* C'est ici qu'on gère la conversion des variables
* Je fais beaucoup de disjonction de cas en fonction de int ou rat, mais selon votre format d'entrée ce ne sera peut-être pas nécessaire
**/
func termToSimplex(t types.Term, map_v_mv *map[string]types.Meta, iv *[]string, passe int, cpt int, pcv []pair_coef_var) ([]string, []pair_coef_var, error) {

	var tab_var = make([]string,0)	
	switch ttype := t.(type) {
	case types.Meta:
		if passe==0{
			// C'est ici que je stock les metavairables, que je regarde si elles sont entière et que je fais la correspondence
			var_for_simplex := ttype.ToString()
			(*map_v_mv)[var_for_simplex] = ttype // Je stock la nouvelle variable de simplexe dans une map pour refaire le lien après
			if typing.IsInt(ttype.GetTypeHint()) {
				(*iv) = append((*iv), var_for_simplex) // Je stock aussi la variable dans la liste des variables entière si elle doit être entière
			}
			tab_var= append(tab_var,var_for_simplex)
			return tab_var,pcv, nil
		}
		if passe==1{
//			if arg_meta,ok:=ttype.(types.Meta);ok{
				var pair pair_coef_var
				pair.variable=ttype
				pair.coef=big.NewRat(1,1)
				pcv=append(pcv,pair)
				return tab_var,pcv,nil
//			}
		}
	case types.Fun:
		
		switch t.GetName(){
			case "sum":

				if passe==0{
		
					var arg1, arg2 types.Term
					arg1 = ttype.GetArgs()[0]
					arg2 = ttype.GetArgs()[1]
					
					var1,_,_:=termToSimplex(arg1,map_v_mv,iv,passe,cpt,pcv)
					var2,_,_:=termToSimplex(arg2,map_v_mv,iv,passe,cpt,pcv)
					if var1!=nil{
						tab_var=append(tab_var,var1...)
					} 
					if var2!=nil{
						tab_var=append(tab_var,var2...)
					}
					

					return tab_var,pcv,nil
				}

				if passe==1{			
					pair,_,_:=funToSimplex(ttype,map_v_mv,iv,cpt,pcv)
					return tab_var,pair,nil 
				}
			case "product":
				if passe==0{
		
					var arg1, arg2 types.Term
					arg1 = ttype.GetArgs()[0]
					arg2 = ttype.GetArgs()[1]
					
					var1,_,_:=termToSimplex(arg1,map_v_mv,iv,passe,cpt,pcv)
					var2,_,_:=termToSimplex(arg2,map_v_mv,iv,passe,cpt,pcv)
					if var1!=nil{
						tab_var=append(tab_var,var1...)
					} 
					if var2!=nil{
						tab_var=append(tab_var,var2...)
					}
					

					return tab_var,pcv,nil
				}
				if passe==1{			
					pair,_,_:=funToSimplex(ttype,map_v_mv,iv,cpt,pcv)
					fmt.Println("pairT = ",pair[0].coef)
					return tab_var,pair,nil 
				}
			case "difference":
				if passe==0{
	
					var arg1, arg2 types.Term
					arg1 = ttype.GetArgs()[0]
					arg2 = ttype.GetArgs()[1]
					
					var1,_,_:=termToSimplex(arg1,map_v_mv,iv,passe,cpt,pcv)
					var2,_,_:=termToSimplex(arg2,map_v_mv,iv,passe,cpt,pcv)
					if var1!=nil{
						tab_var=append(tab_var,var1...)
					} 
					if var2!=nil{
						tab_var=append(tab_var,var2...)
					}
					
					return tab_var,pcv,nil

				}		
				if passe==1{			
					pair,_,_:=funToSimplex(ttype,map_v_mv,iv,cpt,pcv)
					return tab_var,pair,nil 
				}

			default:
				var pair pair_coef_var
				monRat,_:=new(big.Rat).SetString(t.GetName())
				pair.coef=monRat
				pcv=append(pcv,pair)
				return tab_var,pcv,nil
		}
	default:
		fmt.Printf("Unexpected type in termToSimplex\n")
		return tab_var,pcv, errors.New("Error")
	}
	return tab_var,pcv,nil

}

/**
* TODO
**/


func funToSimplex(f types.Fun, map_v_mv *map[string]types.Meta, iv *[]string,cpt int,pcv []pair_coef_var) ([]pair_coef_var,*big.Rat, error) {

	
	switch f.GetID().GetName() {
	
	case "sum":
		arg1 := f.GetArgs()[0]
		arg2 := f.GetArgs()[1]
		

		if arg1.IsFun() && arg2.IsMeta(){
			if arg_fun, ok := arg1.(types.Fun); ok{
				var pair pair_coef_var
				if arg_meta,ok2:=arg2.(types.Meta);ok2{
					pair.variable=arg_meta
					pair.coef=big.NewRat(1,1)
					pcv=append(pcv,pair)
					return funToSimplex(arg_fun,map_v_mv,iv,cpt,pcv)
				}	
			}
		}

		if arg2.IsFun() && arg1.IsMeta(){
			if arg_fun, ok := arg2.(types.Fun); ok{
				var pair pair_coef_var
				if arg_meta,ok2:=arg1.(types.Meta);ok2{
					pair.variable=arg_meta
					pair.coef=big.NewRat(1,1)
					pcv=append(pcv,pair)
					return funToSimplex(arg_fun,map_v_mv,iv,cpt,pcv)
				}
			}
		}

		if arg1.IsMeta() && arg2.IsMeta(){
			if arg_meta1,ok1:=arg1.(types.Meta);ok1{
				if arg_meta2,ok2:=arg2.(types.Meta);ok2{
					var pair1, pair2 pair_coef_var
					pair1.variable=arg_meta1
					pair1.coef=big.NewRat(1,1)
					pair2.variable=arg_meta2
					pair2.coef=big.NewRat(1,1)
					pcv=append(pcv,pair1)
					pcv=append(pcv,pair2)
					return pcv, newRat(),nil
				}
			}
		}


		if arg1.IsFun() && arg2.IsFun(){
			if arg_fun2, ok2 := arg2.(types.Fun); ok2 {
				if arg_fun1, ok1 := arg1.(types.Fun); ok1{
					if !checkBinaryArithmeticFun(arg_fun1.GetID() ) && ! checkBinaryArithmeticFun(arg_fun2.GetID() ){
						test_rat,_:=EvaluateFun(f)
						return pcv,test_rat,nil
					}
				}
			}
		}

	case "difference":

		fmt.Println("ici2")
		arg1 := f.GetArgs()[0]
		arg2 := f.GetArgs()[1]


		if arg1.IsFun() && arg2.IsMeta(){
			fmt.Println("here2")
			if arg_fun, ok := arg1.(types.Fun); ok{
				if arg_meta,ok2:=arg2.(types.Meta);ok2{
					var pair pair_coef_var
					pair.variable=arg_meta
					pair.coef=big.NewRat(1,1)
					pcv=append(pcv,pair)
					return funToSimplex(arg_fun,map_v_mv,iv,cpt,pcv)
				}
			}
		}


		if arg2.IsFun() && arg1.IsMeta(){
			if arg_fun, ok := arg2.(types.Fun); ok{
				if arg_meta,ok2:=arg1.(types.Meta);ok2{
					var pair pair_coef_var
					pair.variable=arg_meta
					pair.coef=big.NewRat(1,1)
					pcv=append(pcv,pair)
					return funToSimplex(arg_fun,map_v_mv,iv,cpt,pcv)
				}
			}
		}



		if arg1.IsMeta() && arg2.IsMeta(){
			if arg_meta1,ok1:=arg1.(types.Meta);ok1{
				if arg_meta2,ok2:=arg2.(types.Meta);ok2{			
					var pair1, pair2 pair_coef_var
					pair1.variable=arg_meta1
					pair1.coef=big.NewRat(1,1)
					pair2.variable=arg_meta2
					pair2.coef=big.NewRat(1,1)
					pcv=append(pcv,pair1)
					pcv=append(pcv,pair2)
					return pcv, newRat(),nil
				}
			}
		}



		if arg1.IsFun() && arg2.IsFun(){
			if arg_fun2, ok2 := arg2.(types.Fun); ok2 {
				if arg_fun1, ok1 := arg1.(types.Fun); ok1{
					if !checkBinaryArithmeticFun(arg_fun1.GetID() ) && ! checkBinaryArithmeticFun(arg_fun2.GetID() ){
						test_rat,_:=EvaluateFun(f)
						return pcv,test_rat,nil
					}
				}
			}
		}
	
	case "product":
		arg1 := f.GetArgs()[0]
		arg2 := f.GetArgs()[1]


		if arg1.IsFun() && arg2.IsMeta(){
			if arg_fun, ok := arg1.(types.Fun); ok{
				if arg_meta,ok2:=arg2.(types.Meta);ok2{
					_,val,_:=funToSimplex(arg_fun,map_v_mv,iv,cpt,pcv)
					var pair pair_coef_var
					pair.variable=arg_meta
					pair.coef=val
					pcv=append(pcv,pair)
					return pcv, newRat(),nil
				}
			}
			
		}


		if arg2.IsFun() && arg1.IsMeta(){
			if arg_fun, ok := arg2.(types.Fun); ok{
				if arg_meta,ok2:=arg1.(types.Meta);ok2{
					_,val,_:=funToSimplex(arg_fun,map_v_mv,iv,cpt,pcv)
					var pair pair_coef_var
					pair.variable=arg_meta
					pair.coef=val
					pcv=append(pcv,pair)
					return pcv, newRat(),nil
				}
			}
		}

		if arg1.IsFun() && arg2.IsFun(){
			if arg_fun2, ok2 := arg2.(types.Fun); ok2 {
				if arg_fun1, ok1 := arg1.(types.Fun); ok1{
					if !checkBinaryArithmeticFun(arg_fun1.GetID() ) && ! checkBinaryArithmeticFun(arg_fun2.GetID() ){
						test_rat,_:=EvaluateFun(f)
						return pcv,test_rat,nil
					}
				}
			}
		}


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
		return pcv,newRat(), errors.New("Error")
	}
	return pcv,newRat(), errors.New("Error")
}