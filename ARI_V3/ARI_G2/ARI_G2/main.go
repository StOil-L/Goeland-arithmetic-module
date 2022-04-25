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

	fmt.Println(" ------------- TEST FONCTIONS BINAIRES ------------- ")

	// Fonctions binaires
	fmt.Println(" ------------- TEST SUM ------------- ")

	TestSumInt()
    TestSumNegInt()
    TestSumRat()
    TestSumRat2()
	TestSumNegRat()
    TestSumNegRat2()

	fmt.Println(" ------------- TEST DIFFERENCE ------------- ")

    TestDiffInt()
    TestDiffInt2()
    TestDiffNegInt()
    TestDiffNegInt2()
    TestDiffRat()
    TestDiffRat2()
    TestDiffNegRat()
    TestDiffNegRat()
    TestDiffNegRat2()
    TestDiffNegRat3()
    TestDiffRat()
    TestDiffRat2()
    TestDiffRat3()

	fmt.Println(" ------------- TEST Prod ------------- ")

	TestProdInt()
    TestProdNegInt()
	TestProdNegInt2()
    TestProdRat()
    TestProdNegRat()
    TestProdRat2()
    TestProdNegRat2()
	TestProdNegRat3()


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

	//Tests fonctions unaires 

	fmt.Println(" ------------- TEST UMINUS ------------- ")

	TestUminusInt()
	TestUminusNegInt()
	TestUminusRat()
	TestUminusNegRat() 
	TestUminusRat2()
	TestUminusNegRat2() 
	
	fmt.Println(" ------------- TEST FLOOR ------------- ")

	TestFloorInt()
	TestFloorNegInt()
	TestFloorRat()
	TestFloorNegRat() 	
	TestFloorRat2()
	TestFloorNegRat2() 	

	fmt.Println(" ------------- TEST CEILING ------------- ")
	
	TestCeilingInt()
	TestCeilingNegInt()
	TestCeilingRat()
	TestCeilingNegRat() 
	TestCeilingRat2()
	TestCeilingNegRat2() 

	fmt.Println(" ------------- TEST TRUNCATE ------------- ")
	
	TestTruncateInt()
	TestTruncateNegInt()
	TestTruncateRat()
	TestTruncateNegRat() 
	TestTruncateRat2()
	TestTruncateNegRat2()

	fmt.Println(" ------------- TEST ROUND ------------- ")
	 
	TestRoundInt()
	TestRoundNegInt()
	TestRoundRat()
	TestRoundNegRat() 
	TestRoundRat2()
	TestRoundNegRat2() 

	fmt.Println(" ------------- TEST Passe1 ------------- ")
/*	 
	// Tests règles simplexe
	TestSimplexeRat1()
	TestSimplexeRat2()
	TestSimplexeRat()
	TestSimplexeSumRat()
	TestSimplexeBeaucoupRat_2()
*/
//	TestSimplexePasse2()
//	TestSimplexePasse2_calcul_sum()
//	TestSimplexeSum()
	TestSimplexeSum_number()
	TestSimplexeProd_number()
	TestSimplexeSum_number_inv()
	TestSimplexeProd_number_inv()
	TestSimplexeSum_var()
	TestSimplexeBeaucoupRat()
	TestSimplexePasse2MultiEg()
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

func TestSumInt() {
	fmt.Println(" -------- TEST Sum Int -------- ")
	fmt.Println(" 1 + 2 = 3")
	un := types.MakerConst(types.MakerId("1"), tInt)
	deux := types.MakerConst(types.MakerId("2"), tInt)
	sum := types.MakeFun(types.MakerId("sum"), []types.Term{un, deux}, typing.GetTypeScheme("sum", typing.MkTypeCross(tInt, tInt)))
	solution,_:=ari.EvaluateFun(sum)
	fmt.Println("solution = ", solution) 
}

func TestSumNegInt() {
	fmt.Println(" -------- TEST Sum Neg Int -------- ")
	fmt.Println(" -1 + 2 = 1")
	moins_un := types.MakerConst(types.MakerId("-1"), tInt)
	deux := types.MakerConst(types.MakerId("2"), tInt)
	sum := types.MakeFun(types.MakerId("sum"), []types.Term{moins_un, deux}, typing.GetTypeScheme("sum", typing.MkTypeCross(tInt, tInt)))
	solution,_:=ari.EvaluateFun(sum)
	fmt.Println("solution = ", solution) 
}

func TestSumRat() {
	fmt.Println(" -------- TEST Sum Rat -------- ")
	fmt.Println(" 1.2 + 2.5 = 3.7")
	un_deux := types.MakerConst(types.MakerId("1.2"), tRat)
	deux_cinq := types.MakerConst(types.MakerId("2.5"), tRat)
	sum := types.MakeFun(types.MakerId("sum"), []types.Term{un_deux, deux_cinq}, typing.GetTypeScheme("sum", typing.MkTypeCross(tRat, tRat)))
	solution,_:=ari.EvaluateFun(sum)
	fmt.Println("solution = ", solution) 
}

func TestSumNegRat() {
	fmt.Println(" -------- TEST Sum Neg Rat -------- ")
	fmt.Println(" -1.2 + 2.5 = 1.3")
	moins_un_deux := types.MakerConst(types.MakerId("-1.2"), tRat)
	deux_cinq := types.MakerConst(types.MakerId("2.5"), tRat)
	sum := types.MakeFun(types.MakerId("sum"), []types.Term{moins_un_deux, deux_cinq}, typing.GetTypeScheme("sum", typing.MkTypeCross(tRat, tRat)))
	solution,_:=ari.EvaluateFun(sum)
	fmt.Println("solution = ", solution) 
}

