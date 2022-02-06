package main


import (
	"fmt"
	"strconv"
    "strings"
    "regexp"
	"bufio"
    "os"
//	"math/big"
//	"time"
    
)


func main() {
	fmt.Println("choisissez le test que vous voulez executer : \n 1 pour : x+y>=2,2x-y>=0,-x+2y>=1 \n 2 pour : x+y>=0,x+y>=1,x+y>=2,x+y>=3,x+y>=4 \n 3 pour : x+y>=0,x+2y>=1,x+3y>=2,x+4y>=3,x+5y>=4 \n 4 pour : x>=1/4,x<=1/5 \n 5 pour : x=1/4 \n 6 pour : construire votre matrice des coefficients et vos contraintes \n 7 pour : faire appel au parseur")
	var x int
	fmt.Scanln(&x)
	var tabVar = make([]string,0)
	if x==1 {
		var tableau = [][]float64{{1,1}, {2,-1}, {-1,2}}
		var tabConst = []float64{2,0,1}
		fmt.Println("x+y>=2,2x-y>=0,-x+2y>=1")
		fmt.Println(simplex(tableau, tabConst,tabVar))
	}
	if x==2{

		var tableau = [][]float64{{1,1}, {1,1}, {1,1}, {1,1}, {1,1}}
		var tabConst = []float64{0,1,2,3,4}
		fmt.Println("x+y>=0,x+y>=1,x+y>=2,x+y>=3,x+y>=4")
		fmt.Println(simplex(tableau,tabConst,tabVar))

	}
	if x==3{
		var tableau = [][]float64{{1,1}, {1,2}, {1,3}, {1,4}, {1,5}}
		var tabConst = []float64{0,1,2,3,4}
		fmt.Println("x+y>=0,x+2y>=1,x+3y>=2,x+4y>=3,x+5y>=4")
		fmt.Println(simplex(tableau,tabConst,tabVar))
	}
	if x==4{
		var tableau = [][]float64{{1}, {-1}}
		var tabConst = []float64{0.25,-0.2}
		fmt.Println("x>=1/4,x<=1/5")
		fmt.Println(simplex(tableau,tabConst,tabVar))
	}

	if x==5{
		var tableau = [][]float64{{1}, {-1}}
		var tabConst = []float64{0.25,-0.25}
		fmt.Println("x=1/4")
		fmt.Println(simplex(tableau,tabConst,tabVar))
	}

	if x==6{
		fmt.Println("veuillez saisir le nombre de lignes de la matrice des coefficients")
		var c int
		var l int
		fmt.Scanln(&l)
		fmt.Println("veuillez saisir le nombre de colonnes de la matrice des coefficients")
		fmt.Scanln(&c)
		var tableau= make([][]float64, l)
		for j:=0; j<l;j++{
			tableau[j]=make([]float64,c)
		}
		cpt:=0
		cpt2:=0
		for k,_ := range tableau{
			for kk,_ := range tableau[0]{
				var a float64
				if cpt==0{
					fmt.Println("veuillez saisir la ligne :",cpt2+1)
				}
				fmt.Scanln(&a)
				tableau[k][kk]=a
				cpt++
				if cpt%c==0{
					cpt=0
					cpt2++
				}
		
			}
			
		}

		fmt.Println("veuillez saisir les contraintes une à une :")
		var tabConst= make([]float64,l)
		for j,_ := range tabConst{
			var a float64
			fmt.Scanln(&a)
			tabConst[j]=a
		}

		fmt.Println(tableau, tabConst)
		fmt.Println(simplex(tableau, tabConst, tabVar))
		
	}
	fmt.Println("\033[0m") 
	
	if x == 7 {
	    // Creation du tableau de coeff et du tableau de contraintes
        var tableau = make([][]float64,0) 
        var tabConst = make([]float64,0)
        var tabVar = make([]string,0)
        fmt.Println("Veuillez saisir le nombre d'equations")
        var nbrEq int
        fmt.Scanln(&nbrEq)
        fmt.Println("nbrEq =",nbrEq)
    	var tabExe = make([]string,0)
    	//var equation string
    	for i:=0; i<nbrEq; i++{
    	    reader := bufio.NewReader(os.Stdin)
            fmt.Print("Entrez une equation: ")
            equation, _ := reader.ReadString('\n')
    		tabExe = append(tabExe, equation)
    	}
    	
    	fmt.Println("tabExe =",tabExe)
        //var tabExe = []string{"20 t - x + y -18 z >= 8","0 t -5 x + y -0 z >= 5","-7 t +3 x +5 y + z >= 33"}
        
        tabConst, tableau, tabVar = addAllConst(tabExe, tableau, tabConst, tabVar)
        fmt.Println("tableau = ",tableau)
        fmt.Println("tabConst = ",tabConst)
        fmt.Println("tabVar = ",tabVar)
        fmt.Println(simplex(tableau, tabConst, tabVar))
	}

}
//donnees: le "Tableau" des coeffs et un tableau contenant les contraintes
//retour : solution s'il y en a une, sinon nil 
func simplex(tableau [][]float64, tabConst []float64, tabVar[]string) (map[string]float64, bool){
	//creation tableau des affectations : taille = nombre de ligne + nombre de colonnes
	alphaTab := createAlphaTab(tableau, tabVar)
	//tableau qui nous donne la postion des variables dans le tableau alphaTab
	var posVarTableau = createPosVarTableau(tableau, tabVar)
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
	
	var IncrementalCoef = make([]float64, 0)
	var IncrementalAff= make([]float64,0)

	//boucle sur le nombre maximum de pivotation que l'on peut avoir
	for true {
		//workingLine est la ligne qui ne respecte pas sa contrainte
		workingLine := checkConst(alphaTab, tabConst, PosConst)
		if workingLine == -1 {
			fmt.Println(" \033[33m La solution est : ") 
			return  alphaTab,true
		}
		//on cherche la colonne du pivot
		columnPivot := pivot(tableau, tabConst, alphaTab, workingLine,
			 posVarTableau, bland, PosConst)
		if columnPivot == -1 {
			fmt.Println(" \033[33m") 
			fmt.Println("Il n'existe pas de solution pour ces contraintes")
			return alphaTab,false 
		} else {
			//on modifie le tableau des coefficients pour la ligne du pivot
			coefficients(tableau,columnPivot,workingLine,IncrementalCoef)
			//calcul des nouveaux alpha
			affectation(tableau,workingLine,alphaTab,posVarTableau,IncrementalAff)
			//time.Sleep(time.Second)
			fmt.Println("\033[34m affectations :" ,alphaTab,"\033[0m")
			fmt.Println("\033[35m matrice des coefficients :",tableau,"\033[0m")
		}
	}
	return alphaTab,false
}



