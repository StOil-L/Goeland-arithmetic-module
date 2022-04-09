/****************/
/* operation.go */
/****************/
/**
* This file provides rules to manage the opration data-structure.
**/

package ari

import (
	typing "ARI/polymorphism"
	"math/big"
)

const (
	EQ int = iota
	NEQ
	LESS
	LESSEQ
	GREAT
	GREATEQ
)

var tInt = typing.MkTypeHint("int")
var tRat = typing.MkTypeHint("rat")
var tProp = typing.MkTypeHint("o")
var tIntIntProp = typing.MkTypeArrow(typing.MkTypeCross(tInt, tInt), tProp)
var tRatRatProp = typing.MkTypeArrow(typing.MkTypeCross(tRat, tRat), tProp)

type operation interface {
	getOperator() int
}

type operationInt struct {
	arg1, arg2, operator int
}

type operationRat struct {
	arg1, arg2 *big.Rat
	operator   int
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

func (or operationRat) getArg1() *big.Rat {
	return or.arg1
}
func (or operationRat) getArg2() *big.Rat {
	return or.arg2
}
func (or operationRat) getOperator() int {
	return or.operator
}

func checkIntIntToProp(ts typing.TypeScheme) bool {
	return ts.Equals(tIntIntProp)
}

func checkRatRatToProp(ts typing.TypeScheme) bool {
	return ts.Equals(tRatRatProp)
}
