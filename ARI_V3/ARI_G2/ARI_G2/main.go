package main

import (
	typing "ARI/polymorphism"
	ari "ARI/rules"
	"ARI/types"
	"fmt"
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
	// Exemple de création de tests
	TestInt()
	TestRat()

	// Tests règle constantes
	TestConstEq1()
	TestConstEq2()
	TestConstLess1()
	TestConstLess2()
	TestConstGreater1()
	TestConstGreater2()
	TestConstLessEq1()
	TestConstLessEq2()
	TestConstLessEq3()
	TestConstGreaterEq1()
	TestConstGreaterEq2()
	TestConstGreaterEq3()
	TestConstIsInt1()
	TestConstIsInt2()
	TestConstIsRat1()
	TestConstIsRat2()

	// Tests règles de normalisation
	TestNormalizationEQ()
	TestNormalizationLess()
	TestNormalizationGreater()
	TestNormalizationNotInf()
	TestNormalizationNotSup()
	TestNormalizationNotInfEq()
	TestNormalizationNotSupEq()
	TestParserRemainder_f()

	// Tests règles simplexe
	TestSimplexeRat1()
	TestSimplexeRat2()
	TestSimplexeRat()
	TestSimplexeSumRat()
	TestSimplexeBeaucoupRat()


}

/*** Test création de termes ***/
/* Tests INT */

/* Test int 1
* tff(sum_1_2_3,conjecture,
*    ( $sum(1,2) = 3 )).
**/
func TestInt() {
	fmt.Println(" -------- TEST 1 -------- ")
	fmt.Println(" 1 + 2 = 3")
	un := types.MakerConst(types.MakerId("1"), tInt)
	deux := types.MakerConst(types.MakerId("2"), tInt)
	trois := types.MakerConst(types.MakerId("3"), tInt)
	sum := types.MakeFun(types.MakerId("sum"), []types.Term{un, deux}, typing.GetTypeScheme("sum", typing.MkTypeCross(tInt, tInt)))
	p := types.MakePred(types.Id_eq, []types.Term{sum, trois}, typing.MkTypeArrow(typing.MkTypeCross(tInt, tInt), tProp))
	fmt.Printf("%v\n", p.ToString())
}

/* Tests RAT */

/* Test rat 1
* tff(x_>_1/4,conjecture,
*    ! [X: $int] :  $greater(X,1/4)).
**/
func TestRat() {
	fmt.Println(" -------- TEST 2 -------- ")
	fmt.Println(" X > 1/4 ")
	un_quart := types.MakerConst(types.MakerId("1/4"), tRat)
	x := types.MakeMeta(0, "X", -1, tRat)
	p := types.MakePred(types.MakerId("greater"), []types.Term{x, un_quart}, typing.GetTypeScheme("greater", typing.MkTypeCross(tRat, tRat)))
	fmt.Printf("%v\n", p.ToString())
}

/*** Tests règles const ***/
func TestConstEq1() {
	fmt.Println(" -------- TEST 3 -------- ")
	fmt.Println(" 3 = 3")
	trois := types.MakerConst(types.MakerId("3"), tInt)
	p := types.MakePred(types.Id_eq, []types.Term{trois, trois}, typing.MkTypeArrow(typing.MkTypeCross(tInt, tInt), tProp))
	fmt.Printf("%v\n", p.ToString())
	fmt.Printf("On applique la règle de const eq (resultat attendu : false)\n")
	res := ari.ApplyConstRule(p)
	fmt.Printf("Resultat : %v\n", res)
}

func TestConstEq2() {
	fmt.Println(" -------- TEST 4 -------- ")
	fmt.Println(" 3 = 4")
	trois := types.MakerConst(types.MakerId("3"), tInt)
	quatre := types.MakerConst(types.MakerId("4"), tInt)
	p := types.MakePred(types.Id_eq, []types.Term{trois, quatre}, typing.MkTypeArrow(typing.MkTypeCross(tInt, tInt), tProp))
	fmt.Printf("%v\n", p.ToString())
	fmt.Printf("On applique la règle de const eq (resultat attendu : true)\n")
	res := ari.ApplyConstRule(p)
	fmt.Printf("Resultat : %v\n", res)
}

