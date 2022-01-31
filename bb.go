package main

import (
	"fmt"
	"math"
)

func main() {
	var tableau = [][]float64{{1,1}, {2,-1}, {-1,2}}
	var tabConst = []float64{2,0,1}

	branch_bound(simplexe(tableau, tabConst), tableau, tabConst)
}

func branch_bound(solution []float64, gotSol bool, tableau [][]float64, tabConst []float64) []flaot64{

	//TODO: Cas d'arret si on a pas de solution
	//Comment savoir qu'on a pas de solution ?

	//Cas d'arret si solution est fait seulement d'entier
	if(estSol(solution)){
		return solution
	}
	for index, element := range solution {
		if(!isInteger(element)){
			for i := 0; i < 2; i++ {
				var tableauBis []float64
				var tabConstBis []float64

				//Copie de tableau et du tableau de contrainte
				for i := 0; i < len(tabConst); i++ {
					tabConstBis = append(tab, tabConst[i])
				}
				for i := 0; i < len(tableau); i++ {
					tableauBis = append(tab, tableau[i])
				}

				//Ajout de la nouvelle contrainte dans les copies de tableau
				if i==0 {
					//Traduction tableau et contrainte de Margaux
				} else {
					//Traduction tableau et contrainte de Margaux
				}

				return go branch_bound(simplexe(tableauBis, tabConstBis), tableauBis, tabConstBis)
			}
		}	
	}	
}

//Verifie que le nombre donné soit un entier
func isInteger(nombre float64) bool{
	return 	math.Mod(nombre,1) == 0

}

//Verifie qu'un tableau contient seulement des entier
func estSol(solution []float64) bool{
	for _, element := range solution {
		if(!isInteger(element)){
			return false
		}
	}
	return true
}

