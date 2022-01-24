/******************************************************************************

                            Online Go Lang Compiler.
                Code, Compile, Run and Debug Go Lang program online.
Write your code in this editor and press "Run" button to execute it.

*******************************************************************************/



package main
import (
    "fmt"
    "strings"
    "regexp"
)
 
func main() {
    // on definit plusieurs equation et on les split
    var pluEqu = " x+y = 5, z+x >=3"
    split := strings.Split(pluEqu, ",")
    fmt.Println("split = ",split)
    
    // on calcul le nbr d'equations
    var nbrEq = len([]string(split))
    println("Length of the string is :", nbrEq)
    
    var equation = "2x+4y <= 1z"
    // tableau de variables
    reV := regexp.MustCompile(`[a-z]`)
    vars := reV.FindAll([]byte(equation), -1)
    fmt.Println("vars = ",vars)
    
    // tableau de coeff
    reC := regexp.MustCompile(`[0-9]`)
    coeff := reC.FindAll([]byte(equation), -1)
    fmt.Println("coeff = ",coeff)
    
    // boucles pour convertir nos variables et coeff qui sont en code ascii en string
    var tableauVar []string
    var lenTV = len([][]byte(vars))
    for cpt := 0; cpt < lenTV; cpt++ {
        sV := string([]byte(vars[cpt]))
        tableauVar = append(tableauVar,sV)
    }
    fmt.Println("tableauVar =",tableauVar)
    
    var tableauCoeff []string
    var lenTC = len([][]byte(coeff))
    for cpt := 0; cpt < lenTC; cpt++ {
        sC := string([]byte(coeff[cpt]))
        tableauCoeff = append(tableauCoeff,sC)
    }
    fmt.Println("tableauCoeff =",tableauCoeff)
    
    // creation de la map qui contient la variable en clé et sa valeur en valeur
    m := make(map[string]string)
    var lenM = len([]string(tableauVar))
    for cpt := 0; cpt < lenM; cpt++ {
        m[tableauVar[cpt]] = tableauCoeff[cpt]
    }
    fmt.Println("map:", m)

    
    //var equation = "x+y <= 5"
    // va aller dans la boucle for
    if strings.Contains(equation, ">=") {
        fmt.Println(">=")
    } else if strings.Contains(equation, "<=") {
        fmt.Println("<=")
    } else if strings.Contains(equation, "=") {
        fmt.Println("=")
    } else {
        fmt.Println("erreur")
    }
    
    

}