func TestConstLess1() {
	fmt.Println(" -------- TEST 5 -------- ")
	fmt.Println(" 3 < 4")
	trois := types.MakerConst(types.MakerId("3"), tInt)
	quatre := types.MakerConst(types.MakerId("4"), tInt)
	p := types.MakePred(types.MakerId("less"), []types.Term{trois, quatre}, typing.MkTypeArrow(typing.MkTypeCross(tInt, tInt), tProp))
	fmt.Printf("%v\n", p.ToString())
	fmt.Printf("On applique la règle de const less (resultat attendu : false)\n")
	res := ari.ApplyConstRule(p)
	fmt.Printf("Resultat : %v\n", res)
}

func TestConstLess2() {
	fmt.Println(" -------- TEST 6 -------- ")
	fmt.Println(" 4 < 3")
	trois := types.MakerConst(types.MakerId("3"), tInt)
	quatre := types.MakerConst(types.MakerId("4"), tInt)
	p := types.MakePred(types.MakerId("less"), []types.Term{quatre, trois}, typing.MkTypeArrow(typing.MkTypeCross(tInt, tInt), tProp))
	fmt.Printf("%v\n", p.ToString())
	fmt.Printf("On applique la règle de const less (resultat attendu : true)\n")
	res := ari.ApplyConstRule(p)
	fmt.Printf("Resultat : %v\n", res)
}

func TestConstGreater1() {
	fmt.Println(" -------- TEST 7 -------- ")
	fmt.Println(" 4 > 3")
	trois := types.MakerConst(types.MakerId("3"), tInt)
	quatre := types.MakerConst(types.MakerId("4"), tInt)
	p := types.MakePred(types.MakerId("great"), []types.Term{quatre, trois}, typing.MkTypeArrow(typing.MkTypeCross(tInt, tInt), tProp))
	fmt.Printf("%v\n", p.ToString())
	fmt.Printf("On applique la règle de const  great (resultat attendu : false)\n")
	res := ari.ApplyConstRule(p)
	fmt.Printf("Resultat : %v\n", res)
}

func TestConstGreater2() {
	fmt.Println(" -------- TEST 8 -------- ")
	fmt.Println(" 3 > 4")
	trois := types.MakerConst(types.MakerId("3"), tInt)
	quatre := types.MakerConst(types.MakerId("4"), tInt)
	p := types.MakePred(types.MakerId("great"), []types.Term{trois, quatre}, typing.MkTypeArrow(typing.MkTypeCross(tInt, tInt), tProp))
	fmt.Printf("%v\n", p.ToString())
	fmt.Printf("On applique la règle de const great (resultat attendu : true)\n")
	res := ari.ApplyConstRule(p)
	fmt.Printf("Resultat : %v\n", res)
}

func TestConstLessEq1() {
	fmt.Println(" -------- TEST 9 -------- ")
	fmt.Println(" 3 <= 4")
	trois := types.MakerConst(types.MakerId("3"), tInt)
	quatre := types.MakerConst(types.MakerId("4"), tInt)
	p := types.MakePred(types.MakerId("lesseq"), []types.Term{trois, quatre}, typing.MkTypeArrow(typing.MkTypeCross(tInt, tInt), tProp))
	fmt.Printf("%v\n", p.ToString())
	fmt.Printf("On applique la règle de const less eq (resultat attendu : false)\n")
	res := ari.ApplyConstRule(p)
	fmt.Printf("Resultat : %v\n", res)
}

