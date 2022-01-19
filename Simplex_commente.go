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
func simplex(tableau [][]float64, tabConst []float64) map[string]float64{
	//creation tableau des affectations : taille = nombre de ligne + nombre de colonnes
	alphaTab := createAlphaTab(tableau)
	var posVarTableau = createPosVarTableau(tableau)
	var bland = make([]string, len(posVarTableau))
	for i:= range posVarTableau{
		bland[i]=posVarTableau[i]
	}
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


	//for a modifier
	for k := 0; k < len(tableau); k++ {
		//workingLine est la ligne qui ne respecte pas sa contrainte
		//on avait &alphaTab, or on ne modifie pas alphaTab dans checkConst
		workingLine := checkConst(alphaTab, tabConst, len(tableau[0]),essaie_de_devenir_nil)
		if workingLine == -1 {
			//on avait un break ici
			// il faut return solution  
			return alphaTab
		}
		//on cherche la colonne du pivot
		//(pour l'instant on utilise pas la regle de bland)
		//on avait pas & pour alphaTab, or on envoie dans pivot un pointeur vers alphaTab
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
						//les nouveaux coeffs sont faux avec une trace, je propose donc une modif
						//la trace que j'ai faite est bonne avec cette proposition
						if j==columnPivot{
							tableau[i][columnPivot]*=tableau[workingLine][columnPivot]
						} else {
							tableau[i][j] += tableau[workingLine][j] * tableau[i][columnPivot]		
						}
					}
				}
			}
			//on considere qu'une variable dont la contrainte est respectee l'a respectera ad vitam
			//a surveiller
			//idee conne pour eviter le probleme du nil 
			tabConst[workingLine] = essaie_de_devenir_nil
			//on modifie les affectations des variables de la base	
			for i := 0; i<len(tableau);i++{
				if i != workingLine {
					var calAlpha float64 
					for j :=0; j<len(tableau[0]);j++{
						calAlpha+= tableau[i][j]*alphaTab[posVarTableau[j + len(tableau)]]
					}
					alphaTab[posVarTableau[i]] = calAlpha
				}
			}
		}
	}
	return alphaTab
}



//Cherche la premiere contrainte qui n'est pas respectee ca place dans tabConst
func checkConst(alphaTab map[string]float64,  tabConst []float64, nbInconnu int, essaie_de_devenir_nil float64) int{
	for index, element := range tabConst {
		if element != essaie_de_devenir_nil {
			//on avait alphaTab[index + nbInconnu] or alpha est
			// construit comme ca initialement : {e1, e2,e3, x,y}
			//si on garde +nbInconnu, ca voudra dire qu'une ligne qui
			// n'existe pas peut etre renvoye
			if alphaTab[fmt.Sprint("e", index)] < element {
				return index
			}
		}
	}
	//il aime pas le return nil, on retourne -1?
	return -1
}

//Renvoie la colonne pivot par rapport a la contrainte a traiter
func pivot(tableau [][]float64,  tabConst []float64, alphaTab map[string]float64, pivotLine int, posVarTableau []string) int{
	for numero_colonne, element1 := range tableau[pivotLine]{
		var teta float64
		var alphaColumn float64
		//il y avait +len(tableau[0]) mais c'est une erreur
		//for bonne_colonne,coeff_bonne_colonne := range(tableau)		
		teta = (tabConst[pivotLine] - (alphaTab[posVarTableau[pivotLine]]) ) / element1
		//renommage alpha inconnu et *alphaTab[index1] devient : *aphaTab[index1 + nombre de lignes]
		//car c'est la variable de la colonne que l'on modifie
		alphaColumn = teta + (alphaTab[posVarTableau[numero_colonne+len(tableau)]])
		//pourquoi imbrication? voir: *$& 
		var alphaLine float64
		//ce for n'est qu'une somme pour determiner l'affectation de la variable ligne du pivot
		for index2, element2 := range tableau[pivotLine] {
			//idem que plus haut, +len(tableau)
			if element1 != element2{
				alphaLine += element2 * (alphaTab[posVarTableau[index2+len(tableau)]])
			
			}
		}
		alphaLine += element1 * alphaColumn
		//*$& a cause d'ici ?
		if  alphaLine >= tabConst[pivotLine] && (teta)>=0 {
			//Changement des place dans alphaTab entre li pivot et la variable d'ecart
			//alphaColumn passe dans la base, donc au debut de alphaTab.
			// L'ajout du nombre de colonnes n'est pas pertinent pour l'index d'alphaTab,
			// car il commence par le nombre de ligne
			alphaTab[posVarTableau[pivotLine]] = alphaLine
			alphaTab[posVarTableau[numero_colonne + len(tableau)]] = alphaColumn
			switchVarStringTab(posVarTableau, pivotLine, numero_colonne + len(tableau))
			return numero_colonne
		}
	}
	//pas de pivot suitable
	return -1	 
}

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

func switchVarStringTab(tab []string, pos1 int, pos2 int){
	valeur := tab[pos1]
	tab[pos1] = tab[pos2]
	tab[pos2] = valeur
}
