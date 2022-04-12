package main

import (
	"fmt"
	"strconv"
    "strings"
    "regexp"
	"bufio"
    "os"
	"math/big"
	"math"
	"sync"
//	"time"
    
)

/** 
 * This function applies the main logic of the simplex algorithm (without objective function).
 * It takes the following parameters:
 *   - `tab_coef`, the matrice with the normalized inequations
 *   - `tab_cont`, an array containing the constraints, tab_cont[0] contains the constraint of the first line of the matrice
 *   - `tab_nom_var`, an array of the system's starting variable
 *   - `incremental_coef`,
 *	 - `incremental_aff`,
 * 	 - `pos_var_tab`, an array containing the variable positions in the matrice starting by the out-base variable
 * 	 - `bland`, an array containing the Bland order of variable
 * 	 - `alpha_tab`, a map associating the name of the variable and his alpha value
 * It returns an affectation that satisfies the constraints. 
 * In case the system has no solutions, the boolean return value is set to false.
 **/
func Simplexe(tab_coef [][]*big.Rat, tab_cont []*big.Rat, tab_nom_var[]string, incremental_coef[]*big.Rat,
	incremental_aff[]*big.Rat, pos_var_tab[]string, bland[]string, alpha_tab map[string]*big.Rat) 
	(bool, []string, []int){

	fmt.Println("tab_cont",tab_cont)
	fmt.Println("alpha_tab",alpha_tab)
	fmt.Println("tab_coef",tab_coef)

	var pos_var_tab_bis =make([]string,len(pos_var_tab))
	for i:=0;i<len(pos_var_tab);i++{
		pos_var_tab_bis[i] = pos_var_tab[i]
	}
	
	
	if len(tab_coef)+len(tab_coef[0]) != len(bland){
		bland=append(bland,fmt.Sprint("e",len(tab_coef)-1))
		pos_var_tab_bis = append(pos_var_tab_bis,fmt.Sprint("e",len(tab_coef)-1))
		for cpt := 0; cpt < len(tab_coef[0]); cpt++{
			tmp := pos_var_tab_bis[len(pos_var_tab)-cpt]
			pos_var_tab_bis[len(pos_var_tab_bis)-1-cpt]=pos_var_tab_bis[len(pos_var_tab_bis)-cpt-2]
			pos_var_tab_bis[len(pos_var_tab_bis)-cpt-2]=tmp
		}

	}
	fmt.Println("pos_var_tab_bis : ",pos_var_tab_bis)
	//incrémental aff ? 

	fmt.Println("\033[0m") 
	//boucle sur le nombre maximum de pivotation que l'on peut avoir
	for true {
		//ligne_pivot est la ligne qui ne respecte pas sa contrainte
		ligne_pivot := checkCont(alpha_tab, tab_cont, pos_var_tab_bis)		
		if ligne_pivot == -1 {
			fmt.Println(" \033[33m La solution est : ") 
			fmt.Println(alpha_tab)
			return true, pos_var_tab_bis
		}
		//on cherche la colonne du pivot
		colonne_pivot := pivot(tab_coef, tab_cont, alpha_tab, ligne_pivot,
			 pos_var_tab_bis, bland)
		if colonne_pivot == -1 {
			fmt.Println(" \033[33m") 
			fmt.Println("Il n'existe pas de solution pour ces contraintes")
			return false, pos_var_tab_bis
		} else {
			//on modifie le tableau des coefficients pour la ligne du pivot
			updateMatrice(tab_coef,colonne_pivot,ligne_pivot,incremental_coef)
			//calcul des nouveaux alpha
			updateAlpha(tab_coef,ligne_pivot,alpha_tab,pos_var_tab_bis,incremental_aff)
			//time.Sleep(time.Second)
			fmt.Println("\033[35m matrice des coefficients :",tab_coef,"\033[0m")
			fmt.Println("\033[34m affectations :" ,alpha_tab,"\033[0m")
			
		}
	}
}