func TestConstLessEq2() {
	fmt.Println(" -------- TEST 10 -------- ")
	fmt.Println(" 4 <= 4")
	quatre := types.MakerConst(types.MakerId("4"), tInt)
	p := types.MakePred(types.MakerId("lesseq"), []types.Term{quatre, quatre}, typing.MkTypeArrow(typing.MkTypeCross(tInt, tInt), tProp))
	fmt.Printf("%v\n", p.ToString())
	fmt.Printf("On applique la règle de const less eq (resultat attendu : false)\n")
	res := ari.ApplyConstRule(p)
	fmt.Printf("Resultat : %v\n", res)
}

func TestConstLessEq3() {
	fmt.Println(" -------- TEST 11 -------- ")
	fmt.Println(" 4 <= 3")
	trois := types.MakerConst(types.MakerId("3"), tInt)
	quatre := types.MakerConst(types.MakerId("4"), tInt)
	p := types.MakePred(types.MakerId("lesseq"), []types.Term{quatre, trois}, typing.MkTypeArrow(typing.MkTypeCross(tInt, tInt), tProp))
	fmt.Printf("%v\n", p.ToString())
	fmt.Printf("On applique la règle de const less eq (resultat attendu : true)\n")
	res := ari.ApplyConstRule(p)
	fmt.Printf("Resultat : %v\n", res)
}

func TestConstGreaterEq1() {
	fmt.Println(" -------- TEST 12 -------- ")
	fmt.Println(" 4 >= 3")
	trois := types.MakerConst(types.MakerId("3"), tInt)
	quatre := types.MakerConst(types.MakerId("4"), tInt)
	p := types.MakePred(types.MakerId("greateq"), []types.Term{quatre, trois}, typing.MkTypeArrow(typing.MkTypeCross(tInt, tInt), tProp))
	fmt.Printf("%v\n", p.ToString())
	fmt.Printf("On applique la règle de const great eq (resultat attendu : false)\n")
	res := ari.ApplyConstRule(p)
	fmt.Printf("Resultat : %v\n", res)
}

func TestConstGreaterEq2() {
	fmt.Println(" -------- TEST 13 -------- ")
	fmt.Println(" 4 >= 4")
	quatre := types.MakerConst(types.MakerId("4"), tInt)
	p := types.MakePred(types.MakerId("greateq"), []types.Term{quatre, quatre}, typing.MkTypeArrow(typing.MkTypeCross(tInt, tInt), tProp))
	fmt.Printf("%v\n", p.ToString())
	fmt.Printf("On applique la règle de const great eq (resultat attendu : false)\n")
	res := ari.ApplyConstRule(p)
	fmt.Printf("Resultat : %v\n", res)
}

func TestConstGreaterEq3() {
	fmt.Println(" -------- TEST 14 -------- ")
	fmt.Println(" 3 >= 4")
	trois := types.MakerConst(types.MakerId("3"), tInt)
	quatre := types.MakerConst(types.MakerId("4"), tInt)
	p := types.MakePred(types.MakerId("greateq"), []types.Term{trois, quatre}, typing.MkTypeArrow(typing.MkTypeCross(tInt, tInt), tProp))
	fmt.Printf("%v\n", p.ToString())
	fmt.Printf("On applique la règle de const great eq (resultat attendu : true)\n")
	res := ari.ApplyConstRule(p)
	fmt.Printf("Resultat : %v\n", res)
}

func TestConstIsInt1() {
	fmt.Println(" -------- TEST 15 -------- ")
	fmt.Println(" isInt(3) ")
	trois := types.MakerConst(types.MakerId("3"), tInt)
	p := types.MakePred(types.MakerId("is_int"), []types.Term{trois}, typing.MkTypeArrow(tInt, tProp))
	fmt.Printf("%v\n", p.ToString())
	fmt.Printf("On applique la règle de const is_int (resultat attendu : false)\n")
	res := ari.ApplyConstRule(p)
	fmt.Printf("Resultat : %v\n", res)
}