func TestSumRat2() {
	fmt.Println(" -------- TEST Sum Rat 2 -------- ")
	fmt.Println(" 1/3 + 2.5 = 83333/100000?")
	un_sur_trois := types.MakerConst(types.MakerId("1/3"), tRat)
	deux_cinq := types.MakerConst(types.MakerId("2.5"), tRat)
	sum := types.MakeFun(types.MakerId("sum"), []types.Term{un_sur_trois, deux_cinq}, typing.GetTypeScheme("sum", typing.MkTypeCross(tRat, tRat)))
	solution,_:=ari.EvaluateFun(sum)
	fmt.Println("solution = ", solution) 
}

func TestSumNegRat2() {
	fmt.Println(" -------- TEST Sum Neg Rat 2 -------- ")
	fmt.Println(" -1/3 + 2.5 = 83333/100000?")
	moins_un_sur_trois := types.MakerConst(types.MakerId("-1/3"), tRat)
	deux_cinq := types.MakerConst(types.MakerId("2.5"), tRat)
	sum := types.MakeFun(types.MakerId("sum"), []types.Term{moins_un_sur_trois, deux_cinq}, typing.GetTypeScheme("sum", typing.MkTypeCross(tRat, tRat)))
	solution,_:=ari.EvaluateFun(sum)
	fmt.Println("solution = ", solution) 
}

func TestDiffInt() {
	fmt.Println(" -------- TEST Diff Int -------- ")
	fmt.Println(" 3 - 1 = 2")
	un := types.MakerConst(types.MakerId("1"), tInt)
	trois := types.MakerConst(types.MakerId("3"), tInt)
	difference := types.MakeFun(types.MakerId("difference"), []types.Term{trois, un}, typing.GetTypeScheme("difference", typing.MkTypeCross(tInt, tInt)))
	solution,_:=ari.EvaluateFun(difference)
	fmt.Println("solution = ", solution) 
}

func TestDiffInt2() {
	fmt.Println(" -------- TEST Diff Int 2-------- ")
	fmt.Println(" 1 - 2 = -1")
	un := types.MakerConst(types.MakerId("1"), tInt)
	deux := types.MakerConst(types.MakerId("2"), tInt)
	difference := types.MakeFun(types.MakerId("difference"), []types.Term{un, deux}, typing.GetTypeScheme("difference", typing.MkTypeCross(tInt, tInt)))
	solution,_:=ari.EvaluateFun(difference)
	fmt.Println("solution = ", solution) 
}

func TestDiffNegInt() {
	fmt.Println(" -------- TEST Diff Neg Int -------- ")
	fmt.Println(" -1 - -2 = 1")
	moins_un := types.MakerConst(types.MakerId("-1"), tInt)
	moins_deux := types.MakerConst(types.MakerId("-2"), tInt)
	difference := types.MakeFun(types.MakerId("difference"), []types.Term{moins_un, moins_deux}, typing.GetTypeScheme("difference", typing.MkTypeCross(tInt, tInt)))
	solution,_:=ari.EvaluateFun(difference)
	fmt.Println("solution = ", solution) 
}

func TestDiffNegInt2() {
	fmt.Println(" -------- TEST Diff Neg Int 2 -------- ")
	fmt.Println(" -3 - 2 = -5")
	moins_trois := types.MakerConst(types.MakerId("-3"), tInt)
	deux := types.MakerConst(types.MakerId("2"), tInt)
	difference := types.MakeFun(types.MakerId("difference"), []types.Term{moins_trois, deux}, typing.GetTypeScheme("difference", typing.MkTypeCross(tInt, tInt)))
	solution,_:=ari.EvaluateFun(difference)
	fmt.Println("solution = ", solution) 
}

func TestDiffRat() {
	fmt.Println(" -------- TEST Diff Rat -------- ")
	fmt.Println(" 1.2 - 2.5 = -1.3")
	un_deux := types.MakerConst(types.MakerId("6/5"), tRat)
	deux_cinq := types.MakerConst(types.MakerId("5/2"), tRat)
	difference := types.MakeFun(types.MakerId("difference"), []types.Term{un_deux, deux_cinq}, typing.GetTypeScheme("difference", typing.MkTypeCross(tRat, tRat)))
	solution,_:=ari.EvaluateFun(difference)
	fmt.Println("solution = ", solution) 
}

func TestDiffRat2() {
	fmt.Println(" -------- TEST Diff Rat 2 -------- ")
	fmt.Println("  2.5 - 1.2 = 1.3")
	un_deux := types.MakerConst(types.MakerId("1.2"), tRat)
	deux_cinq := types.MakerConst(types.MakerId("2.5"), tRat)
	difference := types.MakeFun(types.MakerId("difference"), []types.Term{deux_cinq, un_deux}, typing.GetTypeScheme("difference", typing.MkTypeCross(tRat, tRat)))
	solution,_:=ari.EvaluateFun(difference)
	fmt.Println("solution = ", solution) 
}

func TestDiffRat3() {
	fmt.Println(" -------- TEST Diff Rat 3 -------- ")
	fmt.Println(" 1/3 - 2.5 = -2*16666/100000?")
	un_sur_trois := types.MakerConst(types.MakerId("1/3"), tRat)
	deux_cinq := types.MakerConst(types.MakerId("2.5"), tRat)
	difference := types.MakeFun(types.MakerId("difference"), []types.Term{un_sur_trois, deux_cinq}, typing.GetTypeScheme("difference", typing.MkTypeCross(tRat, tRat)))
	solution,_:=ari.EvaluateFun(difference)
	fmt.Println("solution = ", solution) 
}

func TestDiffNegRat() {
	fmt.Println(" -------- TEST Diff Neg Rat -------- ")
	fmt.Println(" -1.2 - 2.5 = -3.7")
	moins_un_deux := types.MakerConst(types.MakerId("-1.2"), tRat)
	deux_cinq := types.MakerConst(types.MakerId("2.5"), tRat)
	difference := types.MakeFun(types.MakerId("difference"), []types.Term{moins_un_deux, deux_cinq}, typing.GetTypeScheme("difference", typing.MkTypeCross(tRat, tRat)))
	solution,_:=ari.EvaluateFun(difference)
	fmt.Println("solution = ", solution) 
}

