package main

import (
	//"fmt"
	"math"
	"math/big"
)

func main() {
	var tableau = [][]*big.Rat{{big.NewRat(1,1),big.NewRat(1,1)}, {big.NewRat(2,1),big.NewRat(-1,1)}, {big.NewRat(-1,1),big.NewRat(2,1)}}
	var tabConst = []*big.Rat{big.NewRat(2,1),new(big.Rat),big.NewRat(1,1)}
	channel := make(chan []*big.Rat)
	branch_bound(simplexe(tableau, tabConst), tableau, tabConst, channel)
}

func branch_bound(solution []*big.Rat, gotSol bool, tableau [][]*big.Rat, tabConst []*big.Rat, channel chan []*big.Rat) []*big.Rat{

	//Cas d'arret si solution est fait seulement d'entier
	if(estSol(solution)){
		return solution
	}
	for index, element := range solution {
		if(!isInteger(element)){
			for i := 0; i < 2; i++ {
				go func() {
					var tableauBis [][]*big.Rat
					var tabConstBis []*big.Rat
					channelBis := make(chan []*big.Rat)

					//Copie de tableau et du tableau de contrainte
					for j := 0; j < len(tabConst); j++ {
						tabConstBis = append(tabConstBis, tabConst[i])
					}
					for j := 0; j < len(tableau); j++ {
						tableauBis = append(tableauBis, tableau[i])
					}

					//Ajout de la nouvelle contrainte dans les copies de tableau
					if i==0 {
					    var tabInter []*big.Rat
					    partiEntiere, _ := element.Float64()
						tabConstBis = append(tabConstBis, new(big.Rat).SetFloat64(math.Ceil(partiEntiere)))
						for i := 0; i < len(solution); i++ {
						    if i == index {
						        tabInter = append(tabInter, big.NewRat(1,1))
						    }else {
						        tabInter = append(tabInter, new(big.Rat))
						    }
						}
						tableauBis = append(tableauBis, tabInter)
					} else {
						var tabInter []*big.Rat
						partiEntiere, _ := element.Float64()
						tabConstBis = append(tabConstBis, new(big.Rat).SetFloat64(-math.Ceil(partiEntiere)))
						for i := 0; i < len(solution); i++ {
						    if i == index {
						        tabInter = append(tabInter, big.NewRat(-1,1))
						    }else {
						        tabInter = append(tabInter, new(big.Rat))
						    }
						}
						tableauBis = append(tableauBis, tabInter)
					}
					channel <- branch_bound(simplexe(tableauBis, tabConstBis), tableauBis, tabConstBis, channelBis)
				}()
			}
		}	
	}
	sol := <- channel
	return sol	
}

//Verifie que le nombre donné soit un entier
func isInteger(nombre *big.Rat) bool{
	nombreFl, exact := nombre.Float64()
	if exact {
		return 	math.Mod(nombreFl,1) == 0 
	} else {
		return 	false
	}
}

//Verifie qu'un tableau contient seulement des entier
func estSol(solution []*big.Rat) bool{
	for _, element := range solution {
		if(!isInteger(element)){
			return false
		}
	}
	return true
}

