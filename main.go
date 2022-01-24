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
    "strconv"
)

func main() {
    // on definit plusieurs equation et on les split
    var pluEqu = " x+y = 5, z+x >=3"
    split := strings.Split(pluEqu, ",")
    fmt.Println("split = ",split)
    
    // on calcul le nbr d'equations
    var nbrEq = len([]string(split))
    println("Length of the string is :", nbrEq)
    
    var equation = "2x+4y <= z x+2y = 5"
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
    
    // creation de la map qui initialise chaque variable à 1
    m := make(map[string]float64)
    var lenM = len([]string(tableauVar))
    for cpt := 0; cpt < lenM; cpt++ {
        m[tableauVar[cpt]] = 0
    }
    fmt.Println("map:", m)

    fmt.Println("equation = ",equation)
    for cpt := 0; cpt < lenM; cpt++ {
        findVar := regexp.MustCompile(tableauVar[cpt])
        loc := findVar.FindStringIndex(equation)
        loc1 := loc[0]-1
        loc2 := loc[1]-1
        //fmt.Println(loc1,loc2)
        valBefore := equation[loc1:loc2]
        //fmt.Println("valeur avant = ",valBefore)
        
        //on converti valBefore en float64
        valBeforeF, err := strconv.ParseFloat(valBefore, 64)
    	if err == nil {
    	    fmt.Println(valBeforeF, err)
    	}
        
        if valBefore != " " {
            m[tableauVar[cpt]] = m[tableauVar[cpt]] + valBeforeF
        }
    }
    fmt.Println("map:", m)

}

