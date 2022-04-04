package main

import (
	"fmt"
    "strings"
	"bufio"
    "os"
	"math/big"
	"sync"
	"arith"
)

type bAndB struct {
    solBoolStr bool
    solStr  map[string]*big.Rat
}

var lock sync.Mutex

func main() {
	fmt.Println("choisissez le test que vous voulez executer, ajoutez 1 pour le B&B \n\n 1 pour : x+y>=2,2x-y>=0,-x+2y>=1 \n 3 pour : x+y>=0,x+y>=1,x+y>=2,x+y>=3,x+y>=4 \n 5 pour : x+y>=0,x+2y>=1,x+3y>=2,x+4y>=3,x+5y>=4 \n 7 pour : x>=1/4,x<=1/5 \n 9 pour : x=1/4 \n 11 pour : construire votre matrice des coefficients et vos contraintes \n 13 pour : faire appel au parseur \n 15 pour : écrire à la main le système en utilisant le parseur\n\nexemple : l'option 1 lance le simplexe du jeu de test 1, l'option 2 lance le B&B sur le jeu de test 1")
	var x int
	fmt.Scanln(&x)
	var tabVar = make([]string,0)	
	var IncrementalCoef = make([]*big.Rat, 0)
	var IncrementalAff= make([]*big.Rat,0)
	lock = sync.Mutex{}
	

	if x==1 {
		var tableau = [][]*big.Rat{{big.NewRat(1,1),big.NewRat(1,1)}, {big.NewRat(2,1),big.NewRat(-1,1)}, {big.NewRat(-1,1),big.NewRat(2,1)}} //[]*big.Rat{big.NewRat(1,1),new(big.Rat),big.NewRat(1,1)}
		var tabConst = []*big.Rat{big.NewRat(2,1),new(big.Rat),big.NewRat(1,1)}
		fmt.Println("x+y>=2,2x-y>=0,-x+2y>=1")

	//creation tableau des affectations : taille = nombre de ligne + nombre de colonnes
	alphaTab := createAlphaTab(tableau, tabVar)
	//tableau qui nous donne la postion des variables dans le tableau alphaTab
	var posVarTableau = CreatePosVarTableau(tableau, tabVar)
	var bland = make([]string, len(posVarTableau))
	fmt.Println("\033[0m") 

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

		Simplexe(tableau, tabConst,tabVar,IncrementalCoef,IncrementalAff,posVarTableau,bland,PosConst,alphaTab)
	}
	
	if x==2 {
		var tableau = [][]*big.Rat{{big.NewRat(1,1),big.NewRat(1,1)}, {big.NewRat(2,1),big.NewRat(-1,1)}, {big.NewRat(-1,1),big.NewRat(2,1)}} //[]*big.Rat{big.NewRat(1,1),new(big.Rat),big.NewRat(1,1)}
		var tabConst = []*big.Rat{big.NewRat(2,1),new(big.Rat),big.NewRat(1,1)}
		fmt.Println("x+y>=2,2x-y>=0,-x+2y>=1")
		channel := make(chan bAndB)

	//creation tableau des affectations : taille = nombre de ligne + nombre de colonnes
	alphaTab := createAlphaTab(tableau, tabVar)
	//tableau qui nous donne la postion des variables dans le tableau alphaTab
	var posVarTableau = CreatePosVarTableau(tableau, tabVar)
	var bland = make([]string, len(posVarTableau))
	fmt.Println("\033[0m") 

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
	
		a,b,c,Incremental_Coef,Incremental_Aff,posV,rBland,posC:=Simplexe(tableau,tabConst,tabVar,IncrementalCoef,IncrementalAff,posVarTableau,bland,PosConst,alphaTab)
		fmt.Println(Branch_bound(a,b,c, tableau, tabConst, channel,Incremental_Coef,Incremental_Aff,posV,rBland,posC))
	}
	if x==3{
		var tableau = [][]*big.Rat{{big.NewRat(1,1),big.NewRat(1,1)}, {big.NewRat(1,1),big.NewRat(1,1)}, {big.NewRat(1,1),big.NewRat(1,1)}, {big.NewRat(1,1),big.NewRat(1,1)}, {big.NewRat(1,1),big.NewRat(1,1)}}
		var tabConst = []*big.Rat{new(big.Rat),big.NewRat(1,1),big.NewRat(2,1),big.NewRat(3,1),big.NewRat(4,1)}
		fmt.Println("x+y>=0,x+y>=1,x+y>=2,x+y>=3,x+y>=4")

	//creation tableau des affectations : taille = nombre de ligne + nombre de colonnes
	alphaTab := createAlphaTab(tableau, tabVar)
	//tableau qui nous donne la postion des variables dans le tableau alphaTab
	var posVarTableau = CreatePosVarTableau(tableau, tabVar)
	var bland = make([]string, len(posVarTableau))
	fmt.Println("\033[0m") 

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
		
		Simplexe(tableau, tabConst,tabVar,IncrementalCoef,IncrementalAff,posVarTableau,bland,PosConst,alphaTab)
	}

	if x==4{
		var tableau = [][]*big.Rat{{big.NewRat(1,1),big.NewRat(1,1)}, {big.NewRat(1,1),big.NewRat(1,1)}, {big.NewRat(1,1),big.NewRat(1,1)}, {big.NewRat(1,1),big.NewRat(1,1)}, {big.NewRat(1,1),big.NewRat(1,1)}}
		var tabConst = []*big.Rat{new(big.Rat),big.NewRat(1,1),big.NewRat(2,1),big.NewRat(3,1),big.NewRat(4,1)}
		fmt.Println("x+y>=0,x+y>=1,x+y>=2,x+y>=3,x+y>=4")
		channel := make(chan bAndB)

	//creation tableau des affectations : taille = nombre de ligne + nombre de colonnes
	alphaTab := createAlphaTab(tableau, tabVar)
	//tableau qui nous donne la postion des variables dans le tableau alphaTab
	var posVarTableau = CreatePosVarTableau(tableau, tabVar)
	var bland = make([]string, len(posVarTableau))
	fmt.Println("\033[0m") 

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
		
		a,b,c,Incremental_Coef,Incremental_Aff,posV,rBland,posC:=Simplexe(tableau,tabConst,tabVar,IncrementalCoef,IncrementalAff,posVarTableau,bland,PosConst,alphaTab)
		fmt.Println(Branch_bound(a,b,c, tableau, tabConst, channel,Incremental_Coef,Incremental_Aff,posV,rBland,posC))
	}
	if x==5{
		var tableau = [][]*big.Rat{{big.NewRat(1,1),big.NewRat(1,1)}, {big.NewRat(1,1),big.NewRat(2,1)}, {big.NewRat(1,1),big.NewRat(3,1)}, {big.NewRat(1,1),big.NewRat(4,1)}, {big.NewRat(1,1),big.NewRat(5,1)}}
		var tabConst = []*big.Rat{new(big.Rat),big.NewRat(1,1),big.NewRat(2,1),big.NewRat(3,1),big.NewRat(4,1)}
		fmt.Println("x+y>=0,x+2y>=1,x+3y>=2,x+4y>=3,x+5y>=4")

	//creation tableau des affectations : taille = nombre de ligne + nombre de colonnes
	alphaTab := createAlphaTab(tableau, tabVar)
	//tableau qui nous donne la postion des variables dans le tableau alphaTab
	var posVarTableau = CreatePosVarTableau(tableau, tabVar)
	var bland = make([]string, len(posVarTableau))
	fmt.Println("\033[0m") 

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
		
		Simplexe(tableau, tabConst,tabVar,IncrementalCoef,IncrementalAff,posVarTableau,bland,PosConst,alphaTab)

	}

	if x==6{
		var tableau = [][]*big.Rat{{big.NewRat(1,1),big.NewRat(1,1)}, {big.NewRat(1,1),big.NewRat(2,1)}, {big.NewRat(1,1),big.NewRat(3,1)}, {big.NewRat(1,1),big.NewRat(4,1)}, {big.NewRat(1,1),big.NewRat(5,1)}}
		var tabConst = []*big.Rat{new(big.Rat),big.NewRat(1,1),big.NewRat(2,1),big.NewRat(3,1),big.NewRat(4,1)}
		fmt.Println("x+y>=0,x+2y>=1,x+3y>=2,x+4y>=3,x+5y>=4")
		channel := make(chan bAndB)

	//creation tableau des affectations : taille = nombre de ligne + nombre de colonnes
	alphaTab := createAlphaTab(tableau, tabVar)
	//tableau qui nous donne la postion des variables dans le tableau alphaTab
	var posVarTableau = CreatePosVarTableau(tableau, tabVar)
	var bland = make([]string, len(posVarTableau))
	fmt.Println("\033[0m") 

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
		
		a,b,c,Incremental_Coef,Incremental_Aff,posV,rBland,posC:=Simplexe(tableau,tabConst,tabVar,IncrementalCoef,IncrementalAff,posVarTableau,bland,PosConst,alphaTab)
		fmt.Println(Branch_bound(a,b,c, tableau, tabConst, channel,Incremental_Coef,Incremental_Aff,posV,rBland,posC))
	}
	if x==7{
		var tableau = [][]*big.Rat{{big.NewRat(1,1)}, {big.NewRat(-1,1)}}
		var tabConst = []*big.Rat{big.NewRat(1,4),big.NewRat(-1,5)}
		fmt.Println("x>=1/4,x<=1/5")

	//creation tableau des affectations : taille = nombre de ligne + nombre de colonnes
	alphaTab := createAlphaTab(tableau, tabVar)
	//tableau qui nous donne la postion des variables dans le tableau alphaTab
	var posVarTableau = CreatePosVarTableau(tableau, tabVar)
	var bland = make([]string, len(posVarTableau))
	fmt.Println("\033[0m") 

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
		
		Simplexe(tableau, tabConst,tabVar,IncrementalCoef,IncrementalAff,posVarTableau,bland,PosConst,alphaTab)

	}

	if x==8{
		var tableau = [][]*big.Rat{{big.NewRat(1,1)}, {big.NewRat(-1,1)}}
		var tabConst = []*big.Rat{big.NewRat(1,4),big.NewRat(-1,5)}
		fmt.Println("x>=1/4,x<=1/5")
		channel := make(chan bAndB)

	//creation tableau des affectations : taille = nombre de ligne + nombre de colonnes
	alphaTab := createAlphaTab(tableau, tabVar)
	//tableau qui nous donne la postion des variables dans le tableau alphaTab
	var posVarTableau = CreatePosVarTableau(tableau, tabVar)
	var bland = make([]string, len(posVarTableau))
	fmt.Println("\033[0m") 

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
		
		a,b,c,Incremental_Coef,Incremental_Aff,posV,rBland,posC:=Simplexe(tableau,tabConst,tabVar,IncrementalCoef,IncrementalAff,posVarTableau,bland,PosConst,alphaTab)
		fmt.Println(Branch_bound(a,b,c, tableau, tabConst, channel,Incremental_Coef,Incremental_Aff,posV,rBland,posC))
	}
	if x==9{
		var tableau = [][]*big.Rat{{big.NewRat(1,1)}, {big.NewRat(-1,1)}}
		var tabConst = []*big.Rat{big.NewRat(1,4),big.NewRat(-1,4)}
		fmt.Println("x=1/4")

	//creation tableau des affectations : taille = nombre de ligne + nombre de colonnes
	alphaTab := createAlphaTab(tableau, tabVar)
	//tableau qui nous donne la postion des variables dans le tableau alphaTab
	var posVarTableau = CreatePosVarTableau(tableau, tabVar)
	var bland = make([]string, len(posVarTableau))
	fmt.Println("\033[0m") 

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
		
		Simplexe(tableau, tabConst,tabVar,IncrementalCoef,IncrementalAff,posVarTableau,bland,PosConst,alphaTab)

	}

	if x==10{
		var tableau = [][]*big.Rat{{big.NewRat(1,1)}, {big.NewRat(-1,1)}}
		var tabConst = []*big.Rat{big.NewRat(1,4),big.NewRat(-1,4)}
		fmt.Println("x=1/4")
		channel := make(chan bAndB)

	//creation tableau des affectations : taille = nombre de ligne + nombre de colonnes
	alphaTab := createAlphaTab(tableau, tabVar)
	//tableau qui nous donne la postion des variables dans le tableau alphaTab
	var posVarTableau = CreatePosVarTableau(tableau, tabVar)
	var bland = make([]string, len(posVarTableau))
	fmt.Println("\033[0m") 

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
		
		a,b,c,Incremental_Coef,Incremental_Aff,posV,rBland,posC:=Simplexe(tableau,tabConst,tabVar,IncrementalCoef,IncrementalAff,posVarTableau,bland,PosConst,alphaTab)
		fmt.Println(Branch_bound(a,b,c, tableau, tabConst, channel,Incremental_Coef,Incremental_Aff,posV,rBland,posC))
	}

	if x==11{
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

	//creation tableau des affectations : taille = nombre de ligne + nombre de colonnes
	alphaTab := createAlphaTab(tableau, tabVar)
	//tableau qui nous donne la postion des variables dans le tableau alphaTab
	var posVarTableau = CreatePosVarTableau(tableau, tabVar)
	var bland = make([]string, len(posVarTableau))
	fmt.Println("\033[0m") 

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
		
		Simplexe(tableau, tabConst,tabVar,IncrementalCoef,IncrementalAff,posVarTableau,bland,PosConst,alphaTab)

	}
	fmt.Println("\033[0m") 

	if x==12{
		fmt.Println("veuillez saisir le nombre de lignes de la matrice des coefficients")
		var colonne int
		var ligne int
		fmt.Scanln(&ligne)
		fmt.Println("veuillez saisir le nombre de colonnes de la matrice des coefficients")
		fmt.Scanln(&colonne)
		var tableau= make([][]*big.Rat, ligne)
		for j:=0; j<ligne;j++{
			tableau[j]=make([]*big.Rat,colonne)
		}
		cpt:=0
		cpt2:=0
		for k,_ := range tableau{
			for kk,_ := range tableau[0]{
				var aa float64
				if cpt==0{
					fmt.Println("veuillez saisir la ligne :",cpt2+1)
				}
				fmt.Scanln(&aa)
				tableau[k][kk]=new(big.Rat).SetFloat64(aa)
				cpt++
				if cpt%colonne==0{
					cpt=0
					cpt2++
				}
		
			}
			
		}

		fmt.Println("veuillez saisir les contraintes une à une :")
		var tabConst= make([]*big.Rat,ligne)
		for j,_ := range tabConst{
			var aa float64
			fmt.Scanln(&aa)
			tabConst[j]=new(big.Rat).SetFloat64(aa)
		}

		fmt.Println("\nmatrice des coefficients saisis :",tableau,"\ntableau des contraintes saisies :" ,tabConst)
		channel := make(chan bAndB)

	//creation tableau des affectations : taille = nombre de ligne + nombre de colonnes
	alphaTab := createAlphaTab(tableau, tabVar)
	//tableau qui nous donne la postion des variables dans le tableau alphaTab
	var posVarTableau = CreatePosVarTableau(tableau, tabVar)
	var bland = make([]string, len(posVarTableau))
	fmt.Println("\033[0m") 

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
		
		a,b,c,Incremental_Coef,Incremental_Aff,posV,rBland,posC:=Simplexe(tableau,tabConst,tabVar,IncrementalCoef,IncrementalAff,posVarTableau,bland,PosConst,alphaTab)
		fmt.Println(Branch_bound(a,b,c, tableau, tabConst, channel,Incremental_Coef,Incremental_Aff,posV,rBland,posC))
		
	}
	fmt.Println("\033[0m")
	if x == 13 {
	    // Creation du tableau de coeff et du tableau de contraintes
        var tableau = make([][]*big.Rat,0) 
        var tabConst = make([]*big.Rat,0)
        var tabVar = make([]string,0)
        var tabExe = []string{"20 t - x + y -18 z >= 8","0 t -5 x + y -0 z >= 5","-7 t +3 x +5 y + z >= 33"}
        
        tabConst, tableau, tabVar = addAllConst(tabExe, tableau, tabConst, tabVar)
        fmt.Println("tableau = ",tableau)
        fmt.Println("tabConst = ",tabConst)
        fmt.Println("tabVar = ",tabVar)

	//creation tableau des affectations : taille = nombre de ligne + nombre de colonnes
	alphaTab := createAlphaTab(tableau, tabVar)
	//tableau qui nous donne la postion des variables dans le tableau alphaTab
	var posVarTableau = CreatePosVarTableau(tableau, tabVar)
	var bland = make([]string, len(posVarTableau))
	fmt.Println("\033[0m") 

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
		
		Simplexe(tableau, tabConst,tabVar,IncrementalCoef,IncrementalAff,posVarTableau,bland,PosConst,alphaTab)
	}
	if x == 14 {
	    // Creation du tableau de coeff et du tableau de contraintes
        var tableau = make([][]*big.Rat,0) 
        var tabConst = make([]*big.Rat,0)
        var tabVar = make([]string,0)
        var tabExe = []string{"20 t - x + y -18 z >= 8","0 t -5 x + y -0 z >= 5","-7 t +3 x +5 y + z >= 33"}
        
        tabConst, tableau, tabVar = addAllConst(tabExe, tableau, tabConst, tabVar)
        fmt.Println("tableau = ",tableau)
        fmt.Println("tabConst = ",tabConst)
        fmt.Println("tabVar = ",tabVar)
		channel := make(chan bAndB)

	//creation tableau des affectations : taille = nombre de ligne + nombre de colonnes
	alphaTab := createAlphaTab(tableau, tabVar)
	//tableau qui nous donne la postion des variables dans le tableau alphaTab
	var posVarTableau = CreatePosVarTableau(tableau, tabVar)
	var bland = make([]string, len(posVarTableau))
	fmt.Println("\033[0m") 

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
		
		a,b,c,Incremental_Coef,Incremental_Aff,posV,rBland,posC:=Simplexe(tableau,tabConst,tabVar,IncrementalCoef,IncrementalAff,posVarTableau,bland,PosConst,alphaTab)
		fmt.Println(Branch_bound(a,b,c, tableau, tabConst, channel,Incremental_Coef,Incremental_Aff,posV,rBland,posC))
		
	}


	if x == 15 {
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
    	
        tabConst, tableau, tabVar = addAllConst(tabExe, tableau, tabConst, tabVar)
        fmt.Println("tableau = ",tableau)
        fmt.Println("tabConst = ",tabConst)
        fmt.Println("tabVar = ",tabVar)

	//creation tableau des affectations : taille = nombre de ligne + nombre de colonnes
	alphaTab := createAlphaTab(tableau, tabVar)
	//tableau qui nous donne la postion des variables dans le tableau alphaTab
	var posVarTableau = CreatePosVarTableau(tableau, tabVar)
	var bland = make([]string, len(posVarTableau))
	fmt.Println("\033[0m") 

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
		
		Simplexe(tableau, tabConst,tabVar,IncrementalCoef,IncrementalAff,posVarTableau,bland,PosConst,alphaTab)

	}
	if x == 16 {
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
    	
        tabConst, tableau, tabVar = addAllConst(tabExe, tableau, tabConst, tabVar)
        fmt.Println("tableau = ",tableau)
        fmt.Println("tabConst = ",tabConst)
        fmt.Println("tabVar = ",tabVar)
		channel := make(chan bAndB)

	//creation tableau des affectations : taille = nombre de ligne + nombre de colonnes
	alphaTab := createAlphaTab(tableau, tabVar)
	//tableau qui nous donne la postion des variables dans le tableau alphaTab
	var posVarTableau = CreatePosVarTableau(tableau, tabVar)
	var bland = make([]string, len(posVarTableau))
	fmt.Println("\033[0m") 

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
	
		a,b,c,Incremental_Coef,Incremental_Aff,posV,rBland,posC:=Simplexe(tableau,tabConst,tabVar,IncrementalCoef,IncrementalAff,posVarTableau,bland,PosConst,alphaTab)
		fmt.Println(Branch_bound(a,b,c, tableau, tabConst, channel,Incremental_Coef,Incremental_Aff,posV,rBland,posC))
	}
}