func TestConstIsInt2() {
	fmt.Println(" -------- TEST 16 -------- ")
	fmt.Println(" isInt(3/2) ")
	trois_demi := types.MakerConst(types.MakerId("3/2"), tRat)
	p := types.MakePred(types.MakerId("is_int"), []types.Term{trois_demi}, typing.MkTypeArrow(tRat, tProp))
	fmt.Printf("%v\n", p.ToString())
	fmt.Printf("On applique la règle de const is_int (resultat attendu : true)\n")
	res := ari.ApplyConstRule(p)
	fmt.Printf("Resultat : %v\n", res)
}

func TestConstIsRat1() {
	fmt.Println(" -------- TEST 17 -------- ")
	fmt.Println(" isRat(3/2) ")
	trois := types.MakerConst(types.MakerId("3/2"), tRat)
	p := types.MakePred(types.MakerId("is_rat"), []types.Term{trois}, typing.MkTypeArrow(tRat, tProp))
	fmt.Printf("%v\n", p.ToString())
	fmt.Printf("On applique la règle de const is_rat (resultat attendu : false)\n")
	res := ari.ApplyConstRule(p)
	fmt.Printf("Resultat : %v\n", res)
}

func TestConstIsRat2() {
	fmt.Println(" -------- TEST 18 -------- ")
	fmt.Println(" isRat(3) ")
	trois_demi := types.MakerConst(types.MakerId("3"), tInt)
	p := types.MakePred(types.MakerId("is_rat"), []types.Term{trois_demi}, typing.MkTypeArrow(tInt, tProp))
	fmt.Printf("%v\n", p.ToString())
	fmt.Printf("On applique la règle de const is_rat (resultat attendu : true)\n")
	res := ari.ApplyConstRule(p)
	fmt.Printf("Resultat : %v\n", res)
}

/*** Tests règles de normalisation ***/

func TestNormalizationEQ() {
	fmt.Println(" -------- TEST 19 -------- ")
	fmt.Println(" 3 = 3")
	trois := types.MakerConst(types.MakerId("3"), tInt)
	eq := types.MakePred(types.Id_eq, []types.Term{trois, trois}, typing.MkTypeArrow(typing.MkTypeCross(tInt, tInt), tProp))
	fmt.Printf("%v\n", eq.ToString())
	fmt.Printf("On applique la règle de normalisation (égalité)\n")
	res := ari.ApplyNormalizationRule(eq)
	for _, result_rule := range res {
		fmt.Printf("Resultat : %v\n", result_rule.ToString())
	}
}

func TestNormalizationNotEQ() {
	fmt.Println(" -------- TEST 20 -------- ")
	fmt.Println(" 3 != 2")
	deux := types.MakerConst(types.MakerId("2"), tInt)
	trois := types.MakerConst(types.MakerId("3"), tInt)
	eq := types.MakePred(types.Id_neq, []types.Term{trois, deux}, typing.MkTypeArrow(typing.MkTypeCross(tInt, tInt), tProp))
	fmt.Printf("%v\n", eq.ToString())
	fmt.Printf("On applique la règle de normalisation (négalité)\n")
	res := ari.ApplyNormalizationRule(eq)
	for _, result_rule := range res {
		fmt.Printf("Resultat : %v\n", result_rule.ToString())
	}
}

func TestNormalizationLess() {
	fmt.Println(" -------- TEST 21 -------- ")
	fmt.Println(" 2 < 3")
	deux := types.MakerConst(types.MakerId("2"), tInt)
	trois := types.MakerConst(types.MakerId("3"), tInt)
	ineq := types.MakePred(types.MakerId("less"), []types.Term{deux, trois}, typing.MkTypeArrow(typing.MkTypeCross(tInt, tInt), tProp))

	fmt.Printf("%v\n", ineq.ToString())
	fmt.Printf("On applique la règle de normalisation (inf)\n")
	res := ari.ApplyNormalizationRule(ineq)
	for _, result_rule := range res {
		fmt.Printf("Resultat : %v\n", result_rule.ToString())
	}
}

