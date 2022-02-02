package main

import (
	"fmt"
	"math"
)

func main() {
	var tableau = [][]float64{{1,1}, {2,-1}, {-1,2}}
	var tabConst = []float64{2,0,1}
	channel := make(chan []float64)
	branch_bound(simplexe(tableau, tabConst), tableau, tabConst, channel)
}

func branch_bound(solution []float64, gotSol bool, tableau [][]float64, tabConst []float64, channel chan []float64) []flaot64{

	//Cas d'arret si solution est fait seulement d'entier
	if(estSol(solution)){
		return solution
	}
	for index, element := range solution {
		if(!isInteger(element)){
			for i := 0; i < 2; i++ {
				go func() {
					var tableauBis []float64
					var tabConstBis []float64
					channelBis := make(chan []float64)

					//Copie de tableau et du tableau de contrainte
					for j := 0; j < len(tabConst); j++ {
						tabConstBis = append(tabConstBis, tabConst[i])
					}
					for j := 0; j < len(tableau); j++ {
						tableauBis = append(tableauBis, tableau[i])
					}

					//Ajout de la nouvelle contrainte dans les copies de tableau
					if i==0 {
						//Traduction tableau et contrainte de Margaux
					} else {
						//Traduction tableau et contrainte de Margaux
					}

					channel <- branch_bound(simplexe(tableauBis, tabConstBis), tableauBis, tabConstBis, channelBis)
				}
			}
		}	
	}
	sol := <- channel
	return sol	
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

