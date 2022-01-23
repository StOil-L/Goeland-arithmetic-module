package main

//bibliothèque de base
import (
	"fmt"
	"strconv"
)

func main() {
	var tableau = [][]float64{{1,1}, {2,-1}, {-1,2}}
	var tableau2 = [][]float64{{1,1}, {1,1}, {1,1}, {1,1}, {1,1}}
	var tabConst = []float64{2,0,1}
	var tabConst2 = []float64{0,1,2,3,4}
	fmt.Println(simplex(tableau, tabConst))
	fmt.Println(simplex(tableau2,tabConst2))
}
//donnees: le "Tableau" des coeffs et un tableau contenant les contraintes
//retour : solution s'il y en a une, sinon nil 
func simplex(tableau [][]float64, tabConst []float64) map[string]float64{
	//creation tableau des affectations : taille = nombre de ligne + nombre de colonnes
	alphaTab := createAlphaTab(tableau)
	//tableau qui nous donne la postion des variables dans le tableau alphaTab
	var posVarTableau = createPosVarTableau(tableau)
	var bland = make([]string, len(posVarTableau))
	
	for i:=0;i<len(tableau);i++ {
		bland[i+len(tableau[0])]=posVarTableau[i]
	}
	
	for i:=0;i<len(tableau[0]);i++{
		bland[i]=posVarTableau[i+len(tableau)]
	}
	
	var PosConst = make([]int, len(tabConst))
	for i := 0; i<len(tabConst);i++{
		PosConst[i]=i
	}	
	
	//boucle sur le nombre maximum de pivotation que l'on peut avoir
	k := true
	for k {
		//workingLine est la ligne qui ne respecte pas sa contrainte
		workingLine := checkConst(alphaTab, tabConst, PosConst)
		if workingLine == -1 { 
			return alphaTab
		}
		//on cherche la colonne du pivot
		columnPivot := pivot(tableau, tabConst, alphaTab, workingLine, posVarTableau, bland, PosConst)
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
			
			//calcul des nouveaux alpha
			for i := 0; i<len(tableau);i++{
				if i != workingLine {
					var calAlpha float64 
					for j :=0; j<len(tableau[0]);j++{
						calAlpha+= tableau[i][j]*alphaTab[posVarTableau[j + len(tableau)]]
					}
					alphaTab[posVarTableau[i]] = calAlpha
				}
			}

			fmt.Println(alphaTab)
		}
	}
	return alphaTab
}



//Cherche la premiere contrainte qui n'est pas respectee et retourne le numero de la ligne associee
func checkConst(alphaTab map[string]float64,  tabConst []float64,  PosConst []int) int{
	var min int
	min=len(tabConst)
	for _, position := range PosConst  {
		if position!=-1 {
			if min>position && alphaTab[fmt.Sprint("e", position)] < tabConst[position]{
				min=position
			}	
		}	
	}
	if min != len(tabConst) && min != -1{
		return min
	}

	return -1
}

//Renvoie la colonne pivot par rapport a la contrainte a traiter
func pivot(tableau [][]float64,  tabConst []float64, alphaTab map[string]float64, pivotLine int, posVarTableau []string, bland []string, PosConst []int) int{
	var monBool bool
	monBool = false	
	var variablePivot string
	var index int
	for _,vari := range bland{
		HERE:
			if monBool==true {
			    monBool=false
				continue
			}
		for j:=len(tableau); j< len(tableau)+len(tableau[0]);j++{
			if vari==posVarTableau[j]{
				variablePivot=vari
				index=j
				var theta float64
				var coefColumn float64
				coefColumn=tableau[pivotLine][index-len(tableau)]
				if coefColumn != 0 {
					theta = (tabConst[PosConst[pivotLine]] - (alphaTab[posVarTableau[pivotLine]]) ) / coefColumn
				} else  if coefColumn<0 && alphaTab[variablePivot]<=tabConst[PosConst[pivotLine]] || coefColumn==0{
					monBool=true
					goto HERE
				}
				var numero_colonne int
				numero_colonne=index-len(tableau)
				var alphaColumn float64	
				alphaColumn = (alphaTab[variablePivot]) + theta
				var alphaLine float64
				//on calcule alphaLine
				for index2, element2 := range tableau[pivotLine] {
					if coefColumn != element2{
						alphaLine += element2 * (alphaTab[posVarTableau[index2+len(tableau)]])
					}
				}
				alphaLine += coefColumn * alphaColumn
				alphaTab[posVarTableau[pivotLine]] = alphaLine
				alphaTab[variablePivot] = alphaColumn
				switchVarStringTab(posVarTableau, pivotLine, numero_colonne + len(tableau))
				if variablePivot[0]=='e'{
					switchContrainte(PosConst,variablePivot,posVarTableau[pivotLine])
				} else {
					metavar := string(posVarTableau[pivotLine][1])
					var indice int
					if valeur,err := strconv.Atoi(metavar); err==nil{
						indice=valeur
					}
					PosConst[indice]=-1
				}
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

func switchContrainte(PosConst []int,variableColonne string,variableLigne string){
	var1 := string(variableColonne[1])
	var I1 int
	if valeur1,err1 := strconv.Atoi(var1); err1==nil{
		I1=valeur1
	}
	var2 := string(variableLigne[1])
	var I2 int
	if valeur2,err2 := strconv.Atoi(var2); err2==nil{
		I2=valeur2
	}
	PosConst[I1]=PosConst[I2]
	PosConst[I2]=-1

}
