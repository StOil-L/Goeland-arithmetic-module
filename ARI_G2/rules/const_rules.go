/******************/
/* const_rules.go */
/******************/
/**
* This file provides rules to manage const rules
**/

package ari

import (
	"ARI/types"
	"fmt"
)

/** This function returns true if the constraint fits with artithmetic **/
func constRule(o operation) bool {
	fmt.Printf("Operation : %v\n", o.toString())
	// En fonction du type de mon opération
	switch opType := o.(type) {

	// Cas entier
	case operationInt:
		switch opType.getOperator() {
		case EQ:
			return opType.getArg1() == opType.getArg2()
		case NEQ:
			return opType.getArg1() != opType.getArg2()
		case LESS:
			return opType.getArg1() < opType.getArg2()
		case LESSEQ:
			return opType.getArg1() <= opType.getArg2()
		case GREAT:
			return opType.getArg1() > opType.getArg2()
		case GREATEQ:
			return opType.getArg1() >= opType.getArg2()
		default:
			println("Error constRule - int : operator type unknown")
			return false
		}

	// Cas rationnel (rien a changer grace a Julie)
	case operationRat:
		switch opType.getOperator() {
		case EQ:
			return opType.getArg1() == opType.getArg2()
		case NEQ:
			return opType.getArg1() != opType.getArg2()
		case LESS:
			return opType.getArg1().Cmp(opType.getArg2()) == -1
		case LESSEQ:
			return opType.getArg1().Cmp(opType.getArg2()) <= 0
		case GREAT:
			return opType.getArg1().Cmp(opType.getArg2()) > 0
		case GREATEQ:
			return opType.getArg1().Cmp(opType.getArg2()) >= 0
		default:
			println("Error constRule - int : operator type unknown")
			return false
		}

	default:
		println("Error constRule : operation type unknown")
		return false
	}
}

/***
* Type checker  for predicate is_int and is_rat
* Return true if the predicat is is_int or is_rat, plus true if the type fits, false otherwise
**/
func applyTypeRule(p types.Pred) (bool, bool) {
	if p.GetArgs()[0].IsMeta() {
		return false, false
	}
	switch p.GetID().GetName() {
	case "is_int":
		return true, checkIntToProp(p.GetTypeScheme())
	case "is_rat":
		return true, checkRatToProp(p.GetTypeScheme())
	default:
		return false, false
	}
}
