/******************************************************************************
Welcome to GDB Online.
GDB online is an online compiler and debugger tool for C, C++, Python, Java, PHP, Ruby, Perl,
C#, VB, Swift, Pascal, Fortran, Haskell, Objective-C, Assembly, HTML, CSS, JS, SQLite, Prolog.
Code, Compile, Run and Debug online from anywhere in world.

*******************************************************************************/
package main

import (
    "fmt"
    "strconv"
    "strings"
    //"reflect"
    "regexp"
    "bufio"
    "os"
)

func main() {
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
        fmt.Println("tabConst =",tabConst)
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
