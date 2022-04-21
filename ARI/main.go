package main

import (
	typing "ARI/polymorphism"
	ari "ARI/rules"
	"fmt"
	"ARI/types"
)

var tInt typing.TypeHint
var tRat typing.TypeHint
var tProp typing.TypeHint

func main() {
	fmt.Printf("Exemples de problèmes d'arithmétique\n")

	// Initialisation des types de TPTP
	typing.Init()
	ari.Init()

	// Déclaration des types nécessaires pour les tests
	tInt = typing.MkTypeHint("int")
	tRat = typing.MkTypeHint("rat")
	tProp = typing.MkTypeHint("o")

	// Tests
	TestInt1()
	TestRat1()
	TestNormalizationEQ()
	TestNormalizationLess()
	TestNormalizationGreater()
	TestNormalizationNotInf()
	TestNormalizationNotSup()
	TestNormalizationNotInfEq()
	TestNormalizationNotSupEq()
	TestParserRemainder_f()



}

/* Tests INT */

/* Test int 1
* tff(sum_1_2_3,conjecture,
*    ( $sum(1,2) = 3 )).
**/
func TestInt1() {
	fmt.Println(" -------- TEST 1 -------- ")
	fmt.Println(" 1 + 2 = 3")
	un := types.MakerConst(types.MakerId("1"), tInt)
	deux := types.MakerConst(types.MakerId("2"), tInt)
	trois := types.MakerConst(types.MakerId("3"), tInt)
	sum := types.MakeFun(types.MakerId("sum"), []types.Term{un, deux}, typing.GetTypeScheme("sum", typing.MkTypeCross(tInt, tInt)))
	eq := types.MakePred(types.Id_eq, []types.Term{sum, trois}, typing.MkTypeArrow(typing.MkTypeCross(tInt, tInt), tProp))
	fmt.Printf("%v\n", eq.ToString())
}
/* Tests RAT */

/* Test rat 1
* tff(x_>_1/4,conjecture,
*    ! [X: $int] :  $greater(X,1/4)).
**/
func TestRat1() {
	fmt.Println(" -------- TEST 2 -------- ")
	fmt.Println(" X > 1/4 ")
	un_quart := types.MakerConst(types.MakerId("1/4"), tRat)
	x := types.MakeMeta(0, "X", -1, tRat)
	greater := types.MakePred(types.MakerId("greater"), []types.Term{x, un_quart}, typing.GetTypeScheme("greater", typing.MkTypeCross(tRat, tRat)))
	fmt.Printf("%v\n", greater.ToString())
}

func TestNormalizationEQ() {
	fmt.Println(" -------- TEST 3 -------- ")
	fmt.Println(" 3 = 3")
	trois := types.MakerConst(types.MakerId("3"), tInt)
	eq := types.MakePred(types.Id_eq, []types.Term{trois, trois}, typing.MkTypeArrow(typing.MkTypeCross(tInt, tInt), tProp))
	fmt.Printf("%v\n", eq.ToString())
	fmt.Printf("On applique la règle de normalisation (égalité)\n")
	res := ari.NormalizationRule(eq)
	for _, result_rule := range res {
		fmt.Printf("Resultat : %v\n", result_rule.ToString())
	}
}

func TestNormalizationNotEQ() {
	fmt.Println(" -------- TEST 4 -------- ")
	fmt.Println(" 3 != 2")
	deux := types.MakerConst(types.MakerId("2"), tInt)
	trois := types.MakerConst(types.MakerId("3"), tInt)
	eq := types.MakePred(types.Id_neq, []types.Term{trois, deux}, typing.MkTypeArrow(typing.MkTypeCross(tInt, tInt), tProp))
	fmt.Printf("%v\n", eq.ToString())
	fmt.Printf("On applique la règle de normalisation (négalité)\n")
	res := ari.NormalizationRule(eq)
	for _, result_rule := range res {
		fmt.Printf("Resultat : %v\n", result_rule.ToString())
	}
}


func TestNormalizationLess() {
	fmt.Println(" -------- TEST 5 -------- ")
	fmt.Println(" 2 < 3")
	deux := types.MakerConst(types.MakerId("2"), tInt)
	trois := types.MakerConst(types.MakerId("3"), tInt)
	ineq := types.MakePred(types.MakerId("less"), []types.Term{deux, trois}, typing.MkTypeArrow(typing.MkTypeCross(tInt, tInt), tProp))

	fmt.Printf("%v\n", ineq.ToString())
	fmt.Printf("On applique la règle de normalisation (inf)\n")
	res := ari.NormalizationRule(ineq)
	for _, result_rule := range res {
		fmt.Printf("Resultat : %v\n", result_rule.ToString())
	}
}

func TestNormalizationGreater() {
	fmt.Println(" -------- TEST 6 -------- ")
	fmt.Println(" 3 > 2")
	deux := types.MakerConst(types.MakerId("2"), tInt)
	trois := types.MakerConst(types.MakerId("3"), tInt)
	ineq := types.MakePred(types.MakerId("greater"), []types.Term{trois, deux}, typing.MkTypeArrow(typing.MkTypeCross(tInt, tInt), tProp))

	fmt.Printf("%v\n", ineq.ToString())
	fmt.Printf("On applique la règle de normalisation (inf)\n")
	res := ari.NormalizationRule(ineq)
	for _, result_rule := range res {
		fmt.Printf("Resultat : %v\n", result_rule.ToString())
	}
}



