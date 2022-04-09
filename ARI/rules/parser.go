/*************/
/* parser.go */
/*************/
/**
* This file provides function to parse a problem into arithmetic's data-structure.
**/

package ari

import (
	typing "ARI/polymorphism"
	"ARI/types"
	"errors"
	"fmt"
	"math/big"
	"strconv"
)

/*
* sum(1,1) -> 1 + 1
* Ici en string
* Mais à s'adpater à la structure arithmétique d'un problème :
* Struct {
* arg_1
* arg_2
* opérateur
* }
**/

// TODO : compléter les cas (greater, greatereq, ...)
func PredToConst(p types.Pred) operation {
	arg1 := p.GetArgs()[0]
	arg2 := p.GetArgs()[1]
	tScheme := p.GetTypeScheme()

	if arg1.IsMeta() || arg2.IsMeta() {
		return nil
	}

	switch p.GetID().GetName() {
	case types.Id_eq.GetName():
		createOperation(arg1, arg2, tScheme, EQ)
	case types.Id_neq.GetName():
		createOperation(arg1, arg2, tScheme, NEQ)
	case "less":
		createOperation(arg1, arg2, tScheme, LESS)
	case "lesseq":
		createOperation(arg1, arg2, tScheme, LESSEQ)
	case "great":
		createOperation(arg1, arg2, tScheme, GREAT)
	case "greateq":
		createOperation(arg1, arg2, tScheme, GREATEQ)
	}
	fmt.Println("Error parse")
	return nil
}

func createOperation(arg1, arg2 types.Term, tScheme typing.TypeScheme, operator int) operation {
	switch {
	case checkIntIntToProp(tScheme):
		v1, err1 := TermToInt(arg1)
		if err1 == nil {
			fmt.Println("Error createOperation")
			return nil
		}
		v2, err2 := TermToInt(arg2)
		if err2 == nil {
			fmt.Println("Error createOperation")
			return nil
		}
		return operationInt{v1, v2, operator}
	case checkRatRatToProp(tScheme):
		v1, err1 := TermToRat(arg1)
		if err1 == nil {
			fmt.Println("Error createOperation")
			return nil
		}
		v2, err2 := TermToRat(arg2)
		if err2 == nil {
			fmt.Println("Error createOperation")
			return nil
		}
		return operationRat{v1, v2, operator}
	default:
		fmt.Println("Error type unknown")
	}
	fmt.Println("Error createOperation")
	return nil
}

func TermToInt(t types.Term) (int, error) {
	switch ttypes := t.(type) {
	case types.Fun:
		// Constante
		if len(ttypes.GetArgs()) == 0 {
			strconv.Atoi(ttypes.GetID().GetName())
		} else {
			if res, err := funToInt(ttypes); err == nil {
				return res, nil
			} else {
				return 0, err
			}
		}
	default:
		fmt.Println("Error in conversion : format not found")
		return 0, errors.New("Error")
	}
	return 0, errors.New("Error")
}

// TODO : compléter avec les fonction (voir tptp_native)
func funToInt(f types.Fun) (int, error) {
	arg1 := f.GetArgs()[0]
	arg2 := f.GetArgs()[1]
	switch f.GetID().GetName() {
	case "sum":
		res1, err1 := TermToInt(arg1)
		if err1 != nil {
			return 0, err1
		}
		res2, err2 := TermToInt(arg2)
		if err2 != nil {
			return 0, err2
		}
		return res1 + res2, nil
	case "difference":
		res1, err1 := TermToInt(arg1)
		if err1 != nil {
			return 0, err1
		}
		res2, err2 := TermToInt(arg2)
		if err2 != nil {
			return 0, err2
		}
		return res1 - res2, nil
	}
	case "product":
		res1, err1 := TermToInt(arg1)
		if err1 != nil {
			return 0, err1
		}
		res2, err2 := TermToInt(arg2)
		if err2 != nil {
			return 0, err2
		}
		return res1 * res2, nil
	}
	// Que faire dans le cas d'un quotient réél : 1. Renvoi un réél 2. Revoi la division entière
	// cas 1
	case "quotient_t":
		res1, err1 := TermToInt(arg1)
		if err1 != nil {
			return 0, err1
		}
		res2, err2 := TermToInt(arg2)
		// Rajout du cas : res2 == 0
		if err2 != nil || res2 == 0 {
			return 0, err2
		}
		return res1 / res2, nil
	}
	// cas 2
	case "quotient_t":
		res1, err1 := TermToInt(arg1)
		if err1 != nil {
			return 0, err1
		}
		res2, err2 := TermToInt(arg2)
		// Rajout du cas : res2 == 0
		if err2 != nil || res2 == 0 {
			return 0, err2
		}
		return (res1 / res2) - (res1 / res2)%1, nil
	}
	case "remainder_t":
		res1, err1 := TermToInt(arg1)
		if err1 != nil {
			return 0, err1
		}
		res2, err2 := TermToInt(arg2)
		// Rajout du cas : res2 == 0
		if err2 != nil || res2 == 0 {
			return 0, err2
		}
		return res1 % res2, nil
	}
	return 0, nil
}

// TODO : tout
func TermToRat(t types.Term) (*big.Rat, error) {
	switch ttypes := t.(type) {
	case types.Fun:
		// Constante
		if len(ttypes.GetArgs()) == 0 {
			strconv.Atoi(ttypes.GetID().GetName())
		} else {
			if res, err := funToRat(ttypes); err == nil {
				return res, nil
			} else {
				return new(big.Rat), err
			}
		}
	default:
		fmt.Println("Error in conversion Rat: format not found")
		return new(big.Rat), errors.New("Error Rat")
	}

	return new(big.Rat), errors.New("Error Rat")
}

func funToRat(f types.Fun) (*big.Rat, error) {
	arg1 := f.GetArgs()[0]
	arg2 := f.GetArgs()[1]
	switch f.GetID().GetName() {
	case "sum":
		res1, err1 := TermToRat(arg1)
		if err1 != nil {
			return new(big.Rat), err1
		}
		res2, err2 := TermToRat(arg2)
		if err2 != nil {
			return new(big.Rat), err2
		}
		// On créé un nouveau rationnel ou on on met le resultat dans rat1 ?
		return res1.Add(res1, res2), nil
	case "difference":
		res1, err1 := TermToRat(arg1)
		if err1 != nil {
			return new(big.Rat), err1
		}
		res2, err2 := TermToRat(arg2)
		if err2 != nil {
			return new(big.Rat), err2
		}
		return res1.Add(res1, new(big.Rat).Mul(res2, big.NewRat(-1,1))), nil
	}
	case "product":
		res1, err1 := TermToRat(arg1)
		if err1 != nil {
			return new(big.Rat), err1
		}
		res2, err2 := TermToRat(arg2)
		if err2 != nil {
			return new(big.Rat), err2
		}
		return res1.Mul(res1, res2), nil
	}
	case "quotient_t":
		res1, err1 := TermToRat(arg1)
		if err1 != nil {
			return new(big.Rat), err1
		}
		res2, err2 := TermToRat(arg2)
		// Rajout du cas : res2 == 0
		if err2 != nil || res2 == 0 {
			return new(big.Rat), err2
		}
		return res1.Mul(res1, new(big.Rat).Inv(res2)), nil
	}
	
	return new(big.Rat), nil
}
