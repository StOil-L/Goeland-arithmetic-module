package main

import (
	"fmt"
  "math/rand"
  "time"
  "math/big"
)

// Returns an int >= min, < max
func randomInt(min, max int) int {
return min + rand.Intn(max-min)
}

func generateurtest (x int)(){

  rand.Seed(time.Now().UnixNano())
  
  if x == 17 {
    
    fmt.Println("Nombre de variables :")
    var y int
	fmt.Scanln(&y)
     fmt.Println("Nombre de contraintes :")
    var w int
	fmt.Scanln(&w)
    
    var taille int = (y*w)+w
    TableaAlea := make([]int64,taille)

    
    for i := 0;i < taille;i++{
      TableaAlea[i] = int64(randomInt(-10, 11))
    }
    

  var tab_var = make([]string,0)	
	var incremental_coef = make([]*big.Rat, 0)
	var incremental_aff= make([]*big.Rat,0) 
    var tableau = make([][]*big.Rat,w) 
    var tab_cont = make([]*big.Rat,w) 
    var cpt int = 0

    for i :=0;i < w;i++{
      for j :=0; j < y;j++{
        tableau[i] = append(tableau[i],big.NewRat(TableaAlea[cpt],1))
        cpt++
      }
      tab_cont[i] = big.NewRat(TableaAlea[i+(taille-w)],1)
    }
    fmt.Println(tableau)
    fmt.Println(tab_cont)


    //creation tableau des affectations : taille = nombre de ligne + nombre de colonnes
	alpha_tab := create_alpha_tab(tableau, tab_var)
	//tableau qui nous donne la postion des variables dans le tableau alpha_tab
	var pos_var_tab = create_pos_var_tab(tableau, tab_var)
	var bland = make([]string, len(pos_var_tab))
	fmt.Println("\033[0m") 

	for i:=0;i<len(tableau);i++ {
		bland[i+len(tableau[0])]=pos_var_tab[i]
	}
	
	for i:=0;i<len(tableau[0]);i++{
		bland[i]=pos_var_tab[i+len(tableau)]
	}
	
	var PosConst = make([]int, len(tab_cont))
	for i := 0; i<len(tab_cont);i++{
		PosConst[i]=i
	}



    system := info_system{tab_coef: tableau, tab_cont: tab_cont, tab_nom_var: tab_var,
			pos_var_tab: pos_var_tab, bland: bland, alpha_tab: alpha_tab,
			incremental_coef: incremental_coef, incremental_aff: incremental_aff}
    
    var ApelleSimplexe info_system
    var alphatab map[string]*big.Rat
    Simplexe(system)
    ApelleSimplexe,_ = Simplexe(system)
    alphatab = ApelleSimplexe.alpha_tab
    //fmt.Println("\n ",alphatab)
    var verite bool = true

	for i := 0;i < w;i++ {
	  concatenated := fmt.Sprint("e",i)
    var cpt int = alphatab[concatenated].Cmp(tab_cont[i])
    if cpt == -1 {
      verite = false
    }
    //fmt.Println(tab_cont[i])
	  //fmt.Println(alphatab[concatenated])
	}
    if verite {
       fmt.Println("Toutes les contraintes sont respectées")
    }else{
      fmt.Println("Au moins une contrainte n'est pas respectée")
    }
    
  }
}