//Cherche la premiere contrainte qui n'est pas respectee et retourne le numero de la ligne associee
func checkConst(alphaTab map[string]float64,  tabConst []float64,
	  PosConst []int) int{
	var min int
	min=len(tabConst)
	for _, position := range PosConst  {
		if position!=-1 {
			if min>position && 
			alphaTab[fmt.Sprint("e", position)] < tabConst[position]{
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
func pivot(tableau [][]float64,  tabConst []float64,
	 alphaTab map[string]float64, pivotLine int, posVarTableau []string,
	  bland []string, PosConst []int) int{
	var variablePivot string
	var index int
	for _,vari := range bland{
		var coefColumn float64
		for j:=len(tableau);j<len(tableau)+len(tableau[0]);j++{
			if vari==posVarTableau[j]{
				variablePivot=vari
				index=j
				coefColumn=tableau[pivotLine][index-len(tableau)]
			}
		} 	
		for index< len(tableau)+len(tableau[0])  && (coefColumn != 0) &&
		 !(coefColumn<0 && alphaTab[variablePivot]<=tabConst[PosConst[pivotLine]])  {
			var theta float64
			theta = (tabConst[PosConst[pivotLine]] -
				 (alphaTab[posVarTableau[pivotLine]]) ) / coefColumn				
			var numero_colonne int
			numero_colonne=index-len(tableau)
			var alphaColumn float64	
			alphaColumn = (alphaTab[variablePivot]) + theta
			var alphaLine float64
			//on calcule alphaLine
			for index2, element2 := range tableau[pivotLine] {
				if coefColumn != element2{
					alphaLine += element2 * 
					(alphaTab[posVarTableau[index2+len(tableau)]])
				}
			}
			alphaLine += coefColumn * alphaColumn
			alphaTab[posVarTableau[pivotLine]] = alphaLine
			alphaTab[variablePivot] = alphaColumn
			fmt.Println("variable \033[36m colonne:",variablePivot+"\033[0m","variable \033[36m ligne:",posVarTableau[pivotLine]+"\033[0m")
			if variablePivot[0]=='e' {
				switchContrainte(PosConst,variablePivot,posVarTableau[pivotLine])
			} else {
				metavar := string(posVarTableau[pivotLine][1])
				var indice int
				if valeur,err := strconv.Atoi(metavar); err==nil{
					indice=valeur
				}
				PosConst[indice]=-1
			}

			switchVarStringTab(posVarTableau, pivotLine,
				 numero_colonne + len(tableau))
			fmt.Println("\033[36m theta\033[0m =\033[36m",theta,"\033[0m")
			return numero_colonne
		}
	}
	return -1	 
}


//creation d'un dictionnaire (clé:nomDesVariable et valeur:affectation)
func createAlphaTab(tableau [][]float64, tabVar []string) map[string]float64{
	alphaTab := make(map[string]float64)
	for i := 0; i < len(tableau); i++ {
		alphaTab[fmt.Sprint("e", i)] = 0
	}
	if len(tabVar) == 0 {
	    for i := 0; i < len(tableau[0]); i++ {
		alphaTab[fmt.Sprint("v", i)] = 0
	    }
    } else {
        for i := 0; i < len(tableau[0]); i++ {
            alphaTab[tabVar[i]] = 0
	    }
    }
	return alphaTab
}


func createPosVarTableau(tableau [][]float64, tabVar[]string) []string{
	var posVarTableau = make([]string, len(tableau[0])+len(tableau))
	lenTab := len(tableau)
	for i := 0; i < lenTab + len(tableau[0]); i++ {
		if i < lenTab {
			posVarTableau[i] = fmt.Sprint("e", i)
		} else {
		    if len(tabVar) == 0 {
		        posVarTableau[i] = fmt.Sprint("v", i - lenTab)
		    } else {
		        posVarTableau[i] = tabVar[i-lenTab]
		    }
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

	PosConst[I1]=I1
	PosConst[I2]=-1

}

func affectation(tableau [][]float64, workingLine int, 
	alphaTab map[string]float64, posVarTableau []string, IncrementalAff []float64){
	for i := 0; i<len(tableau);i++{
		if i != workingLine {
			var calAlpha float64 
			for j :=0; j<len(tableau[0]);j++{
				calAlpha+= tableau[i][j]*alphaTab[posVarTableau[j + len(tableau)]]
				IncrementalAff=append(IncrementalAff,alphaTab[posVarTableau[j+len(tableau)]])
			}
			alphaTab[posVarTableau[i]] = calAlpha
		}
	}
}


func coefficients(tableau [][]float64, columnPivot int, workingLine int, IncrementalCoef []float64){
	for i := 0; i < len(tableau[0]); i++ {
		if i == columnPivot {
			tableau[workingLine][i] = 1/tableau[workingLine][i]
		} else {
			tableau[workingLine][i] = 
			-tableau[workingLine][i]/tableau[workingLine][columnPivot]
		}
	}
	//on modifie le tableau des coefficients des autres lignes
	for i := 0; i < len(tableau); i++ {
		IncrementalCoef=append(IncrementalCoef,float64(columnPivot))
		IncrementalCoef=append(IncrementalCoef,tableau[workingLine][columnPivot])
		if i != workingLine {
			for j := 0; j < len(tableau[0]); j++ {
				if j==columnPivot{
					tableau[i][columnPivot]*=tableau[workingLine][columnPivot]
				} else {
					tableau[i][j] += tableau[workingLine][j] *
					 tableau[i][columnPivot]		
					 IncrementalCoef=append(IncrementalCoef,tableau[workingLine][j])
		
				}
			}

		}
	}
	
}

// fonction qui prend en parametre un tableau d'equation eqs, une matrice de coeff et un tableau de contraintes 
// et renvoit ces tableaux de coeff et de contraintes remplis
func addAllConst(eqs []string, tableau [][]float64, tabConst []float64, tabVar[]string) ([]float64, [][]float64, []string){
    for _, element := range eqs {
        lastEle, tab, Var := addOneConst(element)
        tableau = append(tableau, tab)
        tabConst = append(tabConst, lastEle)

        for i := 0;i < len(Var);i++{
          var present bool
          present = false
          for j := 0;j < len(tabVar);j++{
            if Var[i] == tabVar[j] {
              present = true
            }
          }
          if present == false {
              tabVar = append(tabVar,Var[i])
            }
        }
        //fmt.Println("tabConst =",tabConst)
    }
    return tabConst, tableau, tabVar
}

// fonction qui prend en parametre une equation et qui implemente les tableaux de coeff et de contraintes 
func addOneConst(eq string) (float64, []float64,[]string){
    // on split la string avec les espaces, ce qui nous donne un tableaux avec tous les elements de la string
    tabEle := strings.Split(eq, " ")
    // tableau qui va contenir les coefficients de l'equation
    var ligneEq []float64
    var TabVar []string
    // indique notre position dans le tableau
    posTab := 0
    for i := 0; i < len(tabEle)-2; i++ {
        ligneEq = append(ligneEq, 1.0)
        // nous permet de savoir si notre caractere est un chiffre
        re := regexp.MustCompile(`[0-9]`)
        isFig := re.FindString(tabEle[i])
        // nous permet de savoir si notre caractere est une lettre
        re2 := regexp.MustCompile(`[a-z]`)
        isLet := re2.FindString(tabEle[i])

        if isLet != "" {
          TabVar = append(TabVar,isLet)
        }       


        // si on a un chiffre
        if isFig != "" {
            // on converti notre caractere qui est une string en float64
            conv,_ := strconv.ParseFloat(tabEle[i], 64)
            // on l'insere dans notre tableau a la position "posTab"
            ligneEq[posTab] = conv

            
        } else {
            // le cas du -
            if tabEle[i] == "-"{
                ligneEq[posTab] = -1.0
            // le cas du +
            } else if tabEle[i] != "+" {
                posTab += 1
            }
        }
    }

    // On ajoute la contraintes dans le tableau de contraintes en l'a convertissant d'abord
    lastEle := tabEle[len(tabEle)-1]
    lastEleC,_ := strconv.ParseFloat(lastEle, 64)
    fmt.Println("lastEleC =",lastEleC)
    
    return lastEleC, ligneEq[0:posTab],TabVar
}

