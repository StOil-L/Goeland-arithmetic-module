/***********************/
/* parser_operation.go */
/***********************/
/**
* This file provides function to parse a arithmetic term into an operation data structure, to be evaluate later.
**/

package ari

import (
	typing "ARI/polymorphism"
	"ARI/types"
	"fmt"
)

/* Transform a predicate into an operation struct */
func predToOperation(p types.Pred) operation {
	var arg1, arg2 types.Term
	var tScheme typing.TypeScheme
	// Assign var depending the number of args
	switch len(p.GetArgs()) {
	case 1:
		arg1 = p.GetArgs()[0]
	case 2:
		arg1 = p.GetArgs()[0]
		arg2 = p.GetArgs()[1]
	default:
		fmt.Printf("Illegal number of arguments in an artihmetic predicate")
		return nil
	}

	tScheme = p.GetTypeScheme()

	// Dans ce cas, on ne peut pas conclure, il faudra demander au simplexe
	if len(arg1.GetMetas()) > 0 || len(arg2.GetMetas()) > 0 {
		return nil
	}

	switch p.GetID().GetName() {
	case types.Id_eq.GetName():
		return createOperation(arg1, arg2, tScheme, EQ)
	case types.Id_neq.GetName():
		return createOperation(arg1, arg2, tScheme, NEQ)
	case "less":
		return createOperation(arg1, arg2, tScheme, LESS)
	case "lesseq":
		return createOperation(arg1, arg2, tScheme, LESSEQ)
	case "great":
		return createOperation(arg1, arg2, tScheme, GREAT)
	case "greateq":
		return createOperation(arg1, arg2, tScheme, GREATEQ)
	default:
		fmt.Println("Error parse")
		return nil
	}

}
