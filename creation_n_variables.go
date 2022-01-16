package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(vari(10))
	k := vari(10)
	fmt.Println("test si e1 existe :",k[0].nom)
	k[0].nom="toto"
	k[0].affectation=-2.5
	fmt.Println("test si e1 s'appelle désormais toto :",k[0].nom,"\ntest si l'affectation de toto vaut -2.5 et non 0 :", k[0].affectation)	

}


type MaVariable struct{

	 nom string
	 affectation float64
}


func vari(n int) []MaVariable{
	var M = make([]MaVariable,0)
	for i :=0;i<n;i++{
			M = append(M,MaVariable{ "e" + strconv.Itoa(i+1), 0 })
		
		}
return M
}