/**
 * This function search the first constraint which isn't respected
 * It takes the following parameters:
 *   - `alpha_tab`, a map associating the name of the variable and his alpha value
 *   - `tab_cont`, an array containing the constraints, tab_cont[0] contains the constraint of the first line of the matrice
 * 	 - `pos_var_tab`, an array containing the variable positions in the matrice starting by the out-base variable
 * It return the `tab_coef`'s line where the constraints isn't respect
 * If all constraints are respected, return -1
 **/
func checkCont(alpha_tab map[string]*big.Rat,  tab_cont []*big.Rat, pos_var_tab []string) int{
	for index, variable := range pos_var_tab[:len(tab_cont)]  {
		if variable[0] == 101 {
			if alpha_tab[variable].Cmp(tab_cont[index]) == -1 {
				return index
			}	
		}	
	}
	return -1
}

func checkContHorsBase(tab_cont []*big.Rat, pos_var_tab []string, nom_var string) bool{
	for index, variable := range pos_var_tab[:len(tab_cont)]  {
		if variable == nom_var {
			return true
		}
	}
	return false
}


/**
 * This function search the pivot and update the alpha and position of the switched varaible
 * It takes the following parameters:
 *   - `tab_coef`, the matrice with the normalized inequations
 *   - `tab_cont`, an array containing the constraints, tab_cont[0] contains the constraint of the first line of the matrice
 * 	- `alpha_tab`, a map associating the name of the variable and his alpha value
 * 	- `pivot_line`, an int which indacate the pivot line in `tab_coef`
 * 	- `pos_var_tab`, an array containing the variable positions in the matrice starting by the out-base variable
 *	- `bland`, an array containing the Bland order of variable
 * It return the line of the pivot
 * If any pivot is found, return -1
 **/
func pivot(tab_coef [][]*big.Rat,  tab_cont []*big.Rat,
	 alpha_tab map[string]*big.Rat, pivot_line int, pos_var_tab []string,
	  bland []string) int{

	var var_pivot string
	var colonne_pivot int
	for _,vari := range bland{
		var coef_colonne = new(big.Rat)
		//Find a possible pivot with Bland order coef and position
		for j:=len(tab_coef); j<len(tab_coef)+len(tab_coef[0]); j++ {
			if vari == pos_var_tab[j] {
				var_pivot = vari
				colonne_pivot = j
				coef_colonne = tab_coef[pivot_line][colonne_pivot-len(tab_coef)]
			}
		}

		var vide string
		var numero_variable_pivot int	
		if var_pivot != vide {
			numero_variable_pivot, _ := strconv.Atoi(var_pivot[1:])
		}

		var var_ecart int
		if var_pivot[0] == 'e' {
			var_ecart = var_pivot
		} else {
			var_ecart = pos_var_tab[pivot_line]
		}
		num_var_ecart := strconv.Atoi(var_ecart[1:])

		//Check if the pivot is suitable
		if colonne_pivot < len(tab_coef)+len(tab_coef[0])  && (coef_colonne.Cmp(new(big.Rat))!=0) && var_pivot!=vide 
		&& (var_pivot[0] !='e' || coef_colonne.Cmp(new(big.Rat))==1 ||
		(coef_colonne.Cmp(new(big.Rat))==-1 && checkContHorsBase(tab_cont, pos_var_tab, var_ecart)>-1 
		&& alpha_tab[var_pivot].Cmp(tab_cont[num_var_ecart])>0))  {
		//	 time.Sleep(time.Second)
			var theta = new(big.Rat)
			theta.Mul(new(big.Rat).Add(tab_cont[num_var_ecart], new(big.Rat).Neg(alpha_tab[pos_var_tab[pivot_line]])), 
				new(big.Rat).Inv(coef_colonne))
			var alpha_colonne = new(big.Rat)	
			alpha_colonne.Add(alpha_tab[var_pivot], theta)
			alpha_tab[var_pivot]=alpha_colonne
			var alpha_ligne = new(big.Rat)
			//on calcule alpha_ligne
			for index2, element2 := range tab_coef[pivot_line] {
					alpha_ligne.Add(alpha_ligne, new(big.Rat).Mul(element2, alpha_tab[pos_var_tab[index2+len(tab_coef)]]))
			}
			fmt.Println("alpha_colonne", alpha_colonne)
			fmt.Println("alpha_ligne",alpha_ligne)
			alpha_tab[posVarTableau[pivot_line]].Set(alpha_ligne)
			alpha_tab[variablePivot].Set(alpha_colonne)
			fmt.Println("\033[0m variable \033[36m colonne:",var_pivot+"\033[0m","variable \033[36m ligne:",
			posVarTableau[pivot_line]+"\033[0m")
			pos_var_tab[pivot_line], pos_var_tab[colonne_pivot] = pos_var_tab[colonne_pivot], pos_var_tab[pivot_line]
			fmt.Println("\033[36m theta\033[0m =\033[36m",theta,"\033[0m")
			return colonne_pivot-len(tab_coef)
		}
	}
	return -1	 
}

