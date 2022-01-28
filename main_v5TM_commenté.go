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
)

func main() {
    // Creation du tableau de coeff et du tableau de contraintes
    var tableau = make([][]float64,0) 
    var tabConst = make([]float64,0)

    tableaux("20 r - x + y -18 z <= 8", tableau, tabConst)
}


// fonction qui prend en parametre une equation et qui implemente les tableaux de coeff et de contraintes
func tableaux(eq string, tableau [][]float64, tabConst []float64) {
    // on split la string avec les espaces, ce qui nous donne un tableaux avec tous les elements de la string
    tabEle := strings.Split(eq, " ")
    // tableau qui va contenir les coefficients de l'equation
    var ligneEq []float64
    // indique notre position dans le tableau
    posTab := 0
    for i := 0; i < len(tabEle)-2; i++ {
        ligneEq = append(ligneEq, 1.0)
        // nous permet de savoir si notre caractere est un chiffre
        re := regexp.MustCompile(`[0-9]`)
        isFig := re.FindString(tabEle[i])
        
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

    // On ajoute notre ligne de coeff dans notre tableau de coeff
    tableau = append(tableau, ligneEq[0:posTab])
    fmt.Println(tableau)
    
    // On ajoute la contraintes dans le tableau de contraintes en l'a convertissant d'abord
    lastEle := tabEle[len(tabEle)-1]
    lastEleC,_ := strconv.ParseFloat(lastEle, 64)
    tabConst = append(tabConst, lastEleC)
    fmt.Println(tabConst)
}


