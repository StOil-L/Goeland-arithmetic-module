package main


import (
	"fmt"
	"strconv"
    "strings"
    "regexp"
//	"bufio"
//  "os"
	"math/big"
	"math"
//	"time"
    
)

type bAndB struct {
    solBoolStr bool
    solStr  map[string]*big.Rat
}

func main() {
	//fmt.Println("choisissez le test que vous voulez executer : \n 1 pour : x+y>=2,2x-y>=0,-x+2y>=1 \n 2 pour : x+y>=0,x+y>=1,x+y>=2,x+y>=3,x+y>=4 \n 3 pour : x+y>=0,x+2y>=1,x+3y>=2,x+4y>=3,x+5y>=4 \n 4 pour : x>=1/4,x<=1/5 \n 5 pour : x=1/4 \n 6 pour : construire votre matrice des coefficients et vos contraintes \n 7 pour : faire appel au parseur")
	//var x int
	//fmt.Scanln(&x)
	var tabVar = make([]string,0)
	var tableau = [][]*big.Rat{{big.NewRat(1,1)}, {big.NewRat(-1,1)}}
	var tabConst = []*big.Rat{big.NewRat(1,4),big.NewRat(-1,4)}
	channel := make(chan bAndB)
	a,b:=simplex(tableau,tabConst,tabVar)
	fmt.Println(a,b,"\033[0m ")
	fmt.Println(branch_bound(a,b, tableau, tabConst, channel))
/*	
	if x==1 {
		var tableau = [][]*big.Rat{{big.NewRat(1,1),big.NewRat(1,1)}, {big.NewRat(2,1),big.NewRat(-1,1)}, {big.NewRat(-1,1),big.NewRat(2,1)}} //[]*big.Rat{big.NewRat(1,1),new(big.Rat),big.NewRat(1,1)}
		var tabConst = []*big.Rat{big.NewRat(2,1),new(big.Rat),big.NewRat(1,1)}
		fmt.Println("x+y>=2,2x-y>=0,-x+2y>=1")
		fmt.Println(simplex(tableau, tabConst,tabVar))
	}
	if x==2{

		var tableau = [][]*big.Rat{{big.NewRat(1,1),big.NewRat(1,1)}, {big.NewRat(1,1),big.NewRat(1,1)}, {big.NewRat(1,1),big.NewRat(1,1)}, {big.NewRat(1,1),big.NewRat(1,1)}, {big.NewRat(1,1),big.NewRat(1,1)}}
		var tabConst = []*big.Rat{new(big.Rat),big.NewRat(1,1),big.NewRat(2,1),big.NewRat(3,1),big.NewRat(4,1)}
		fmt.Println("x+y>=0,x+y>=1,x+y>=2,x+y>=3,x+y>=4")
		fmt.Println(simplex(tableau,tabConst,tabVar))

	}
	if x==3{
		var tableau = [][]*big.Rat{{big.NewRat(1,1),big.NewRat(1,1)}, {big.NewRat(1,1),big.NewRat(2,1)}, {big.NewRat(1,1),big.NewRat(3,1)}, {big.NewRat(1,1),big.NewRat(4,1)}, {big.NewRat(1,1),big.NewRat(5,1)}}
		var tabConst = []*big.Rat{new(big.Rat),big.NewRat(1,1),big.NewRat(2,1),big.NewRat(3,1),big.NewRat(4,1)}
		fmt.Println("x+y>=0,x+2y>=1,x+3y>=2,x+4y>=3,x+5y>=4")
		fmt.Println(simplex(tableau,tabConst,tabVar))
	}
	if x==4{
		var tableau = [][]*big.Rat{{big.NewRat(1,1)}, {big.NewRat(-1,1)}}
		var tabConst = []*big.Rat{big.NewRat(1,4),big.NewRat(-1,5)}
		fmt.Println("x>=1/4,x<=1/5")
		fmt.Println(simplex(tableau,tabConst,tabVar))
	}

	if x==5{
		var tableau = [][]*big.Rat{{big.NewRat(1,1)}, {big.NewRat(-1,1)}}
		var tabConst = []*big.Rat{big.NewRat(1,4),big.NewRat(-1,4)}
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
		var tableau= make([][]*big.Rat, l)
		for j:=0; j<l;j++{
			tableau[j]=make([]*big.Rat,c)
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
				tableau[k][kk]=new(big.Rat).SetFloat64(a)
				cpt++
				if cpt%c==0{
					cpt=0
					cpt2++
				}
		
			}
			
		}

		fmt.Println("veuillez saisir les contraintes une à une :")
		var tabConst= make([]*big.Rat,l)
		for j,_ := range tabConst{
			var a float64
			fmt.Scanln(&a)
			tabConst[j]=new(big.Rat).SetFloat64(a)
		}

		fmt.Println("\nmatrice des coefficients saisis :",tableau,"\ntableau des contraintes saisies :" ,tabConst)
		fmt.Println(simplex(tableau, tabConst, tabVar))
		
	}
	fmt.Println("\033[0m") 
	
	if x == 7 {
	    // Creation du tableau de coeff et du tableau de contraintes
        var tableau = make([][]*big.Rat,0) 
        var tabConst = make([]*big.Rat,0)
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
			equation = strings.TrimSuffix(equation, "\n")
			tabExe = append(tabExe, equation)
    	}
    	
    	fmt.Println("tabExe =",tabExe,string(tabExe[0][22]))
    //    var tabExe = []string{"20 t - x + y -18 z >= 8","0 t -5 x + y -0 z >= 5","-7 t +3 x +5 y + z >= 33"}
        
        tabConst, tableau, tabVar = addAllConst(tabExe, tableau, tabConst, tabVar)
        fmt.Println("tableau = ",tableau)
        fmt.Println("tabConst = ",tabConst)
        fmt.Println("tabVar = ",tabVar)
        fmt.Println(simplex(tableau, tabConst, tabVar))
	}
*/
}
//donnees: le "Tableau" des coeffs et un tableau contenant les contraintes
//retour : solution s'il y en a une, sinon nil 
func simplex(tableau [][]*big.Rat, tabConst []*big.Rat, tabVar[]string) (map[string]*big.Rat, bool){
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
	
	var IncrementalCoef = make([]*big.Rat, 0)
	var IncrementalAff= make([]*big.Rat,0)

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
func checkConst(alphaTab map[string]*big.Rat,  tabConst []*big.Rat,
	  PosConst []int) int{
	var min int
	min=len(tabConst)
	for _, position := range PosConst  {
		if position!=-1 {
			if min>position && 
			alphaTab[fmt.Sprint("e", position)].Cmp(tabConst[position])==-1{
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
func pivot(tableau [][]*big.Rat,  tabConst []*big.Rat,
	 alphaTab map[string]*big.Rat, pivotLine int, posVarTableau []string,
	  bland []string, PosConst []int) int{
	var variablePivot string
	var index int
	for _,vari := range bland{
		var coefColumn = new(big.Rat)
		for j:=len(tableau);j<len(tableau)+len(tableau[0]);j++{
			if vari==posVarTableau[j]{
				variablePivot=vari
				index=j
				coefColumn=tableau[pivotLine][index-len(tableau)]
			}
		} 	
		for index< len(tableau)+len(tableau[0])  && (coefColumn.Cmp(new(big.Rat))!=0) &&
		 !(coefColumn.Cmp(new(big.Rat))==-1 && alphaTab[variablePivot].Cmp(tabConst[PosConst[pivotLine]])<=0)  {
			var theta = new(big.Rat)/*
			theta = (tabConst[PosConst[pivotLine]] -
				 (alphaTab[posVarTableau[pivotLine]]) ) / coefColumn	
			fmt.Println("Ligne 224: ", tabConst[PosConst[pivotLine]])
			fmt.Println("Ligne 225: ", alphaTab[posVarTableau[pivotLine]])*/
			theta.Mul(new(big.Rat).Add(tabConst[PosConst[pivotLine]], new(big.Rat).Neg(alphaTab[posVarTableau[pivotLine]])), new(big.Rat).Inv(coefColumn))
			var numero_colonne int
			numero_colonne=index-len(tableau)
			var alphaColumn = new(big.Rat)	/*
			alphaColumn = (alphaTab[variablePivot]) + theta*/
			alphaColumn.Add(alphaTab[variablePivot], theta)
			var alphaLine = new(big.Rat)
			//on calcule alphaLine
			for index2, element2 := range tableau[pivotLine] {
				if coefColumn.Cmp(element2) != 0{/*
					alphaLine += element2 * 
					(alphaTab[posVarTableau[index2+len(tableau)]])*/
					alphaLine.Add(alphaLine, new(big.Rat).Mul(element2, alphaTab[posVarTableau[index2+len(tableau)]]))
				}
			}/*
			alphaLine += coefColumn * alphaColumn*/
			alphaLine.Add(alphaLine, new(big.Rat).Mul(coefColumn, alphaColumn))
			alphaTab[posVarTableau[pivotLine]].Set(alphaLine)
			alphaTab[variablePivot].Set(alphaColumn)
			fmt.Println("\033[0m variable \033[36m colonne:",variablePivot+"\033[0m","variable \033[36m ligne:",posVarTableau[pivotLine]+"\033[0m")
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
func createAlphaTab(tableau [][]*big.Rat, tabVar []string) map[string]*big.Rat{
	alphaTab := make(map[string]*big.Rat)
	for i := 0; i < len(tableau); i++ {
	    //fmt.Println(alphaTab[fmt.Sprint("e", i)]) //pour debug
		alphaTab[fmt.Sprint("e", i)] = new(big.Rat)
	}
	if len(tabVar) == 0 {
	    for i := 0; i < len(tableau[0]); i++ {
		alphaTab[fmt.Sprint("v", i)] = new(big.Rat)
	    }
    } else {
        for i := 0; i < len(tableau[0]); i++ {
            alphaTab[tabVar[i]] = new(big.Rat)
	    }
    }
	return alphaTab
}


func createPosVarTableau(tableau [][]*big.Rat, tabVar[]string) []string{
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

func affectation(tableau [][]*big.Rat, workingLine int, 
	alphaTab map[string]*big.Rat, posVarTableau []string, IncrementalAff []*big.Rat){
	for i := 0; i<len(tableau);i++{
		if i != workingLine {
			var calAlpha = new(big.Rat)
			for j :=0; j<len(tableau[0]);j++{/*
				calAlpha+= tableau[i][j]*alphaTab[posVarTableau[j + len(tableau)]]*/
				calAlpha.Add(calAlpha, new(big.Rat).Mul(tableau[i][j], alphaTab[posVarTableau[j + len(tableau)]]))
				IncrementalAff=append(IncrementalAff,alphaTab[posVarTableau[j+len(tableau)]])
			}
			alphaTab[posVarTableau[i]].Set(calAlpha)
		}
	}
}


func coefficients(tableau [][]*big.Rat, columnPivot int, workingLine int, IncrementalCoef []*big.Rat){
	for i := 0; i < len(tableau[0]); i++ {
		if i == columnPivot {/*
			tableau[workingLine][i] = 1/tableau[workingLine][i]*/
			tableau[workingLine][i].Inv(tableau[workingLine][i])
		} else {/*
			tableau[workingLine][i] = 
			-tableau[workingLine][i]/tableau[workingLine][columnPivot]*/
			tableau[workingLine][i].Mul(new(big.Rat).Neg(tableau[workingLine][i]), new(big.Rat).Inv(tableau[workingLine][columnPivot]))
		}
	}
	//on modifie le tableau des coefficients des autres lignes
	for i := 0; i < len(tableau); i++ {
		IncrementalCoef=append(IncrementalCoef,big.NewRat(int64(columnPivot), 1))
		IncrementalCoef=append(IncrementalCoef,tableau[workingLine][columnPivot])
		if i != workingLine {
			for j := 0; j < len(tableau[0]); j++ {
				if j==columnPivot{/*
					tableau[i][columnPivot]*=tableau[workingLine][columnPivot]*/
					tableau[i][columnPivot].Mul(tableau[i][columnPivot], tableau[workingLine][columnPivot])
				} else {/*
					tableau[i][j] += tableau[workingLine][j] *
					 tableau[i][columnPivot]		*/
					tableau[i][j].Add(tableau[i][j], new(big.Rat).Mul(tableau[workingLine][j], tableau[i][columnPivot]))
					 IncrementalCoef=append(IncrementalCoef,tableau[workingLine][j])
		
				}
			}

		}
	}
	
}

// fonction qui prend en parametre un tableau d'equation eqs, une matrice de coeff et un tableau de contraintes 
// et renvoit ces tableaux de coeff et de contraintes remplis
func addAllConst(eqs []string, tableau [][]*big.Rat, tabConst []*big.Rat, tabVar[]string) ([]*big.Rat, [][]*big.Rat, []string){
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
func addOneConst(eq string) (*big.Rat, []*big.Rat,[]string){
    // on split la string avec les espaces, ce qui nous donne un tableaux avec tous les elements de la string
    tabEle := strings.Split(eq, " ")
    // tableau qui va contenir les coefficients de l'equation
    var ligneEq []*big.Rat
    var TabVar []string
    // indique notre position dans le tableau
    posTab := 0
    for i := 0; i < len(tabEle)-2; i++ {
        ligneEq = append(ligneEq, big.NewRat(1,1))
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
            conv,_ := new(big.Rat).SetString(tabEle[i])
           // on l'insere dans notre tableau a la position "posTab"
            ligneEq[posTab]=conv

            
        } else {
            // le cas du -
            if tabEle[i] == "-"{
                ligneEq[posTab].SetFloat64(-1.0)
            // le cas du +
            } else if tabEle[i] != "+" {
                posTab += 1
            }
        }
    }

    // On ajoute la contraintes dans le tableau de contraintes en l'a convertissant d'abord
    lastEle := tabEle[len(tabEle)-1]
	lastEleC,_ := new(big.Rat).SetString(lastEle)
 
    return lastEleC, ligneEq[0:posTab],TabVar
}



func branch_bound(solution map[string]*big.Rat, gotSol bool, tableau [][]*big.Rat, tabConst []*big.Rat, channel chan bAndB) (map[string]*big.Rat, bool){
    var tabVar = make([]string,0)

    //Cas d'arret si solution est fait seulement d'entier
    if (!gotSol) {
        return solution, false
    } else if (estSol(solution)){
        return solution, true
    }

	for index, element := range solution {
		if(!isInteger(element)){
			for i := 0; i < 2; i++ {
				go func() {
					var tableauBis [][]*big.Rat
					var tabConstBis []*big.Rat
					channelBis := make(chan bAndB)

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
						    if string(i) == index {
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
						    if string(i) == index {
						        tabInter = append(tabInter, big.NewRat(-1,1))
						    }else {
						        tabInter = append(tabInter, new(big.Rat))
						    }
						}
						tableauBis = append(tableauBis, tabInter)
					}
					a,b :=simplex(tableauBis,tabConstBis,tabVar)
                    sol, solBool := branch_bound(a,b, tableauBis, tabConstBis, channelBis)
                    stBAndB := bAndB{solBoolStr: solBool, solStr: sol}
                    channel <- stBAndB
                }()
            }
        }  
        break
    }
    stBAndB := <- channel
    if(!stBAndB.solBoolStr){
        stBAndB = <- channel
    }
    return stBAndB.solStr, stBAndB.solBoolStr
}

//Verifie que le nombre donné soit un entier
func isInteger(nombre *big.Rat) bool{
		return nombre.IsInt()
}

//Verifie qu'un tableau contient seulement des entier
func estSol(solution map[string]*big.Rat) bool{
	for _, element := range solution {
		if(!isInteger(element)){
			return false
		}
	}
	return true
}