func TestDiffNegRat2() {
	fmt.Println(" -------- TEST Diff Neg Rat 2 -------- ")
	fmt.Println(" -1.2 - -2.5 = 1.3")
	moins_un_deux := types.MakerConst(types.MakerId("-1.2"), tRat)
	moins_deux_cinq := types.MakerConst(types.MakerId("-2.5"), tRat)
	difference := types.MakeFun(types.MakerId("difference"), []types.Term{moins_un_deux, moins_deux_cinq}, typing.GetTypeScheme("difference", typing.MkTypeCross(tRat, tRat)))
	solution,_:=ari.EvaluateFun(difference)
	fmt.Println("solution = ", solution) 
}

func TestDiffNegRat3() {
	fmt.Println(" -------- TEST Diff Neg Rat 3 -------- ")
	fmt.Println(" -1/3 - -2.5 = 83333/100000?")
	moins_un_sur_trois := types.MakerConst(types.MakerId("-1/3"), tRat)
	moins_deux_cinq := types.MakerConst(types.MakerId("-2.5"), tRat)
	difference := types.MakeFun(types.MakerId("difference"), []types.Term{moins_un_sur_trois, moins_deux_cinq}, typing.GetTypeScheme("sum", typing.MkTypeCross(tRat, tRat)))
	solution,_:=ari.EvaluateFun(difference)
	fmt.Println("solution = ", solution) 
}

func TestProdInt() {
	fmt.Println(" -------- TEST Prod Int -------- ")
	fmt.Println(" 5 * 2 = 10")
	cinq := types.MakerConst(types.MakerId("5"), tInt)
	deux := types.MakerConst(types.MakerId("2"), tInt)
	product := types.MakeFun(types.MakerId("product"), []types.Term{cinq, deux}, typing.GetTypeScheme("product", typing.MkTypeCross(tInt, tInt)))
	solution,_:=ari.EvaluateFun(product)
	fmt.Println("solution = ", solution) 
}

func TestProdNegInt() {
	fmt.Println(" -------- TEST Prod Neg Int -------- ")
	fmt.Println(" -5 * 2 = -10")
	moins_cinq := types.MakerConst(types.MakerId("-5"), tInt)
	deux := types.MakerConst(types.MakerId("2"), tInt)
	product := types.MakeFun(types.MakerId("product"), []types.Term{moins_cinq, deux}, typing.GetTypeScheme("product", typing.MkTypeCross(tInt, tInt)))
	solution,_:=ari.EvaluateFun(product)
	fmt.Println("solution = ", solution) 
}

func TestProdNegInt2() {
	fmt.Println(" -------- TEST Prod Neg Int 2 -------- ")
	fmt.Println(" -5 * -2 = 10")
	moins_cinq := types.MakerConst(types.MakerId("-5"), tInt)
	moins_deux := types.MakerConst(types.MakerId("-2"), tInt)
	product := types.MakeFun(types.MakerId("product"), []types.Term{moins_cinq, moins_deux}, typing.GetTypeScheme("product", typing.MkTypeCross(tInt, tInt)))
	solution,_:=ari.EvaluateFun(product)
	fmt.Println("solution = ", solution) 
}

func TestProdRat() {
	fmt.Println(" -------- TEST Prod Rat -------- ")
	fmt.Println(" 1.2 * 2.5 = 3.0")
	un_deux := types.MakerConst(types.MakerId("1.2"), tRat)
	deux_cinq := types.MakerConst(types.MakerId("2.5"), tRat)
	product := types.MakeFun(types.MakerId("product"), []types.Term{un_deux, deux_cinq}, typing.GetTypeScheme("product", typing.MkTypeCross(tRat, tRat)))
	solution,_:=ari.EvaluateFun(product)
	fmt.Println("solution = ", solution) 
}

func TestProdNegRat() {
	fmt.Println(" -------- TEST Prod Neg Rat -------- ")
	fmt.Println(" -1.2 * 2.7 = -3.24")
	moins_un_deux := types.MakerConst(types.MakerId("-1.2"), tRat)
	deux_sept := types.MakerConst(types.MakerId("2.7"), tRat)
	product := types.MakeFun(types.MakerId("product"), []types.Term{moins_un_deux, deux_sept}, typing.GetTypeScheme("product", typing.MkTypeCross(tRat, tRat)))
	solution,_:=ari.EvaluateFun(product)
	fmt.Println("solution = ", solution) 
}

func TestProdRat2() {
	fmt.Println(" -------- TEST Prod Rat 2 -------- ")
	fmt.Println(" 1/3 * 2.5 = 3333333/25000000?")
	un_sur_trois := types.MakerConst(types.MakerId("1/3"), tRat)
	deux_cinq := types.MakerConst(types.MakerId("2.5"), tRat)
	product := types.MakeFun(types.MakerId("product"), []types.Term{un_sur_trois, deux_cinq}, typing.GetTypeScheme("product", typing.MkTypeCross(tRat, tRat)))
	solution,_:=ari.EvaluateFun(product)
	fmt.Println("solution = ", solution) 
}

func TestProdNegRat2() {
	fmt.Println(" -------- TEST Prod Neg Rat 2 -------- ")
	fmt.Println(" -1.2 * -2.7 = 3.24")
	moins_un_deux := types.MakerConst(types.MakerId("-1.2"), tRat)
	moins_deux_sept := types.MakerConst(types.MakerId("-2.7"), tRat)
	product := types.MakeFun(types.MakerId("product"), []types.Term{moins_un_deux, moins_deux_sept}, typing.GetTypeScheme("product", typing.MkTypeCross(tRat, tRat)))
	solution,_:=ari.EvaluateFun(product)
	fmt.Println("solution = ", solution) 
}

