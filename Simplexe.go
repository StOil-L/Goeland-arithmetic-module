package main

import (
	"fmt"
)

func main() {
	var tableau = [][]float64{{1,1}, {2,-1}, {-1,2}}
	var constTab = []float64{2,0,1}
	fmt.Println(simplexe(tableau, constTab))
}

func simplexe(tableau [][]float64, tabConst []float64) [][]float64{
	var alphaTab [len(tableau) + len(tableau[0])]float64
	//for a modifier
	for k := 0; k < len(tableau); k++ {
		actualConstraint := checkConst(&alphaTab, tabConst, len(tableau[0]))
		if actualConstraint == nil {
			break
		}
		colonnePivot := pivot(tableau, tabConst, alphaTab, actualConstraint)
		if colonnePivot == nil {
			fmt.Println("Il n'existe pas de solution pour ces contraintes") 
		} else {
			for i := 0; i < len(tableau[0]); i++ {
				if i == colonnePivot {
					tableau[actualConstraint][i] = 1/tableau[actualConstraint][i]
				} else {
					tableau[actualConstraint][i] = -tableau[actualConstraint][i]/tableau[actualConstraint][colonnePivot]
				}
			}
			for i := 0; i < len(tableau); i++ {
				if i != actualConstraint {
					for j := 0; j < len(tableau[0]); j++ {
						tableau[i][j] += tableau[actualConstraint][j] * tableau[i][colonnePivot]
					}
				}
			}
			tabConst[actualConstraint] = nil
			for i := len(tableau[0]); i < len(alphaTab); i++ {
				if i !=  actualConstraint + len(tableau[0]){
					alphaTab[i] = 0
					for j := 0; j < len(tableau[0]); j++ {
						alphaTab[i] += tableau[i - len(tableau[0])][j] * alphaTab[j]
					}
				}
			}
		}
	}
	return tableau
}

//Cherche la première contrainte qui n'est pas respecté ca place dans tabConst
func checkConst(alphaTab []float64,  tabConst []float64, nbInconnu int) int{
	for index, element := range tabConst {
		if element != nil {
			if alphaTab[index + nbInconnu] < element {
				return index
			}
		}
	}
	return nil
}

//Renvoie la colonne pivot par rapport a la contrainte a traiter
func pivot(tableau []float64,  tabConst []float64, alphaTab *[]float64, pivotLane int) int{
	var teta float64
	for index1, element1 := range tableau[pivotLane]{
		teta = (tabConst[pivotLane] - (*alphaTab[pivotLane]) + len(tableau[0])) / element1
		alphaInconnu := teta + (*alphaTab[index1])
		var varAlphaEcart float64
		for index2, element2 := range tableau[pivotLane] {
			varAlphaEcart += element2 * (*alphaTab[index2])
		}
		if  varAlphaEcart >= tabConst[pivotLane] {
			//Changement des place dans alphaTab entre li pivot et la variable d'ecart
			*alphaTab[index1 + len(tableau[0])] = alphaInconnu
			*alphaTab[index1] = varAlphaEcart
			return index1
		}
	}
	return nil	 
}