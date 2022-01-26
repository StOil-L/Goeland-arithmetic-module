/******************************************************************************
Welcome to GDB Online.
GDB online is an online compiler and debugger tool for C, C++, Python, Java, PHP, Ruby, Perl,
C#, VB, Swift, Pascal, Fortran, Haskell, Objective-C, Assembly, HTML, CSS, JS, SQLite, Prolog.
Code, Compile, Run and Debug online from anywhere in world.

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
    var pluEqu = " x+3y <= 2, 3x-4y <=4, 8x + y <=0"
    split := strings.Split(pluEqu, ",")
    fmt.Println("split = ",split)
    
    // on calcul le nbr d'equations
    var nbrEq = len(split)
    println("Length of the string is :", nbrEq)
    
    var equation = "2x+4y <= z"
    // tableau de variables
    reV := regexp.MustCompile(`[a-z]`)
    vars := reV.FindAll([]byte(equation), -1)
    fmt.Println("vars = ",vars)
    
    // tableau de coeff
    reC := regexp.MustCompile(`[0-9]`)
    coeff := reC.FindAll([]byte(equation), -1)
    fmt.Println("coeff = ",coeff)
    
    // tableau de coeff
    reTEST := regexp.MustCompile(`[a-z0-9]`)
    Test := reTEST.FindAll([]byte(equation), -1)
    fmt.Println("Test = ",Test)
    // boucles pour convertir nos variables et coeff qui sont en code ascii en string
    var tableauTest []string
    var lentest = len(Test)
    for cpt := 0; cpt < lentest; cpt++ {
        s := string([]byte(Test[cpt]))
        tableauTest = append(tableauTest,s)
    }
    fmt.Println("tableauTest =",tableauTest)
    
    // boucles pour convertir nos variables et coeff qui sont en code ascii en string
    var tableauVar []string
    var lenTV = len(vars)
    for cpt := 0; cpt < lenTV; cpt++ {
        sV := string([]byte(vars[cpt]))
        tableauVar = append(tableauVar,sV)
    }
    fmt.Println("tableauVar =",tableauVar)
    
    // creation de la map qui initialise chaque variable à 1
    m := make(map[string]float64)
    var lenM = len(tableauVar)
    for cpt := 0; cpt < lenM; cpt++ {
        m[tableauVar[cpt]] = 1
    }
    fmt.Println("map:", m)

    fmt.Println("equation = ",equation)
    for cpt := 0; cpt < lenM; cpt++ {
        findVar := regexp.MustCompile(tableauVar[cpt])
        loc := findVar.FindStringIndex(equation)
        fmt.Println(loc)
        loc1 := loc[0]-1
        loc2 := loc[1]-1
        fmt.Println(loc1,loc2)
        valBefore := equation[loc1:loc2]
        //fmt.Println("valeur avant = ",valBefore)
        
        //on converti valBefore en float64
        valBeforeF, err := strconv.ParseFloat(valBefore, 64)
    	if err == nil {
    	    fmt.Println(valBeforeF, err)
    	}
        
        if valBefore != " " {
            m[tableauVar[cpt]] = valBeforeF
        }
    }
    fmt.Println("map:", m)

}
