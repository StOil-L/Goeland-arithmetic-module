package main

import (
	"fmt"
  "math/big"
  "math"
	// "bibli"       pour rajouter une bibli tu met son nom entre double quote sans ','

)



func main() {

  //Création de nombre Rationnel grâce à la librairie big

  Test1 := big.NewRat(1,3)
  // Rationnel = 1/3

  Test2 := big.NewRat(1,7)
  // Rationnel = 1/7

  TestInt := big.NewRat(1,1)
  // Rationnel = 1/1 donc c'est un entier

  TestInt2 := big.NewRat(1,4)
  // Rationnel = 1/1 donc ce n'est pas un entier

  TestMulti  := big.NewRat(0,1)
  TestAddi := big.NewRat(0,1)
  //Je crée 2 variables Rationnel que j'initialise à 0 car je ne sais pas encore les créer sans les initialiser

  TestMulti.Mul(Test1, Test2)
  //Multiplie 2 Rat et enregistre le resultat dans TestMulti

  TestAddi.Add(Test1,Test2)
  //Additionne 2 Rat et enregistre le resultat dans TestAddi

  TestInt.IsInt()
  TestInt2.IsInt()
  //Renvoie Vrai si TestInt est un entier et Faux si il ne l'est pas

	fmt.Println(TestMulti)
  fmt.Println(TestAddi)
  if TestInt.IsInt() {
    fmt.Println("C'est un entier")
  }
  if TestInt2.IsInt() {
    fmt.Println("C'est un entier")
  }
          
  z2 := new(big.Rat).SetFloat64(1.25)   // z2 := 5/4
  z3:=new(big.Rat).SetFloat64(4.2)
  c:=new(big.Rat)
  d:=new(big.Rat)
  fmt.Println("c: ", c)
 
  var tab =make([]big.Rat,0)
  tab=append(tab,*z2)
  tab=append(tab,*z3)
  tab=append(tab,*d)
  fmt.Println(c.Add(z2,z3))
  fmt.Println(z2,z3,c)
  tab=append(tab,*c)

  fmt.Println(tab)
  e:=new(big.Rat)
  fmt.Println(e.Add(&tab[0],&tab[1]))

  f:=new(big.Rat).SetFloat64(math.Sqrt(2))
  fmt.Println(f)
}