/** 
 * This function update tab_coef.
 * It takes the following parameters:
 *   - `tab_coef`, the matrice with the normalized inequations
 *   - `colonne_pivot`, column of the matrice where the pivot occur
 *   - `ligne_pivot`, line of the matrice where the pivot occur
 *   - `incremental_coef`,
 **/
func updateMatrice(tab_coef [][]*big.Rat, colonne_pivot int, ligne_pivot int, incremental_coef []*big.Rat) {
	
	//ajout numéro colonne pivot 
	incremental_coef=append(incremental_coef,big.NewRat(int64(colonne_pivot), 1))
	//ajout pivot
	var inv_pivot = new(big.Rat)
	inv_pivot.Set(new(big.Rat).Inv(tab_coef[ligne_pivot][colonne_pivot]))
	for i := 0; i < len(tab_coef[0]); i++ {
		if i == colonne_pivot {
			incremental_coef=append(incremental_coef,inv_pivot)
			tab_coef[ligne_pivot][i]=inv_pivot
		} else {
			tab_coef[ligne_pivot][i].Mul(new(big.Rat).Neg(tab_coef[ligne_pivot][i]), inv_pivot)
			incremental_coef=append(incremental_coef,tab_coef[ligne_pivot][i])
		}

	}
		
	//on modifie le tableau des coefficients des autres lignes
	for i := 0; i < len(tab_coef); i++ {
		//conservation du coefficient non modifier ligne actuel/colonne pivot
		//necessaire pour les prochains calcul de coef de la ligne
		var  tab_i_pivot=new(big.Rat)
		tab_i_pivot.Set(tab_coef[i][colonne_pivot])
	
		if i != ligne_pivot {
			for j := 0; j < len(tab_coef[0]); j++ {
				if j==colonne_pivot{
					tab_coef[i][colonne_pivot].Mul(tab_coef[i][colonne_pivot], tab_coef[ligne_pivot][colonne_pivot])
				} else {
					tab_coef[i][j].Add(tab_coef[i][j], new(big.Rat).Mul(tab_coef[ligne_pivot][j],tab_i_pivot))
		
				}
			}
		}
	}
}

/** 
 * This function update alpha_tab.
 * It takes the following parameters:
 *   - `tab_coef`, the matrice with the normalized inequations
 *   - `ligne_pivot`, line of the matrice where the pivot occur
 * 	 - `alpha_tab`, a map associating the name of the variable and his alpha value
 * 	 - `pos_var_tab`, an array containing the variable positions in the matrice starting by the out-base variable
 *   - `incremental_aff`,
 **/
