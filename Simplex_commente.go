package main

//bibliothèque de base
import (
	"fmt"
)

func main() {
	var tableau = [][]float64{{1,1}, {2,-1}, {-1,2}}
	var tabConst = []float64{2,0,1}
	fmt.Println(simplex(tableau, tabConst))
}
//donnees: le "Tableau" des coeffs et un tableau contenant les contraintes
//retour : solution s'il y en a une, sinon nil 
func simplex(tableau [][]float64, tabConst []float64) map[string]float64{
	//creation tableau des affectations : taille = nombre de ligne + nombre de colonnes
	alphaTab := createAlphaTab(tableau)
	//tableau qui nous donne la postion des variables dans le tableau alphaTab
	var posVarTableau = createPosVarTableau(tableau)
	//on ne peut pas mettre nil dans un tableau d'entier, on cree donc une
	//valeur qui ne peut exister dans le tableau et que l'on retrouvera 
	var essaie_de_devenir_nil float64
	for i :=0;i<=len(tabConst); i++{
		var aide_pour_devenir_nil bool
		aide_pour_devenir_nil=true
		for j := range tabConst{
			if i==j{
				aide_pour_devenir_nil=false
			}			
		}
		if aide_pour_devenir_nil==true{
			essaie_de_devenir_nil=float64(i)
			break
		}
	}


	//boucle sur le nombre maximum de pivotation que l'on peut avoir
	for k := 0; k < len(tableau); k++ {
		//workingLine est la ligne qui ne respecte pas sa contrainte
		workingLine := checkConst(alphaTab, tabConst, len(tableau[0]), essaie_de_devenir_nil)
		if workingLine == -1 { 
			return alphaTab
		}
		//on cherche la colonne du pivot
		columnPivot := pivot(tableau, tabConst, alphaTab, workingLine, posVarTableau)
		if columnPivot == -1 {
			fmt.Println("Il n'existe pas de solution pour ces contraintes")
			return alphaTab 
		} else {
			//on modifie le tableau des coefficients pour la ligne du pivot
			for i := 0; i < len(tableau[0]); i++ {
				if i == columnPivot {
					tableau[workingLine][i] = 1/tableau[workingLine][i]
				} else {
					tableau[workingLine][i] = -tableau[workingLine][i]/tableau[workingLine][columnPivot]
				}
			}
			//on modifie le tableau des coefficients des autres lignes
			for i := 0; i < len(tableau); i++ {
				if i != workingLine {
					for j := 0; j < len(tableau[0]); j++ {
						if j==columnPivot{
							tableau[i][columnPivot]*=tableau[workingLine][columnPivot]
						} else {
							tableau[i][j] += tableau[workingLine][j] * tableau[i][columnPivot]		
						}
					}
				}
			}
			tabConst[workingLine] = essaie_de_devenir_nil
			//calcul des nouveaux alpha
			fmt.Println(alphaTab)
			fmt.Println(posVarTableau)
			for i := 0; i<len(tableau);i++{
				if i != workingLine {
					var calAlpha float64 
					for j :=0; j<len(tableau[0]);j++{
						fmt.Println(calAlpha, tableau[i][j], alphaTab[posVarTableau[j + len(tableau)]], posVarTableau[j + len(tableau)])
						calAlpha+= tableau[i][j]*alphaTab[posVarTableau[j + len(tableau)]]
					}
					alphaTab[posVarTableau[i]] = calAlpha
				}
			}
		}
	}
	return alphaTab
}



//Cherche la premiere contrainte qui n'est pas respectee et retourne le numero de la ligne associee
func checkConst(alphaTab map[string]float64,  tabConst []float64, nbInconnu int, essaie_de_devenir_nil float64) int{
	for index, element := range tabConst {
		if element != essaie_de_devenir_nil {
			if alphaTab[fmt.Sprint("e", index)] < element {
				return index
			}
		}
	}
	return -1
}

//Renvoie la colonne pivot par rapport a la contrainte a traiter
func pivot(tableau [][]float64,  tabConst []float64, alphaTab map[string]float64, pivotLine int, posVarTableau []string) int{
	for numero_colonne, coefColumn := range tableau[pivotLine]{
		if string(posVarTableau[numero_colonne + len(tableau)][0]) != "e" && coefColumn != 0 {
			var teta float64
			var alphaColumn float64		
			teta = (tabConst[pivotLine] - (alphaTab[posVarTableau[pivotLine]]) ) / coefColumn
			alphaColumn = teta + (alphaTab[posVarTableau[numero_colonne+len(tableau)]])
			var alphaLine float64
			//on calcule alphaLine
			for index2, element2 := range tableau[pivotLine] {
				if coefColumn != element2{
					alphaLine += element2 * (alphaTab[posVarTableau[index2+len(tableau)]])
				}
			}
			alphaLine += coefColumn * alphaColumn
			//on verifie la suitabilite de alphaLine
			if  alphaLine >= tabConst[pivotLine] {
				alphaTab[posVarTableau[pivotLine]] = alphaLine
				alphaTab[posVarTableau[numero_colonne + len(tableau)]] = alphaColumn
				switchVarStringTab(posVarTableau, pivotLine, numero_colonne + len(tableau))
				return numero_colonne
			}
		}
	}
	return -1	 
}

//creation d'un dictionnaire (clé:nomDesVariable et valeur:affectation)
func createAlphaTab(tableau [][]float64) map[string]float64{
	alphaTab := make(map[string]float64)
	for i := 0; i < len(tableau); i++ {
		alphaTab[fmt.Sprint("e", i)] = 0
	}
	for i := 0; i < len(tableau[0]); i++ {
		alphaTab[fmt.Sprint("v", i)] = 0
	}
	return alphaTab
}


func createPosVarTableau(tableau [][]float64) []string{
	var posVarTableau = make([]string, len(tableau[0])+len(tableau))
	lenTab := len(tableau)
	for i := 0; i < lenTab + len(tableau[0]); i++ {
		if i < lenTab {
			posVarTableau[i] = fmt.Sprint("e", i)
		} else {
			posVarTableau[i] = fmt.Sprint("v", i - lenTab)
		}
	}
	return posVarTableau
}

//echange deux places de valeur dans un tableau
func switchVarStringTab(tab []string, pos1 int, pos2 int){
	valeur := tab[pos1]
	tab[pos1] = tab[pos2]
	tab[pos2] = valeur
}