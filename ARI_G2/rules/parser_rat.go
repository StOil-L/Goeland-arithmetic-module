/*****************/
/* parser_rat.go */
/*****************/
/**
* This file provides function to parse a arithmetic term into a *big.Rat.
**/

package ari

import (
	"ARI/types"
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

var zero_rat = &big.Rat{}

func termToRat(t types.Term) (*big.Rat, error) {
	switch ttypes := t.(type) {
	case types.Fun:
		// Constante
		if len(ttypes.GetArgs()) == 0 {
			// Julie : Ici je split le nombre (de la forme a/b) par "/" et je créé le rat correspondant
			// Cas particlier pour 0
			n_d := strings.Split(ttypes.GetID().ToString(), "/")
			if len(n_d) == 2 {
				n, err_n := strconv.ParseInt(n_d[0], 10, 64)
				if err_n != nil {
					return zero_rat, err_n
				}
				d, err_d := strconv.ParseInt(n_d[1], 10, 64)
				if err_d != nil {
					return zero_rat, err_d
				}
				return big.NewRat(n, d), nil
			}
			if len(n_d) == 1 && n_d[0] == "0" {
				return zero_rat, nil
			}

			fmt.Printf("Error termToRat")
		} else {
			if res, err := funToRat(ttypes); err == nil {
				return res, nil
			} else {
				return nil, err
			}
		}

	default:
		fmt.Println("Error in conversion Rat: format not found")
		return nil, errors.New("error Rat")
	}

	return nil, errors.New("error Rat")
}

// TODO : compléter la fonction
func funToRat(f types.Fun) (*big.Rat, error) {
	arg1 := f.GetArgs()[0]
	arg2 := f.GetArgs()[1]
	switch f.GetID().GetName() {
	case "sum":
		res1, err1 := termToRat(arg1)
		if err1 != nil {
			return nil, err1
		}
		res2, err2 := termToRat(arg2)
		if err2 != nil {
			return nil, err2
		}
		// On créé un nouveau rationnel ou on on met le resultat dans rat1 ?
		return res1.Add(res1, res2), nil
	case "difference":
		res1, err1 := termToRat(arg1)
		if err1 != nil {
			return nil, err1
		}
		res2, err2 := termToRat(arg2)
		if err2 != nil {
			return nil, err2
		}
		return res1.Add(res1, res2.Mul(res2, big.NewRat(-1, 1))), nil

	case "product":
		res1, err1 := termToRat(arg1)
		if err1 != nil {
			return nil, err1
		}
		res2, err2 := termToRat(arg2)
		if err2 != nil {
			return nil, err2
		}
		return res1.Mul(res1, res2), nil

	case "quotient_t":
		res1, err1 := termToRat(arg1)
		if err1 != nil {
			return nil, err1
		}
		res2, err2 := termToRat(arg2)
		// Rajout du cas : res2 == 0
		if err2 != nil || (res2.Cmp(zero_rat) == 0) {
			return nil, err2
		}
		return res1.Mul(res1, new(big.Rat).Inv(res2)), nil
	}

	return new(big.Rat), nil
}