func updateAlpha(tab_coef [][]*big.Rat, ligne_pivot int, 
	alpha_tab map[string]*big.Rat, pos_var_tab []string, incremental_aff []*big.Rat){
	for i := 0; i<len(tab_coef);i++{
		if i != ligne_pivot {
			var cal_alpha = new(big.Rat)
			for j :=0; j<len(tab_coef[0]);j++{
				cal_alpha.Add(cal_alpha, new(big.Rat).Mul(tab_coef[i][j], alpha_tab[pos_var_tab[j + len(tab_coef)]]))
			}
			alpha_tab[pos_var_tab[i]].Set(cal_alpha)
		}
	}
	for j :=0; j<len(tab_coef[0]);j++{
		incremental_aff=append(incremental_aff,alpha_tab[pos_var_tab[j + len(tab_coef)]])
	}
}

/** 
 * This function transform a system of equations on different varaible needed for the simplex
 * It takes the following parameters:
 *   - `eqs`, a string which represent a system of equations
 * It returns the matrice of coefficient, the array of constraint and table of the system's starting variable
 **/
func AddAllConst(eqs []string, tab_coef [][]*big.Rat, tab_cont []*big.Rat, tab_nom_var []string) {
    for _, element := range eqs {
        dernier_element, tab, liste_nom_var := parseOneConst(element)
        tab_coef = append(tab_coef, tab)
        tab_cont = append(tab_cont, dernier_element)
        addVarIfNotExists(liste_nom_var, tab_nom_var)
    }
}

/** 
 * This function add the element of an array of string that doesn't exist in the other one
 * It takes the following parameters:
 *   - `liste_nom_var`, an array of string that will give his new element
 * 	 - `tab_nom_var`, an array of string that will receive the new element
 **/
func addVarIfNotExists(liste_nom_var []string, tab_nom_var []string) {
	for i := 0; i < len(liste_nom_var); i++ {
		present := false
		for j := 0; j < len(tab_nom_var); j++ {
			if liste_nom_var[i] == tab_nom_var[j] {
				present = true
			}
			if !present {
				tab_nom_var = append(tab_nom_var, liste_nom_var[i])
			}
		}
	}
}

/** 
 * This function transform one equations on different varaible needed for the simplex
 * It takes the following parameters:
 *   - `eq`, a string which represent one equation
 * It returns a line of the matrice of coefficient, the constraint of the equation and the table of the equation's starting variable
 **/
func parseOneConst(eq string) (*big.Rat, []*big.Rat,[]string){
    // on split la string avec les espaces, ce qui nous donne un tableaux avec tous les elements de la string
    tab_element := strings.Split(eq, " ")
    // tableau qui va contenir les coefficients de l'equation
    var ligneEq []*big.Rat
    var tab_nom_var []string
    // indique notre position dans le tableau
    posTab := 0
    for i := 0; i < len(tab_element)-2; i++ {
        ligneEq = append(ligneEq, big.NewRat(1,1))
        // nous permet de savoir si notre caractere est un chiffre
        re := regexp.MustCompile(`[0-9.,]`)
        isFig := re.FindString(tab_element[i])
        // nous permet de savoir si notre caractere est une lettre
        re2 := regexp.MustCompile(`[a-z]`)
        isLet := re2.FindString(tab_element[i])

        if isLet != "" {
          tab_nom_var = append(tab_nom_var,isLet)
        }       


        // si on a un chiffre
        if isFig != "" {
            // on converti notre caractere qui est une string en float64
            conv,_ := new(big.Rat).SetString(tab_element[i])
           // on l'insere dans notre tableau a la position "posTab"
            ligneEq[posTab]=conv

            
        } else {
            // le cas du -
            if tab_element[i] == "-"{
                ligneEq[posTab].SetFloat64(-1.0)
            // le cas du +
            } else if tab_element[i] != "+" {
                posTab += 1
            }
        }
    }

    // On ajoute la contraintes dans le tableau de contraintes en l'a convertissant d'abord
    lastEle := tab_element[len(tab_element)-1]
	lastEleC,_ := new(big.Rat).SetString(lastEle)
 
    return lastEleC, ligneEq[0:posTab],tab_nom_var
}