func TestProdNegRat3() {
	fmt.Println(" -------- TEST Sum Prod Rat 3 -------- ")
	fmt.Println(" -1/3 + 2.5 = -3333333/25000000?")
	moins_un_sur_trois := types.MakerConst(types.MakerId("-1/3"), tRat)
	deux_cinq := types.MakerConst(types.MakerId("2.5"), tRat)
	product := types.MakeFun(types.MakerId("product"), []types.Term{moins_un_sur_trois, deux_cinq}, typing.GetTypeScheme("product", typing.MkTypeCross(tRat, tRat)))
	solution,_:=ari.EvaluateFun(product)
	fmt.Println("solution = ", solution) 
}

func TestUminusInt() {
	fmt.Println(" -------- TEST Uminus Int -------- ")
	fmt.Println(" 4  devient -4 ")
	quatre := types.MakerConst(types.MakerId("4"),tInt) 
	//MakerFun c'est pour créer la fonction type uminus ^^
	uminus := types.MakerFun(types.MakerId("uminus"),[]types.Term{quatre}, tInt)
	solution,_:=ari.EvaluateFun(uminus)
	fmt.Println("solution = ", solution) 
	//MkTypeCross c'est le produit cartésien, en gros c'est pour les opérations binaire, je pense que nous il nous suffit de mettre MkTypeArrow dans le prédicat.
	//arrow c'est juste pour la flèche : genre   f(x) -> 2x  ^^
}

func TestUminusNegInt() {
	fmt.Println(" -------- TEST Uminus Neg Int -------- ")
	fmt.Println(" -4  devient 4 ")
	moins_quatre := types.MakerConst(types.MakerId("-4"),tInt) 
	uminus := types.MakerFun(types.MakerId("uminus"),[]types.Term{moins_quatre}, tInt)
	solution,_:=ari.EvaluateFun(uminus)
	fmt.Println("solution = ", solution) 
}

func TestUminusRat() {
	fmt.Println(" -------- TEST Uminus Rat -------- ")
	fmt.Println(" 4.5  devient -9/2 ")
	quatre_cinq := types.MakerConst(types.MakerId("9/2"),tRat) 
	uminus := types.MakerFun(types.MakerId("uminus"),[]types.Term{quatre_cinq}, tRat)
	solution,_:=ari.EvaluateFun(uminus)
	fmt.Println("solution = ", solution) 
}
func TestUminusNegRat() {
	fmt.Println(" -------- TEST Uminus Neg Rat -------- ")
	fmt.Println(" -4.5  devient 9/2 ")
	moins_quatre_cinq := types.MakerConst(types.MakerId("-9/2"),tRat) 
	uminus := types.MakerFun(types.MakerId("uminus"),[]types.Term{moins_quatre_cinq}, tRat)
	solution,_:=ari.EvaluateFun(uminus)
	fmt.Println("solution = ", solution) 
}

func TestUminusRat2() {
	fmt.Println(" -------- TEST Uminus Rat 2 -------- ")
	fmt.Println(" 1/3  devient -1/3 ")
	un_trois := types.MakerConst(types.MakerId("1/3"),tRat) 
	uminus := types.MakerFun(types.MakerId("uminus"),[]types.Term{un_trois}, tRat)
	solution,_:=ari.EvaluateFun(uminus)
	fmt.Println("solution = ", solution) 
}

func TestUminusNegRat2() {
	fmt.Println(" -------- TEST Uminus Neg Rat 2 -------- ")
	fmt.Println(" -1/3  devient 1/3 ")
	moins_un_trois := types.MakerConst(types.MakerId("-1/3"),tRat) 
	uminus := types.MakerFun(types.MakerId("uminus"),[]types.Term{moins_un_trois}, tRat)
	solution,_:=ari.EvaluateFun(uminus)
	fmt.Println("solution = ", solution) 
}


func TestFloorInt() {
	fmt.Println(" -------- TEST Floor Int -------- ")
	fmt.Println(" 4  devient 4 ")
	quatre := types.MakerConst(types.MakerId("4"),tInt)
	floor := types.MakerFun(types.MakerId("floor"),[]types.Term{quatre}, tInt)
	solution,_:=ari.EvaluateFun(floor)
	fmt.Println("solution = ", solution) 
}

func TestFloorNegInt() {
	fmt.Println(" -------- TEST Floor Neg Int -------- ")
	fmt.Println(" -4  devient -4 ")
	moins_quatre := types.MakerConst(types.MakerId("-4"),tInt) 
	floor := types.MakerFun(types.MakerId("floor"),[]types.Term{moins_quatre}, tInt)
	solution,_:=ari.EvaluateFun(floor)
	fmt.Println("solution = ", solution) 
}

func TestFloorRat() {
	fmt.Println(" -------- TEST Floor Rat -------- ")
	fmt.Println(" 4.7  devient 4.0 ")
	quatre_sept := types.MakerConst(types.MakerId("47/10"),tRat) 
	floor := types.MakerFun(types.MakerId("floor"),[]types.Term{quatre_sept}, tRat)
	solution,_:=ari.EvaluateFun(floor)
	fmt.Println("solution = ", solution) 
}

func TestFloorNegRat() {
	fmt.Println(" -------- TEST Floor Neg Rat -------- ")
	fmt.Println(" -4.2  devient -5 ")
	moins_quatre_deux := types.MakerConst(types.MakerId("-21/5"),tRat) 
	floor := types.MakerFun(types.MakerId("floor"),[]types.Term{moins_quatre_deux}, tRat)
	solution,_:=ari.EvaluateFun(floor)
	fmt.Println("solution = ", solution) 
}

func TestFloorRat2() {
	fmt.Println(" -------- TEST Floor Rat 2 -------- ")
	fmt.Println(" 1/3  devient 0.0 ")
	un_trois := types.MakerConst(types.MakerId("1/3"),tRat) 
	floor := types.MakerFun(types.MakerId("floor"),[]types.Term{un_trois}, tRat)
	solution,_:=ari.EvaluateFun(floor)
	fmt.Println("solution = ", solution) 
}

