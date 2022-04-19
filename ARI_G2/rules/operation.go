/****************/
/* operation.go */
/****************/
/**
* This file provides rules to manage the opration data-structure.
**/

package ari

import (
	typing "ARI/polymorphism"
	"ARI/types"
	"fmt"
	"math/big"
	"strconv"
)

const (
	EQ int = iota
	NEQ
	LESS
	LESSEQ
	GREAT
	GREATEQ
)

var tInt typing.TypeHint
var tRat typing.TypeHint
var tProp typing.TypeHint
var tIntIntProp typing.TypeScheme
var tRatRatProp typing.TypeScheme
var tIntProp typing.TypeScheme
var tRatProp typing.TypeScheme

type operation interface {
	getOperator() int
	toString() string
}

type operationInt struct {
	arg1, arg2, operator int
}

type operationRat struct {
	arg1, arg2 *big.Rat
	operator   int
}

func opToString(i int) string {
	switch i {
	case EQ:
		return "="
	case NEQ:
		return "!="
	case LESS:
		return "<"
	case LESSEQ:
		return "<="
	case GREAT:
		return ">"
	case GREATEQ:
		return ">="
	default:
		return "Operator type unknown"
	}
}

func (oi operationInt) getArg1() int {
	return oi.arg1
}
func (oi operationInt) getArg2() int {
	return oi.arg2
}
func (oi operationInt) getOperator() int {
	return oi.operator
}
func (oi operationInt) toString() string {
	return strconv.Itoa(oi.getArg1()) + " " + opToString(oi.getOperator()) + " " + strconv.Itoa(oi.getArg2())
}

func (or operationRat) getArg1() *big.Rat {
	return or.arg1
}
func (or operationRat) getArg2() *big.Rat {
	return or.arg2
}
func (or operationRat) getOperator() int {
	return or.operator
}
func (or operationRat) toString() string {
	return or.getArg1().RatString() + " " + opToString(or.getOperator()) + " " + or.getArg2().RatString()
}

func Init() {
	tInt = typing.MkTypeHint("int")
	tRat = typing.MkTypeHint("rat")
	tProp = typing.MkTypeHint("o")
	tIntIntProp = typing.MkTypeArrow(typing.MkTypeCross(tInt, tInt), tProp)
	tRatRatProp = typing.MkTypeArrow(typing.MkTypeCross(tRat, tRat), tProp)
	tIntProp = typing.MkTypeArrow(tInt, tProp)
	tRatProp = typing.MkTypeArrow(tRat, tProp)
}

func checkIntIntToProp(ts typing.TypeScheme) bool {
	return ts.Equals(tIntIntProp)
}

func checkRatRatToProp(ts typing.TypeScheme) bool {
	return ts.Equals(tRatRatProp)
}

func checkIntToProp(ts typing.TypeScheme) bool {
	return ts.Equals(tIntProp)
}

func checkRatToProp(ts typing.TypeScheme) bool {
	return ts.Equals(tRatProp)
}

func checkUnaryArithmeticPredicate(id types.Id) bool {
	s := id.ToString()
	return s == "is_int" || s == "is_rat"
}

func checkBinaryArithmeticPredicate(id types.Id) bool {
	s := id.ToString()
	return id.Equals(types.Id_eq) || id.Equals(types.Id_neq) || s == "less" || s == "lesseq" || s == "great" || s == "greateq"
}

func createOperation(arg1, arg2 types.Term, tScheme typing.TypeScheme, operator int) operation {
	switch {
	case checkIntIntToProp(tScheme):
		v1, err1 := termToInt(arg1)
		if err1 != nil {
			fmt.Printf("Error createOperation : term1 %v - %v is not an int : %v\n", arg1.ToString(), v1, err1)
			return nil
		}
		v2, err2 := termToInt(arg2)
		if err2 != nil {
			fmt.Printf("Error createOperation : term2 %v - %v is not an int : %v\n", arg2.ToString(), v2, err2)
			return nil
		}
		return operationInt{v1, v2, operator}
	case checkRatRatToProp(tScheme):
		v1, err1 := termToRat(arg1)
		if err1 != nil {
			fmt.Printf("Error createOperation : term1 %v - %v is not a rat : %v\n", arg1.ToString(), v1, err1)
			return nil
		}
		v2, err2 := termToRat(arg2)
		if err2 != nil {
			fmt.Printf("Error createOperation : term2 %v - %v is not a rat : %v\n", arg2.ToString(), v2, err2)
			return nil
		}
		return operationRat{v1, v2, operator}
	default:
		fmt.Println("Error type unknown")
	}
	fmt.Printf("Error createOperation : %v %v %v\n", arg1.ToString(), operator, arg2.ToString())
	return nil
}