func TestNormalizationGreater() {
	fmt.Println(" -------- TEST 22 -------- ")
	fmt.Println(" 3 > 2")
	deux := types.MakerConst(types.MakerId("2"), tInt)
	trois := types.MakerConst(types.MakerId("3"), tInt)
	ineq := types.MakePred(types.MakerId("greater"), []types.Term{trois, deux}, typing.MkTypeArrow(typing.MkTypeCross(tInt, tInt), tProp))

	fmt.Printf("%v\n", ineq.ToString())
	fmt.Printf("On applique la règle de normalisation (inf)\n")
	res := ari.ApplyNormalizationRule(ineq)
	for _, result_rule := range res {
		fmt.Printf("Resultat : %v\n", result_rule.ToString())
	}
}

func TestNormalizationNotInf() {
	fmt.Println(" -------- TEST 23 -------- ")
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
	res := ari.ApplyNormalizationRule(NotIneq)
	for _, result_rule := range res {
		fmt.Printf("Resultat : %v\n", result_rule.ToString())
	}
}

func TestNormalizationNotSup() {
	fmt.Println(" -------- TEST 24 -------- ")
	fmt.Println(" ! (2 > 3)")
	deux := types.MakerConst(types.MakerId("2"), tInt)
	trois := types.MakerConst(types.MakerId("3"), tInt)
	Ineq := types.MakePred(types.MakerId("greater"), []types.Term{deux, trois}, typing.MkTypeArrow(typing.MkTypeCross(tInt, tInt), tProp))
	NotIneq := types.RefuteForm(Ineq)
	fmt.Printf("%v\n", NotIneq.ToString())
	fmt.Printf("On applique la règle de normalisation (negSup))\n")
	res := ari.ApplyNormalizationRule(NotIneq)
	for _, result_rule := range res {
		fmt.Printf("Resultat : %v\n", result_rule.ToString())
	}
}

func TestNormalizationNotInfEq() {
	fmt.Println(" -------- TEST 25 -------- ")
	fmt.Println(" ! (3 <= 2)")
	deux := types.MakerConst(types.MakerId("2"), tInt)
	trois := types.MakerConst(types.MakerId("3"), tInt)
	Ineq := types.MakePred(types.MakerId("lessereq"), []types.Term{trois, deux}, typing.MkTypeArrow(typing.MkTypeCross(tInt, tInt), tProp))
	NotIneq := types.RefuteForm(Ineq)
	fmt.Printf("%v\n", NotIneq.ToString())
	fmt.Printf("On applique la règle de normalisation (negInfEq))\n")
	res := ari.ApplyNormalizationRule(NotIneq)
	for _, result_rule := range res {
		fmt.Printf("Resultat : %v\n", result_rule.ToString())
	}
}

func TestNormalizationNotSupEq() {
	fmt.Println(" -------- TEST 26 -------- ")
	fmt.Println(" ! (2 >= 3)")
	deux := types.MakerConst(types.MakerId("2"), tInt)
	trois := types.MakerConst(types.MakerId("3"), tInt)
	Ineq := types.MakePred(types.MakerId("greatereq"), []types.Term{deux, trois}, typing.MkTypeArrow(typing.MkTypeCross(tInt, tInt), tProp))
	NotIneq := types.RefuteForm(Ineq)
	fmt.Printf("%v\n", NotIneq.ToString())
	fmt.Printf("On applique la règle de normalisation (negSupEq))\n")
	res := ari.ApplyNormalizationRule(NotIneq)
	for _, result_rule := range res {
		fmt.Printf("Resultat : %v\n", result_rule.ToString())
	}
}

func TestParserRemainder_f() {
	fmt.Println(" -------- TEST 27 -------- ")
	fmt.Println("  (8%3)")
	huit := types.MakerConst(types.MakerId("8"), tInt)
	trois := types.MakerConst(types.MakerId("3"), tInt)
	modulo := types.MakeFun(types.MakerId("remainder_f"), []types.Term{huit, trois}, typing.GetTypeScheme("remainder_f", typing.MkTypeCross(tInt, tInt)))
	res, _ := ari.EvaluateFun(modulo)
	fmt.Printf("8 modulo floor 3 = %d \n", res)
}