func TestFloorNegRat2() {
	fmt.Println(" -------- TEST Floor Neg Rat 2 -------- ")
	fmt.Println(" -1/3  devient -1.0 ")
	moins_un_trois := types.MakerConst(types.MakerId("-1/3"),tRat) 
	floor := types.MakerFun(types.MakerId("floor"),[]types.Term{moins_un_trois}, tRat)
	solution,_:=ari.EvaluateFun(floor)
	fmt.Println("solution = ", solution) 
}

func TestCeilingInt() {
	fmt.Println(" -------- TEST Ceiling Int -------- ")
	fmt.Println(" 4  devient 4 ")
	quatre := types.MakerConst(types.MakerId("4"),tInt)
	ceiling := types.MakerFun(types.MakerId("ceiling"),[]types.Term{quatre}, tInt)
	solution,_:=ari.EvaluateFun(ceiling)
	fmt.Println("solution = ", solution) 
}

func TestCeilingNegInt() {
	fmt.Println(" -------- TEST Ceiling Neg Int -------- ")
	fmt.Println(" -4  devient -4 ")
	moins_quatre := types.MakerConst(types.MakerId("-4"),tInt) 
	ceiling := types.MakerFun(types.MakerId("ceiling"),[]types.Term{moins_quatre}, tInt)
	solution,_:=ari.EvaluateFun(ceiling)
	fmt.Println("solution = ", solution) 
}

func TestCeilingRat() {
	fmt.Println(" -------- TEST Ceiling Rat -------- ")
	fmt.Println(" 4.7  devient 5.0 ")
	quatre_sept := types.MakerConst(types.MakerId("47/10"),tRat) 
	ceiling := types.MakerFun(types.MakerId("ceiling"),[]types.Term{quatre_sept}, tRat)
	solution,_:=ari.EvaluateFun(ceiling)
	fmt.Println("solution = ", solution) 
}
func TestCeilingNegRat() {
	fmt.Println(" -------- TEST Ceiling Neg Rat -------- ")
	fmt.Println(" -4.2  devient -4 ")
	moins_quatre_deux := types.MakerConst(types.MakerId("-21/5"),tRat) 
	ceiling := types.MakerFun(types.MakerId("ceiling"),[]types.Term{moins_quatre_deux}, tRat)
	solution,_:=ari.EvaluateFun(ceiling)
	fmt.Println("solution = ", solution) 
}

func TestCeilingRat2() {
	fmt.Println(" -------- TEST Ceiling Rat 2 -------- ")
	fmt.Println(" 1/3  devient 1.0 ")
	un_trois := types.MakerConst(types.MakerId("1/3"),tRat) 
	ceiling := types.MakerFun(types.MakerId("ceiling"),[]types.Term{un_trois}, tRat)
	solution,_:=ari.EvaluateFun(ceiling)
	fmt.Println("solution = ", solution) 
}

func TestCeilingNegRat2() {
	fmt.Println(" -------- TEST Ceiling Neg Rat 2 -------- ")
	fmt.Println(" -1/3  devient -0.0 ")
	moins_un_trois := types.MakerConst(types.MakerId("-1/3"),tRat) 
	ceiling := types.MakerFun(types.MakerId("ceiling"),[]types.Term{moins_un_trois}, tRat)
	solution,_:=ari.EvaluateFun(ceiling)
	fmt.Println("solution = ", solution) 
}


func TestTruncateInt() {
	fmt.Println(" -------- TEST Truncate Int -------- ")
	fmt.Println(" 4  devient 4 ")
	quatre := types.MakerConst(types.MakerId("4"),tInt)
	truncate := types.MakerFun(types.MakerId("truncate"),[]types.Term{quatre}, tInt)
	solution,_:=ari.EvaluateFun(truncate)
	fmt.Println("solution = ", solution) 
}

func TestTruncateNegInt() {
	fmt.Println(" -------- TEST Truncate Neg Int -------- ")
	fmt.Println(" -4  devient -4 ")
	moins_quatre := types.MakerConst(types.MakerId("-4"),tInt) 
	truncate := types.MakerFun(types.MakerId("truncate"),[]types.Term{moins_quatre}, tInt)
	solution,_:=ari.EvaluateFun(truncate)
	fmt.Println("solution = ", solution) 
}

func TestTruncateRat() {
	fmt.Println(" -------- TEST Truncate Rat -------- ")
	fmt.Println(" 4.7  devient 4.0 ")
	quatre_sept := types.MakerConst(types.MakerId("47/10"),tRat) 
	truncate := types.MakerFun(types.MakerId("truncate"),[]types.Term{quatre_sept}, tRat)
	solution,_:=ari.EvaluateFun(truncate)
	fmt.Println("solution = ", solution) 
}

func TestTruncateNegRat() {
	fmt.Println(" -------- TEST Truncate Neg Rat -------- ")
	fmt.Println(" -4.2  devient -4.0 ")
	moins_quatre_deux := types.MakerConst(types.MakerId("-21/5"),tRat) 
	truncate := types.MakerFun(types.MakerId("truncate"),[]types.Term{moins_quatre_deux}, tRat)
	solution,_:=ari.EvaluateFun(truncate)
	fmt.Println("solution = ", solution) 
}

func TestTruncateRat2() {
	fmt.Println(" -------- TEST Truncate Rat 2 -------- ")
	fmt.Println(" 1/3  devient 0.0 ")
	un_trois := types.MakerConst(types.MakerId("1/3"),tRat) 
	truncate := types.MakerFun(types.MakerId("truncate"),[]types.Term{un_trois}, tRat)
	solution,_:=ari.EvaluateFun(truncate)
	fmt.Println("solution = ", solution) 
}

