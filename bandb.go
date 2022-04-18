package main

import (
	"fmt"
	"math/big"
	"math"  
)

/** 
 * This function applies the main logic of the branch and bound algorithm
 * It takes the following parameters:
 *	 - `solution`, solution of the system of equation return by the simplex in other way, it's `alpah_tab`
 *	 - `gotSol`, a boolean return by the simplex saying if the simplex find a solution
 *   - `tab_nom_var`, an array of the system's starting variable
 *   - `tab_coef`, the matrice with the normalized inequations
 *   - `tab_cont`, an array containing the constraints, tab_cont[0] contains the constraint of the first line of the matrice
 *   - `channel`, 
 *   - `incremental_coef`,
 *	 - `incremental_aff`,
 * 	 - `pos_var_tab`, an array containing the variable positions in the matrice starting by the out-base variable
 * 	 - `bland`, an array containing the Bland order of variable
 * It returns a boolean which says if it exist a solution and the `alpha_tab` which represent the solution of the system
 **/
func Branch_bound(gotSol bool, channel chan bAndB, system info_system) (map[string]*big.Rat, bool){

	fmt.Println("\033[0m ") 
	
	solutionEntiere,index:=estSol(system.alpha_tab,system.tab_nom_var)
	
	//Cas d'arret si solution est fait seulement d'entier
	if (!gotSol) {
        return system.alpha_tab, false
    } else if (solutionEntiere){
        return system.alpha_tab, true
    }
	go goBandB(false, channel, index, system)
	go goBandB(true, channel, index, system)
	
	stBAndB := <- channel
	if(!stBAndB.solBoolStr){
		stBAndB = <- channel
		if(stBAndB.solBoolStr){
			close(channel)
		}
	} else {
		select{
			case <- channel :
			
			default :
				close(channel)
		}
	}
    return stBAndB.solStr, stBAndB.solBoolStr
}

func goBandB(inf_sup bool, channel chan bAndB, index int, system info_system) {
	select {
		case <- channel :
			return
		default :
			var tab_coef_bis [][]*big.Rat
			var tab_cont_bis []*big.Rat
			channelBis := make(chan bAndB)
			//Copie de tableau et du tableau de contrainte
			tab_cont_bis = deepCopyTableau(system.tab_cont)
			tab_coef_bis = deepCopyMatrice(system.tab_coef)
			//Ajout de la nouvelle contrainte dans les copies de tableau
			if inf_sup {
				partiEntiere, _ := system.alpha_tab[system.tab_nom_var[index]].Float64()
				tab_cont_bis = append(tab_cont_bis, new(big.Rat).SetFloat64(math.Ceil(partiEntiere)))
				var tabInter []*big.Rat
				for i := 0; i < len(system.tab_nom_var); i++ {
					if i == index {
						tabInter = append(tabInter, big.NewRat(1,1))
					}else {
						tabInter = append(tabInter, new(big.Rat))
					}
				}
				tab_coef_bis = append(tab_coef_bis, tabInter)
			} else  {
				var tabInter []*big.Rat
				partiEntiere, _ := system.alpha_tab[system.tab_nom_var[index]].Float64()
				tab_cont_bis = append(tab_cont_bis, new(big.Rat).SetFloat64(-math.Floor(partiEntiere)))
				for i := 0; i < len(system.tab_nom_var); i++ {
					if i == index {
						tabInter = append(tabInter, big.NewRat(-1,1))
					}else {
						tabInter = append(tabInter, new(big.Rat))
					}
				}
				tab_coef_bis = append(tab_coef_bis, tabInter)
			}
				
			system.tab_coef = tab_coef_bis
			system.tab_cont = tab_cont_bis

			//incrémental
			system = incremental(system) 
			//fin incrémental

			system, gotSol := Simplexe(system)

			system.tab_coef = tab_coef_bis
			system.tab_cont = tab_cont_bis

			sol, solBool := Branch_bound(gotSol, channelBis, system)

			stBAndB := bAndB{solBoolStr: solBool, solStr: sol}
			select {
				case channel <- stBAndB:
				case <- channel:
			}
	}
}

/** 
 * This function check if the solution is natural
 * It takes the following parameters:
 * 	 - `solution`, a map associating the name of the variable and his alpha value
 *   - `tab_nom_var`, an array of the system's starting variable
 * It returns a boolean which inform of the solution is natural and the index of the decimal number in `solution` if there is one
 **/
