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
    fmt.Println(split)
    
    // on calcul le nbr d'equations
    var nbrEq = len([]string(split))
    println("Length of the string is :", nbrEq)
    
    var equation = "x+y <= 5"
    // tableau de variables
    re := regexp.MustCompile(`[a-z]`)
    vars := re.FindAll([]byte(equation), -1)
    fmt.Println(vars)
    
    var tableauVar []string
    var lent = len([][]byte(vars))
    for cpt := 0; cpt < lent; cpt++ {
        s := string([]byte(vars[cpt]))
        tableauVar = append(tableauVar,s)
    }
    fmt.Println(tableauVar)
    
    
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