func TestTruncateNegRat2() {
	fmt.Println(" -------- TEST Truncate Neg Rat 2 -------- ")
	fmt.Println(" -1/3  devient -0.0 ")
	moins_un_trois := types.MakerConst(types.MakerId("-1/3"),tRat) 
	truncate := types.MakerFun(types.MakerId("truncate"),[]types.Term{moins_un_trois}, tRat)
	solution,_:=ari.EvaluateFun(truncate)
	fmt.Println("solution = ", solution) 
}


func TestRoundInt() {
	fmt.Println(" -------- TEST Round Int -------- ")
	fmt.Println(" 4  devient 4 ")
	quatre := types.MakerConst(types.MakerId("4"),tInt)
	round := types.MakerFun(types.MakerId("round"),[]types.Term{quatre}, tInt)
	solution,_:=ari.EvaluateFun(round)
	fmt.Println("solution = ", solution) 
}

func TestRoundNegInt() {
	fmt.Println(" -------- TEST Round Neg Int -------- ")
	fmt.Println(" -4  devient -4 ")
	moins_quatre := types.MakerConst(types.MakerId("-4"),tInt) 
	round := types.MakerFun(types.MakerId("round"),[]types.Term{moins_quatre}, tInt)
	solution,_:=ari.EvaluateFun(round)
	fmt.Println("solution = ", solution) 
}

func TestRoundRat() {
	fmt.Println(" -------- TEST Round Rat -------- ")
	fmt.Println(" 4.7  devient 5.0 ")
	quatre_sept := types.MakerConst(types.MakerId("47/10"),tRat) 
	round := types.MakerFun(types.MakerId("round"),[]types.Term{quatre_sept}, tRat)
	solution,_:=ari.EvaluateFun(round)
	fmt.Println("solution = ", solution) 
}

func TestRoundNegRat() {
	fmt.Println(" -------- TEST Round Neg Rat -------- ")
	fmt.Println(" -4.2  devient -4.0 ")
	moins_quatre_deux := types.MakerConst(types.MakerId("-21/5"),tRat) 
	round := types.MakerFun(types.MakerId("round"),[]types.Term{moins_quatre_deux}, tRat)
	solution,_:=ari.EvaluateFun(round)
	fmt.Println("solution = ", solution) 
}

func TestRoundRat2() {
	fmt.Println(" -------- TEST Round Rat 2 -------- ")
	fmt.Println(" 1/3  devient 0.0 ")
	un_trois := types.MakerConst(types.MakerId("1/3"),tRat) 
	round := types.MakerFun(types.MakerId("round"),[]types.Term{un_trois}, tRat)
	solution,_:=ari.EvaluateFun(round)
	fmt.Println("solution = ", solution) 
}

func TestRoundNegRat2() {
	fmt.Println(" -------- TEST Round Neg Rat 2 -------- ")
	fmt.Println(" -5/3  devient -2.0 ")
	moins_cinq_trois := types.MakerConst(types.MakerId("-5/3"),tRat) 
	round := types.MakerFun(types.MakerId("round"),[]types.Term{moins_cinq_trois}, tRat)
	solution,_:=ari.EvaluateFun(round)
	fmt.Println("solution = ", solution) 
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
	fmt.Println(" X + Y = 3/2")
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
	fmt.Println(" -------- TEST 28.998 -------- ")
	fmt.Println(" (((((X + Y)-Z)+K)+Y)-Z)+K) = 3/2")
	x := types.MakerMeta("X_0__", -1)
	y := types.MakerMeta("Y", -1)
	z := types.MakerMeta("Z", -1)
	k := types.MakerMeta("K", -1)
	trois_demi := types.MakerConst(types.MakerId("3/2"), tRat)
	sum := types.MakeFun(types.MakerId("sum"), []types.Term{x, y}, typing.GetTypeScheme("sum", typing.MkTypeCross(tRat, tRat)))
	diff := types.MakeFun(types.MakerId("difference"), []types.Term{sum, z}, typing.GetTypeScheme("difference", typing.MkTypeCross(tRat, tRat)))
	prod := types.MakeFun(types.MakerId("sum"), []types.Term{diff, k}, typing.GetTypeScheme("product", typing.MkTypeCross(tRat, tRat)))
	sum2 := types.MakeFun(types.MakerId("sum"), []types.Term{prod, y}, typing.GetTypeScheme("sum", typing.MkTypeCross(tRat, tRat)))
	diff2 := types.MakeFun(types.MakerId("difference"), []types.Term{sum2, z}, typing.GetTypeScheme("difference", typing.MkTypeCross(tRat, tRat)))
	prod2 := types.MakeFun(types.MakerId("sum"), []types.Term{diff2, k}, typing.GetTypeScheme("product", typing.MkTypeCross(tRat, tRat)))
	p := types.MakePred(types.Id_eq, []types.Term{prod2, trois_demi}, typing.MkTypeArrow(typing.MkTypeCross(tRat, tRat), tProp))
	systeme := []types.Form{p}
	found, solution := ari.ApplySimplexRule(systeme)
	fmt.Printf("Solution trouvée : %v = %v \n", found, solution.ToString())
}