/*** Tests règle simplexe ***/

/** Julie
* Là je crée un exemple pour l'appel à la fonction SimplexRule
* Il prend en paramètre le problème suivant :
* x <= 3/2
* x >= 2/3
* Il doit renvoyer une solution !
**/



func TestSimplexeRat1() {
	fmt.Println(" -------- TEST 27.5 -------- ")
	fmt.Println(" X = 3/2 et X = 2/3 ")
	x := types.MakerMeta("X", -1)
	trois_demi := types.MakerConst(types.MakerId("3/2"), tRat)
	deux_tiers := types.MakerConst(types.MakerId("2/3"), tRat)
	p1 := types.MakePred(types.Id_eq, []types.Term{x, trois_demi}, typing.MkTypeArrow(typing.MkTypeCross(tRat, tRat), tProp))
	p2 := types.MakePred(types.Id_eq, []types.Term{x, deux_tiers}, typing.MkTypeArrow(typing.MkTypeCross(tRat, tRat), tProp))
	systeme := []types.Form{p1, p2}
	found, solution := ari.ApplySimplexRule(systeme)
	fmt.Printf("Solution trouvée : %v = %v \n", found, solution.ToString())
}


func TestSimplexeRat2() {
	fmt.Println(" -------- TEST 28 -------- ")
	fmt.Println(" X <= 3/2 et Y >= 2/3 ")
	x := types.MakerMeta("X", -1)
	y := types.MakerMeta("Y", -1)
	trois_demi := types.MakerConst(types.MakerId("3/2"), tRat)
	deux_tiers := types.MakerConst(types.MakerId("2/3"), tRat)
	p1 := types.MakePred(types.MakerId("lesseq"), []types.Term{x, trois_demi}, typing.MkTypeArrow(typing.MkTypeCross(tRat, tRat), tProp))
	p2 := types.MakePred(types.MakerId("greateq"), []types.Term{y, deux_tiers}, typing.MkTypeArrow(typing.MkTypeCross(tRat, tRat), tProp))
	systeme := []types.Form{p1, p2}
	found, solution := ari.ApplySimplexRule(systeme)
	fmt.Printf("Solution trouvée : %v = %v \n", found, solution.ToString())
}

func TestSimplexeRat() {
	fmt.Println(" -------- TEST 28.5 -------- ")
	fmt.Println(" X <= 3/2 et Y >= 2/3  et  3/2 = Z et X = 2/3 ")
	x := types.MakerMeta("X", -1)
	y := types.MakerMeta("Y", -1)
	z := types.MakerMeta("Z", -1)
	trois_demi := types.MakerConst(types.MakerId("3/2"), tRat)
	deux_tiers := types.MakerConst(types.MakerId("2/3"), tRat)
	p1 := types.MakePred(types.MakerId("lesseq"), []types.Term{x, trois_demi}, typing.MkTypeArrow(typing.MkTypeCross(tRat, tRat), tProp))
	p2 := types.MakePred(types.MakerId("greateq"), []types.Term{y, deux_tiers}, typing.MkTypeArrow(typing.MkTypeCross(tRat, tRat), tProp))
	p3 := types.MakePred(types.Id_eq, []types.Term{ trois_demi, z}, typing.MkTypeArrow(typing.MkTypeCross(tRat, tRat), tProp))
	p4 := types.MakePred(types.Id_eq, []types.Term{x, deux_tiers}, typing.MkTypeArrow(typing.MkTypeCross(tRat, tRat), tProp))
	systeme := []types.Form{p1, p2, p3, p4}
	found, solution := ari.ApplySimplexRule(systeme)
	fmt.Printf("Solution trouvée : %v = %v \n", found, solution.ToString())
}