/** 
 * This function create the map `alpha_tab` which associate the name of the variable and his alpha value.
 * It takes the following parameters:
 *   - `tab_coef`, the matrice with the normalized inequations
 *   - `tab_nom_var`, an array of the system's starting variable
 * It returns the create map 
 **/
func createAlphaTab(tab_coef [][]*big.Rat, tab_nom_var []string) map[string]*big.Rat{
	alpha_tab := make(map[string]*big.Rat)
	//Creation variable d'ecart
	for i := 0; i < len(tab_coef); i++ {
		alpha_tab[fmt.Sprint("e", i)] = new(big.Rat)
	}
	//Creation variable initial
	if len(tab_nom_var) == 0 {
	    for i := 0; i < len(tab_coef[0]); i++ {
		alpha_tab[fmt.Sprint("x", i)] = new(big.Rat)
	    }
    } else {
        for i := 0; i < len(tab_coef[0]); i++ {
            alpha_tab[tab_nom_var[i]] = new(big.Rat)
	    }
    }
	return alpha_tab
}

/** 
 * This function create the array which contains the position each variable in the matrice.
 * It takes the following parameters:
 *   - `tab_coef`, the matrice with the normalized inequations
 *   - `tab_nom_var`, an array of the system's starting variable
 * It returns the create array 
 **/