func TestNormalizationNotInf() {
	fmt.Println(" -------- TEST 7 -------- ")
	fmt.Println(" ! (3 < 2)")
	deux := types.MakerConst(types.MakerId("2"), tInt)
	trois := types.MakerConst(types.MakerId("3"), tInt)
	Ineq := types.MakePred(types.MakerId("¬less"), []types.Term{trois, deux}, typing.MkTypeArrow(typing.MkTypeCross(tInt, tInt), tProp))
	Ineq2 := types.MakePred(types.MakerId("less"), []types.Term{trois, deux}, typing.MkTypeArrow(typing.MkTypeCross(tInt, tInt), tProp))
	//pour montrer que Not(Pred) == Pred : notInf
	NotIneq := types.RefuteForm(Ineq2)
	fmt.Printf("%v\n", Ineq.ToString())
	fmt.Printf("%v\n", NotIneq.ToString())
	fmt.Printf("On applique la règle de normalisation (negInf))\n")
	res := ari.NormalizationNegRule(NotIneq)
	for _, result_rule := range res {
		fmt.Printf("Resultat : %v\n", result_rule.ToString())
	}
}

func TestNormalizationNotSup() {
	fmt.Println(" -------- TEST 8 -------- ")
	fmt.Println(" ! (2 > 3)")
	deux := types.MakerConst(types.MakerId("2"), tInt)
	trois := types.MakerConst(types.MakerId("3"), tInt)
	Ineq := types.MakePred(types.MakerId("greater"), []types.Term{deux, trois}, typing.MkTypeArrow(typing.MkTypeCross(tInt, tInt), tProp))
	NotIneq := types.RefuteForm(Ineq)
	fmt.Printf("%v\n", NotIneq.ToString())
	fmt.Printf("On applique la règle de normalisation (negSup))\n")
	res := ari.NormalizationNegRule(NotIneq)
	for _, result_rule := range res {
		fmt.Printf("Resultat : %v\n", result_rule.ToString())
	}
}

func TestNormalizationNotInfEq() {
	fmt.Println(" -------- TEST 9 -------- ")
	fmt.Println(" ! (3 <= 2)")
	deux := types.MakerConst(types.MakerId("2"), tInt)
	trois := types.MakerConst(types.MakerId("3"), tInt)
	Ineq := types.MakePred(types.MakerId("lessereq"), []types.Term{trois, deux}, typing.MkTypeArrow(typing.MkTypeCross(tInt, tInt), tProp))
	NotIneq := types.RefuteForm(Ineq)
	fmt.Printf("%v\n", NotIneq.ToString())
	fmt.Printf("On applique la règle de normalisation (negInfEq))\n")
	res := ari.NormalizationNegRule(NotIneq)
	for _, result_rule := range res {
		fmt.Printf("Resultat : %v\n", result_rule.ToString())
	}
}

func TestNormalizationNotSupEq() {
	fmt.Println(" -------- TEST 10 -------- ")
	fmt.Println(" ! (2 >= 3)")
	deux := types.MakerConst(types.MakerId("2"), tInt)
	trois := types.MakerConst(types.MakerId("3"), tInt)
	Ineq := types.MakePred(types.MakerId("greatereq"), []types.Term{deux, trois}, typing.MkTypeArrow(typing.MkTypeCross(tInt, tInt), tProp))
	NotIneq := types.RefuteForm(Ineq)
	fmt.Printf("%v\n", NotIneq.ToString())
	fmt.Printf("On applique la règle de normalisation (negSupEq))\n")
	res := ari.NormalizationNegRule(NotIneq)
	for _, result_rule := range res {
		fmt.Printf("Resultat : %v\n", result_rule.ToString())
	}
}

func TestParserRemainder_f() {
	fmt.Println(" -------- TEST 11 -------- ")
	fmt.Println("  (8%3)")
	huit := types.MakerConst(types.MakerId("8"), tInt)
	trois := types.MakerConst(types.MakerId("3"), tInt)
	modulo:=types.MakeFun(types.MakerId("remainder_f"), []types.Term{huit, trois}, typing.GetTypeScheme("remainder_f", typing.MkTypeCross(tInt, tInt)))		
	res,_ := ari.FunToInt(modulo)
	fmt.Printf("8 modulo floor 3 = %d \n",res)
}
func TestParserUminus() {
	fmt.Println(" -------- TEST 12 -------- ")
	fmt.Println("  (8%3)")
	huit := types.MakerConst(types.MakerId("8"), tInt)
	trois := types.MakerConst(types.MakerId("3"), tInt)
	modulo:=types.MakeFun(types.MakerId("remainder_f"), []types.Term{huit, trois}, typing.GetTypeScheme("remainder_f", typing.MkTypeCross(tInt, tInt)))		
	res,_ := ari.FunToInt(modulo)
	fmt.Printf("8 modulo floor 3 = %d \n",res)
}