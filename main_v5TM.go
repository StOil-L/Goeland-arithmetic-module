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
    //var tableau [][]float64
    //var tabConst []float64
    //element := "-15"
    //fmt.Println(strconv.ParseFloat(element, 64))
    var tableau = make([][]float64,0) 
    var tabConst = make([]float64,0)
    tableaux("-2 x +4 y <= 8", tableau, tabConst)
}



func tableaux(eq string, tableau [][]float64, tabConst []float64) {
    //var equation = "-2 x +4 y <= 8"
    //[-15, x, +5, y, =, 20]
    tabEle := strings.Split(eq, " ")
    var ligneEq []float64
    posTab := 0
    for i := 0; i < len(tabEle)-2; i++ {
        ligneEq = append(ligneEq, 1.0)

        //fmt.Println(tabEle[i]," est du Type ",reflect.TypeOf(tabEle[i]).String())
        re := regexp.MustCompile(`[0-9]`)
        isChar := re.FindString(tabEle[i])

        if isChar != "" {
            fmt.Println("tabEle[i]",tabEle[i])
            //fmt.Println(strconv.ParseFloat(tabEle[i], 64))
            conv,_ := strconv.ParseFloat(tabEle[i], 64)
            //fmt.Println("conv",conv)
            //fmt.Println("error",error)
            //ligneEq[posTab] = strconv.ParseFloat(tabEle[i], 64)
            ligneEq[posTab] = conv
            //fmt.Println("type = ",reflect.TypeOf(tabEle[i]).String())
        } else {
            posTab += 1
        }
    }
    //fmt.Println("ligneEq",ligneEq)
    fmt.Println("postab = ", posTab)
    tableau = append(tableau, ligneEq[0:posTab])
    fmt.Println(tableau)
    //tabConst = append(tabConst, tabEle[len(tabEle)-1])
}