func TestSimplexeSumRat() {
	fmt.Println(" -------- TEST 28.75 -------- ")
	fmt.Println(" X + Y = 3")
	x := types.MakerMeta("X", -1)
	y := types.MakerMeta("Y", -1)
	trois_demi := types.MakerConst(types.MakerId("3/2"), tRat)
	sum := types.MakeFun(types.MakerId("sum"), []types.Term{x, y}, typing.GetTypeScheme("sum", typing.MkTypeCross(tRat, tRat)))
	p := types.MakePred(types.Id_eq, []types.Term{sum, trois_demi}, typing.MkTypeArrow(typing.MkTypeCross(tRat, tRat), tProp))
	systeme := []types.Form{p}
	found, solution := ari.ApplySimplexRule(systeme)
	fmt.Printf("Solution trouvée : %v = %v \n", found, solution.ToString())
}



func TestSimplexeBeaucoupRat() {
	fmt.Println(" -------- TEST 28.999 -------- ")
	fmt.Println(" ((X + Y)-Z)*K = 3")
	x := types.MakerMeta("X", -1)
	y := types.MakerMeta("Y", -1)
	z := types.MakerMeta("Z", -1)
	k := types.MakerMeta("K", -1)
	trois_demi := types.MakerConst(types.MakerId("3/2"), tRat)
	sum := types.MakeFun(types.MakerId("sum"), []types.Term{x, y}, typing.GetTypeScheme("sum", typing.MkTypeCross(tRat, tRat)))
	diff := types.MakeFun(types.MakerId("difference"), []types.Term{sum, z}, typing.GetTypeScheme("difference", typing.MkTypeCross(tRat, tRat)))
	prod := types.MakeFun(types.MakerId("product"), []types.Term{diff, k}, typing.GetTypeScheme("product", typing.MkTypeCross(tRat, tRat)))
	p := types.MakePred(types.Id_eq, []types.Term{prod, trois_demi}, typing.MkTypeArrow(typing.MkTypeCross(tRat, tRat), tProp))
	systeme := []types.Form{p}
	found, solution := ari.ApplySimplexRule(systeme)
	fmt.Printf("Solution trouvée : %v = %v \n", found, solution.ToString())
}



func TestSimplexeInt() {
	fmt.Println(" -------- TEST 29 -------- ")
	fmt.Println(" X <= 3 et X >= 2 ")
	x := types.MakerMeta("X", -1)
	trois := types.MakerConst(types.MakerId("3"), tInt)
	deux := types.MakerConst(types.MakerId("2"), tInt)
	p1 := types.MakePred(types.MakerId("lesseq"), []types.Term{x, trois}, typing.MkTypeArrow(typing.MkTypeCross(tInt, tInt), tProp))
	p2 := types.MakePred(types.MakerId("greateq"), []types.Term{x, deux}, typing.MkTypeArrow(typing.MkTypeCross(tInt, tInt), tProp))
	systeme := []types.Form{p1, p2}
	found, solution := ari.ApplySimplexRule(systeme)
	fmt.Printf("Solution trouvée : %v = %v \n", found, solution.ToString())
}

func TestSimplexeRatint() {
	fmt.Println(" -------- TEST 30 -------- ")
	fmt.Println(" X <= 3/2 et Y >= 2 - X rat, Y int")
	x := types.MakerMeta("X", -1)
	trois_demi := types.MakerConst(types.MakerId("3/2"), tRat)
	deux := types.MakerConst(types.MakerId("2"), tRat)
	p1 := types.MakePred(types.MakerId("lesseq"), []types.Term{x, trois_demi}, typing.MkTypeArrow(typing.MkTypeCross(tRat, tRat), tProp))
	p2 := types.MakePred(types.MakerId("greateq"), []types.Term{x, deux}, typing.MkTypeArrow(typing.MkTypeCross(tInt, tInt), tProp))
	systeme := []types.Form{p1, p2}
	found, solution := ari.ApplySimplexRule(systeme)
	fmt.Printf("Solution trouvée : %v = %v \n", found, solution.ToString())
}