func CreatePosVarTableau(tab_coef [][]*big.Rat, tab_nom_var[]string) []string{
	//len(tab_coef[0]) correspond au nombre de ligne soit au nombre de variable d'ecart
	//len(tab_coef) correspond au nombre de colonne soit au nombre de variable initial
	var pos_var_tab = make([]string, len(tab_coef[0])+len(tab_coef))
	taille_tab := len(tab_coef)
	for i := 0; i < lenTab + len(tab_coef[0]); i++ {
		if i < taille_tab {
			pos_var_tab[i] = fmt.Sprint("e", i)
		} else {
		    if len(tab_nom_var) == 0 {
		        pos_var_tab[i] = fmt.Sprint("x", i-taille_tab)
		    } else {
		        pos_var_tab[i] = tab_nom_var[i-taille_tab]
		    }
		}
	}
	
	return pos_var_tab
}

//Remplace = par >= avec bool pour savoir si ce n'étais pas déja une inequation
func remplaceEgal(eq string)(string, bool){
	tab_ascii := []rune(eq)
	for index, ascii := range tab_ascii {
		// 61 -> = | 60 -> < | 62 -> >
    	if index > 1 && ascii == 61 && eq[index-1] != 60 && eq[index-1] != 62 {
    		return string(append(tab_ascii[:index],append([]rune{62}, tab_ascii[index:]...)...)), true
    	}
	}
	return eq, false
}