func estSol(solution map[string]*big.Rat, tab_nom_var []string) (bool,int){
	index:=0
	for ( index < len(tab_nom_var) && solution[tab_nom_var[index]].IsInt()){
		index+=1
	}
	if index < len(tab_nom_var){
			return false,index
	}
	fmt.Printf("test %d \n", len(tab_nom_var))
	return true,index
}

/** 
 * This function make a deep copy of the matrice give in parameter
 * It takes the following parameters:
 * 	 - `tab`, matrice that will be copied
 * It returns the copied matrice
 **/
func deepCopyMatrice(tab [][]*big.Rat) [][]*big.Rat {
	var tab_copy = make([][]*big.Rat,len(tab))
	for index_tab_ligne := 0 ; index_tab_ligne<len(tab) ; index_tab_ligne++{
		tab_copy[index_tab_ligne] = append(tab_copy[index_tab_ligne], deepCopyTableau(tab[index_tab_ligne])...)
	}
	return tab_copy
}

/** 
 * This function make a deep copy of the array give in parameter
 * It takes the following parameters:
 * 	 - `tab`, array that will be copied
 * It returns the copied array
 **/
func deepCopyTableau(tab []*big.Rat) []*big.Rat {
	var tab_copy =make([]*big.Rat,len(tab))
		for index_tab_colonne := 0 ; index_tab_colonne<len(tab) ; index_tab_colonne++{
			tab_copy[index_tab_colonne] = new(big.Rat)
			tab_copy[index_tab_colonne].Set(tab[index_tab_colonne])
		}
	return tab_copy
}

func incremental(system info_system) (info_system){
	alpha_tab_bis := make(map[string]*big.Rat)		
			cpt:=0
			cpt2:=0
			for cpt < len(system.incremental_coef){		
				var case_pivot =new(big.Rat)
				case_pivot.Set(system.tab_coef[len(system.tab_coef)-1][system.incremental_coef[cpt].Num().Int64()])
				for j := 0; j < len(system.tab_coef[0]); j++ {				
					if int64(j)==system.incremental_coef[cpt].Num().Int64(){
						system.tab_coef[len(system.tab_coef)-1][j].Mul(system.tab_coef[len(system.tab_coef)-1][j],
						system.incremental_coef[int64(cpt)+system.incremental_coef[cpt].Num().Int64()+1])
					} else {
						system.tab_coef[len(system.tab_coef)-1][j].Add(system.tab_coef[len(system.tab_coef)-1][j],
						new(big.Rat).Mul(case_pivot,system.incremental_coef[j+1]))			
					}

				}		
				var calAlpha = new(big.Rat)
				for j :=0; j<len(system.tab_coef[0]);j++{
					calAlpha.Add(calAlpha,new(big.Rat).Mul(system.incremental_aff[j+cpt2],system.tab_coef[len(system.tab_coef)-1][j]))
				}


				for i := 0; i < len(system.tab_coef)-1; i++ {
					alpha_tab_bis[fmt.Sprint("e", i)] = new(big.Rat)
					alpha_tab_bis[fmt.Sprint("e", i)].Set(system.alpha_tab[fmt.Sprint("e", i)]) 
				}
				alpha_tab_bis[fmt.Sprint("e", len(system.tab_coef)-1)]=new(big.Rat)
				alpha_tab_bis[fmt.Sprint("e", len(system.tab_coef)-1)].Set(calAlpha)
				if len(system.tab_nom_var) == 0 {
					for i := 0; i < len(system.tab_coef[0]); i++ {
					alpha_tab_bis[fmt.Sprint("x", i)] = new(big.Rat)
					alpha_tab_bis[fmt.Sprint("x", i)].Set(system.alpha_tab[fmt.Sprint("x", i)])
					}
				} else {
					for i := 0; i < len(system.tab_coef[0]); i++ {
						alpha_tab_bis[system.tab_nom_var[i]] = new(big.Rat)
						alpha_tab_bis[system.tab_nom_var[i]].Set(system.alpha_tab[system.tab_nom_var[i]])
					
					}
				}

				cpt+=1+len(system.tab_coef[0])
				cpt2+=len(system.tab_coef[0])
			}
			return system
}