package main

import (
	"fmt"
)

func main() {
	var tableau = [][]float64{{1,1}, {2,-1}, {-1,2}}
	var constTab = []float64{2,0,1}
	fmt.Println(simplex(tableau, constTab))
}
//donnees: le "Tableau" des coeffs et un tableau contenant les contraintes
//retour : solution s'il y en a une, sinon nil 
func simplex(tableau [][]float64, tabConst []float64) [][]float64{
	//creation tableau des affectations : taille = nombre de ligne + nombre de colonnes
	var alphaTab [len(tableau) + len(tableau[0])]float64
	//initialisation des affectations
	for k:= 0; k< len(alphaTab); k++ {
		alphaTab[k]=0
	}	
	//for a modifier
	for k := 0; k < len(tableau); k++ {
		//actualConstraint est la ligne qui ne respecte pas sa contrainte
		//on avait &alphaTab, or on ne modifie pas alphaTab dans checkConst
		actualConstraint := checkConst(alphaTab, tabConst, len(tableau[0]))
		if actualConstraint == nil {
			//on avait un break ici
			// il faut return solution  
		}
		//on cherche la colonne du pivot
		//(pour l'instant on utilise pas la règle de bland)
		//on avait pas & pour alphaTab, or on envoie dans pivot un pointeur vers alphaTab
		colonnePivot := pivot(tableau, tabConst, &alphaTab, actualConstraint)
		if colonnePivot == nil {
			fmt.Println("Il n'existe pas de solution pour ces contraintes") 
		} else {
			//on modifie le tableau des coefficients pour la ligne du pivot
			for i := 0; i < len(tableau[0]); i++ {
				if i == colonnePivot {
					tableau[actualConstraint][i] = 1/tableau[actualConstraint][i]
				} else {
					tableau[actualConstraint][i] = -tableau[actualConstraint][i]/tableau[actualConstraint][colonnePivot]
				}
			}
			//on modifie le tableau des coefficients des autres lignes
			for i := 0; i < len(tableau); i++ {
				if i != actualConstraint {
					for j := 0; j < len(tableau[0]); j++ {
						//les nouveaux coeffs sont faux avec une trace, je propose donc une modif
						//la trace que j'ai faite est bonne avec cette proposition
						if j==colonnePivot{
							tableau[i][colonnePivot]*=tableau[actualConstraint][colonnePivot]
						} else {
							tableau[i][j] += tableau[actualConstraint][j] * tableau[i][colonnePivot]		
						}
					}
				}
			}
			//on considere qu'une variable dont la contrainte est respectee l'a respectera ad vitam
			//a surveiller 
			tabConst[actualConstraint] = nil
			//on modifie les affectations des variables de la base			
			//je propose ici : !***!!  une autre version
			for i := len(tableau[0]); i < len(alphaTab); i++ {
				//a quoi sert cette conditionnelle?
				if i !=  actualConstraint + len(tableau[0]){
					alphaTab[i] = 0
					//Avec une trace je ne retrouve pas le bon alpha pour j
					for j := 0; j < len(tableau[0]); j++ {
						alphaTab[i] += tableau[i - len(tableau[0])][j] * alphaTab[j]
					}
				}
			}
		}
	}
	return tableau
}

/*  !***!!
for i := 0; i<len(tableau);i++{
	if i != actualConstraint {
		for j :=0; j<len(tableau[0]);j++{
			alphaTab[i]+= tableau[i][j]*alphaTab[len(tableau)]
		}
	}
}
*/



//Cherche la premiere contrainte qui n'est pas respectee ca place dans tabConst
func checkConst(alphaTab []float64,  tabConst []float64, nbInconnu int) int{
	for index, element := range tabConst {
		if element != nil {
			//on avait alphaTab[index + nbInconnu] or alpha est
			// construit comme ca initialement : {e1, e2,e3, x,y}
			//si on garde +nbInconnu, ca voudra dire qu'une ligne qui
			// n'existe pas peut etre renvoye
			if alphaTab[index] < element {
				return index
			}
		}
	}
	//il aime pas le return nil, on retourne -1?
	return -1
}

//Renvoie la colonne pivot par rapport a la contrainte a traiter
func pivot(tableau []float64,  tabConst []float64, alphaTab *[]float64, pivotLine int) int{
	var teta float64
	for index1, element1 := range tableau[pivotLine]{
		//il y avait +len(tableau[0]) mais c'est une erreur
		teta = (tabConst[pivotLine] - (*alphaTab[pivotLine]) ) / element1
		//renommage alpha inconnu et *alphaTab[index1] devient : *aphaTab[index1 + nombre de lignes]
		//car c'est la variable de la colonne que l'on modifie
		alphaColumnPivot := teta + (*alphaTab[index1+len(tableau)])
		//pourquoi imbrication? voir: *$& 
		var varAlphaEcart float64
		//ce for n'est qu'une somme pour determiner l'affectation de la variable ligne du pivot
		for index2, element2 := range tableau[pivotLine] {
			//idem que plus haut, +len(tableau)
			varAlphaEcart += element2 * (*alphaTab[index2+len(tableau)])
		}
		//*$& a cause d'ici ?
		if  varAlphaEcart >= tabConst[pivotLine] {
			//Changement des place dans alphaTab entre li pivot et la variable d'ecart
			//alphaColumnPivot passe dans la base, donc au debut de alphaTab.
			// L'ajout du nombre de colonnes n'est pas pertinent pour l'index d'alphaTab,
			// car il commence par le nombre de ligne
			*alphaTab[index1] = alphaColumnPivot
			*alphaTab[index1 + len(tableau)] = varAlphaEcart
			return index1
		}
	}
	//pas de pivot suitable
	return nil	 
}