//Retourne le négatif d'une inéquation
func negEq(eq string)(string){
	tab_ascii := []rune(eq)

	//gestion du premier signe
	if tab_ascii[0] != 45 {
		if tab_ascii[0] >= 48 && tab_ascii[0] <= 57 {
			tab_ascii = append([]rune{43} ,tab_ascii...)
		}else {
			tab_ascii = append([]rune{43, 32} ,tab_ascii...)
		}
	} else {
		tab_ascii = tab_ascii[2:]
	}

	//gestion du dernier signe
	for i:=len(tab_ascii)-1 ; i > 0 ; i--{
		if tab_ascii[i-1] == 32 && tab_ascii[i] != 45 {
			tab_ascii = append(tab_ascii[:i], append([]rune{45}, tab_ascii[i:]...)...)
			i = -1
		} else if tab_ascii[i-1] == 32 && tab_ascii[i] == 45 {
			tab_ascii = append(tab_ascii[:i], tab_ascii[i:]...)
			i = -1
		}
	}

	//gestion des autres signe 
	for index, ascii := range tab_ascii {
		// 43 -> +
    	if ascii == 43 {
    		tab_ascii[index] = 45
    	// 45 -> -
    	} else if ascii == 45 {
    		tab_ascii[index] = 43
    	}
	}

	return string(tab_ascii)
}

//Retourne les inequations correspondante au equation d'entré
func getIneq(eq string)([]string){
	ineq, verif := remplaceEgal(eq)
	if verif {
		return []string{ineq, negEq(ineq)}
	} else {
		return []string{ineq}
	}
}