func TestSimplexeBeaucoupRat_2() {
	fmt.Println(" -------- TEST 28.999 -------- ")
	fmt.Println(" 1) (((((X_0__ + Y)-Z)*K)+Y)-Z)*K) = (A+B)*(C+K)  2)  X_0__ + Y = 3/2 + M       ")
	x := types.MakerMeta("X_0__", -1)
	y := types.MakerMeta("Y", -1)
	z := types.MakerMeta("Z", -1)
	k := types.MakerMeta("K", -1)
	a := types.MakerMeta("A", -1)
	b := types.MakerMeta("B", -1)
	c := types.MakerMeta("C", -1)
	m := types.MakerMeta("M", -1)
	//t1
	sum_x_y := types.MakeFun(types.MakerId("sum"), []types.Term{x, y}, typing.GetTypeScheme("sum", typing.MkTypeCross(tRat, tRat)))
	diff_z := types.MakeFun(types.MakerId("difference"), []types.Term{sum_x_y, z}, typing.GetTypeScheme("difference", typing.MkTypeCross(tRat, tRat)))
	prod_k := types.MakeFun(types.MakerId("product"), []types.Term{diff_z, k}, typing.GetTypeScheme("product", typing.MkTypeCross(tRat, tRat)))
	sum_y := types.MakeFun(types.MakerId("sum"), []types.Term{prod_k, y}, typing.GetTypeScheme("sum", typing.MkTypeCross(tRat, tRat)))
	diff2_z := types.MakeFun(types.MakerId("difference"), []types.Term{sum_y, z}, typing.GetTypeScheme("difference", typing.MkTypeCross(tRat, tRat)))
	prod2 := types.MakeFun(types.MakerId("product"), []types.Term{diff2_z, k}, typing.GetTypeScheme("product", typing.MkTypeCross(tRat, tRat)))
	//t2
	sum_a_b := types.MakeFun(types.MakerId("sum"), []types.Term{a, b}, typing.GetTypeScheme("sum", typing.MkTypeCross(tRat, tRat)))
	sum_c_k := types.MakeFun(types.MakerId("sum"), []types.Term{c, k}, typing.GetTypeScheme("sum", typing.MkTypeCross(tRat, tRat)))
	prod := types.MakeFun(types.MakerId("product"), []types.Term{sum_a_b, sum_c_k}, typing.GetTypeScheme("product", typing.MkTypeCross(tRat, tRat)))

	//t1 déjà fait
	
	//t2
	trois_demi := types.MakerConst(types.MakerId("3/2"), tRat)
	sum_m_trois_demi := types.MakeFun(types.MakerId("sum"), []types.Term{ trois_demi,m}, typing.GetTypeScheme("sum", typing.MkTypeCross(tRat, tRat)))
	p1 := types.MakePred(types.Id_eq, []types.Term{prod2, prod}, typing.MkTypeArrow(typing.MkTypeCross(tRat, tRat), tProp))
	p2 := types.MakePred(types.Id_eq, []types.Term{sum_x_y, sum_m_trois_demi}, typing.MkTypeArrow(typing.MkTypeCross(tRat, tRat), tProp))
	systeme := []types.Form{p1,p2}
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
func TestSimplexePasse2(){

fmt.Println(" -------- TEST Passe2 -------- ")
	fmt.Println(" X+Y = 3/2  ")
	x := types.MakerMeta("X", -1)
	y := types.MakerMeta("Y", -1)
	trois_demi := types.MakerConst(types.MakerId("3/2"), tRat)
	sum := types.MakerFun(types.MakerId("sum"),[]types.Term{x, y}, typing.GetTypeScheme("sum", typing.MkTypeCross(tRat, tRat)))
	p := types.MakePred(types.Id_eq, []types.Term{sum, trois_demi}, typing.MkTypeArrow(typing.MkTypeCross(tRat, tRat), tProp))
	systeme := []types.Form{p}
	found, solution := ari.ApplySimplexRule(systeme)
	fmt.Printf("Solution trouvée : %v = %v \n", found, solution.ToString())

}

func TestSimplexePasse2_calcul_sum(){

	fmt.Println(" -------- TEST Passe2_2 -------- ")
		fmt.Println(" X(3/2 + 3/2) = 3/2  ")
		x := types.MakerMeta("X", -1)
		trois_demi := types.MakerConst(types.MakerId("3/2"), tRat)
		sum := types.MakerFun(types.MakerId("sum"),[]types.Term{trois_demi, trois_demi}, typing.GetTypeScheme("sum", typing.MkTypeCross(tRat, tRat)))
		prod := types.MakerFun(types.MakerId("product"),[]types.Term{x, sum}, typing.GetTypeScheme("product", typing.MkTypeCross(tRat, tRat)))
		p := types.MakePred(types.Id_eq, []types.Term{prod, trois_demi}, typing.MkTypeArrow(typing.MkTypeCross(tRat, tRat), tProp))
		systeme := []types.Form{p}
		found, solution := ari.ApplySimplexRule(systeme)
		fmt.Printf("Solution trouvée : %v = %v \n", found, solution.ToString())
	
}

func TestSimplexeSum(){

	fmt.Println(" -------- TEST Passe2_2 somme naze-------- ")
	fmt.Println(" X + 3/2 = 3/2  ")
	x := types.MakerMeta("X", -1)
	trois_demi := types.MakerConst(types.MakerId("3/2"), tRat)
	sum := types.MakerFun(types.MakerId("sum"),[]types.Term{x, trois_demi}, typing.GetTypeScheme("sum", typing.MkTypeCross(tRat, tRat)))
	p := types.MakePred(types.Id_eq, []types.Term{sum, trois_demi}, typing.MkTypeArrow(typing.MkTypeCross(tRat, tRat), tProp))
	systeme := []types.Form{p}
	found, solution := ari.ApplySimplexRule(systeme)
	fmt.Printf("Solution trouvée : %v = %v \n", found, solution.ToString())

}

func TestSimplexeSum_number(){

	fmt.Println(" -------- TEST Passe2 sum number-------- ")
	fmt.Println(" (3/2 + 3/2) * x = 3/2  ")
	x := types.MakerMeta("X", -1)
	trois_demi := types.MakerConst(types.MakerId("3/2"), tRat)
	sum := types.MakerFun(types.MakerId("sum"),[]types.Term{trois_demi, trois_demi}, typing.GetTypeScheme("sum", typing.MkTypeCross(tRat, tRat)))
	prod := types.MakerFun(types.MakerId("product"), []types.Term{sum,x}, typing.GetTypeScheme("product",typing.MkTypeCross(tRat, tRat)))
	p := types.MakePred(types.Id_eq, []types.Term{prod, trois_demi}, typing.MkTypeArrow(typing.MkTypeCross(tRat, tRat), tProp))
	systeme := []types.Form{p}
	found, solution := ari.ApplySimplexRule(systeme)
	fmt.Printf("Solution trouvée : %v = %v \n", found, solution.ToString())

}


func TestSimplexeProd_number(){

	fmt.Println(" -------- TEST Passe2 prod number-------- ")
	fmt.Println(" (3/2 * 3/2) * x = 3/2  ")
	x := types.MakerMeta("X", -1)
	trois_demi := types.MakerConst(types.MakerId("3/2"), tRat)
	prod2 := types.MakerFun(types.MakerId("product"),[]types.Term{trois_demi, trois_demi}, typing.GetTypeScheme("product", typing.MkTypeCross(tRat, tRat)))
	prod := types.MakerFun(types.MakerId("product"), []types.Term{prod2,x}, typing.GetTypeScheme("product",typing.MkTypeCross(tRat, tRat)))
	p := types.MakePred(types.Id_eq, []types.Term{prod, trois_demi}, typing.MkTypeArrow(typing.MkTypeCross(tRat, tRat), tProp))
	systeme := []types.Form{p}
	found, solution := ari.ApplySimplexRule(systeme)
	fmt.Printf("Solution trouvée : %v = %v \n", found, solution.ToString())

}



func TestSimplexeSum_number_inv(){

	fmt.Println(" -------- TEST Passe2 sum number inv-------- ")
	fmt.Println("  3/2 =(3/2 + 3/2) * x  ")
	x := types.MakerMeta("X", -1)
	trois_demi := types.MakerConst(types.MakerId("3/2"), tRat)
	sum := types.MakerFun(types.MakerId("sum"),[]types.Term{trois_demi, trois_demi}, typing.GetTypeScheme("sum", typing.MkTypeCross(tRat, tRat)))
	prod := types.MakerFun(types.MakerId("product"), []types.Term{sum,x}, typing.GetTypeScheme("product",typing.MkTypeCross(tRat, tRat)))
	p := types.MakePred(types.Id_eq, []types.Term{trois_demi,prod}, typing.MkTypeArrow(typing.MkTypeCross(tRat, tRat), tProp))
	systeme := []types.Form{p}
	found, solution := ari.ApplySimplexRule(systeme)
	fmt.Printf("Solution trouvée : %v = %v \n", found, solution.ToString())

}


func TestSimplexeProd_number_inv(){

	fmt.Println(" -------- TEST Passe2 prod number inv-------- ")
	fmt.Println("  3/2 = (3/2 * 3/2) * x   ")
	x := types.MakerMeta("X", -1)
	trois_demi := types.MakerConst(types.MakerId("3/2"), tRat)
	prod2 := types.MakerFun(types.MakerId("product"),[]types.Term{trois_demi, trois_demi}, typing.GetTypeScheme("product", typing.MkTypeCross(tRat, tRat)))
	prod := types.MakerFun(types.MakerId("product"), []types.Term{prod2,x}, typing.GetTypeScheme("product",typing.MkTypeCross(tRat, tRat)))
	p := types.MakePred(types.Id_eq, []types.Term{ trois_demi, prod}, typing.MkTypeArrow(typing.MkTypeCross(tRat, tRat), tProp))
	systeme := []types.Form{p}
	found, solution := ari.ApplySimplexRule(systeme)
	fmt.Printf("Solution trouvée : %v = %v \n", found, solution.ToString())

}


func TestSimplexeSum_var(){

	fmt.Println(" -------- TEST Passe2 sum number inv-------- ")
	fmt.Println("  X + Y = 3/2 ")
	x := types.MakerMeta("X", -1)
	y := types.MakerMeta("Y", -1)
	trois_demi := types.MakerConst(types.MakerId("3/2"), tRat)
	sum := types.MakerFun(types.MakerId("sum"),[]types.Term{x, y}, typing.GetTypeScheme("sum", typing.MkTypeCross(tRat, tRat)))
	p := types.MakePred(types.Id_eq, []types.Term{sum,trois_demi}, typing.MkTypeArrow(typing.MkTypeCross(tRat, tRat), tProp))
	systeme := []types.Form{p}
	found, solution := ari.ApplySimplexRule(systeme)
	fmt.Printf("Solution trouvée : %v = %v \n", found, solution.ToString())

}


func TestSimplexePasse2MultiEg() {
	fmt.Println(" -------- TEST passe2 multiEg -------- ")
	fmt.Println(" X = 3/2 et Y = 2/3  et  3/2 = Z et X = 2/3 ")
	x := types.MakerMeta("X", -1)
	y := types.MakerMeta("Y", -1)
	z := types.MakerMeta("Z", -1)
	trois_demi := types.MakerConst(types.MakerId("3/2"), tRat)
	deux_tiers := types.MakerConst(types.MakerId("2/3"), tRat)
	p1 := types.MakePred(types.Id_eq, []types.Term{x, trois_demi}, typing.MkTypeArrow(typing.MkTypeCross(tRat, tRat), tProp))
	p2 := types.MakePred(types.Id_eq, []types.Term{y, deux_tiers}, typing.MkTypeArrow(typing.MkTypeCross(tRat, tRat), tProp))
	p3 := types.MakePred(types.Id_eq, []types.Term{ trois_demi, z}, typing.MkTypeArrow(typing.MkTypeCross(tRat, tRat), tProp))
	p4 := types.MakePred(types.Id_eq, []types.Term{x, deux_tiers}, typing.MkTypeArrow(typing.MkTypeCross(tRat, tRat), tProp))
	systeme := []types.Form{p1, p2, p3, p4}
	found, solution := ari.ApplySimplexRule(systeme)
	fmt.Printf("Solution trouvée : %v = %v \n", found, solution.ToString())
}