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
	typing "ARI/polymorphism"
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
			if res, err := FunToRat(ttypes); err == nil {
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
func FunToRat(f types.Fun) (*big.Rat, error) {
	switch f.GetID().GetName() {
	case "sum":
		arg1 := f.GetArgs()[0]
		arg2 := f.GetArgs()[1]
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
		arg1 := f.GetArgs()[0]
		arg2 := f.GetArgs()[1]
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
		arg1 := f.GetArgs()[0]
		arg2 := f.GetArgs()[1]
		res1, err1 := termToRat(arg1)
		if err1 != nil {
			return nil, err1
		}
		res2, err2 := termToRat(arg2)
		if err2 != nil {
			return nil, err2
		}
		return res1.Mul(res1, res2), nil
		
	case "quotient_e":
		arg1 := f.GetArgs()[0]
		arg2 := f.GetArgs()[1]
		res1, err1 := termToRat(arg1)
		if err1 != nil {
			return nil, err1
		}
		res2, err2 := termToRat(arg2)
		if err2 != nil || (res2.Cmp(zero_rat) == 0) {
			return nil, err2
		}
		quo := res1.Mul(res1, new(big.Rat).Inv(res2))
		if (quo.Num()).Cmp(big.NewInt(0)) == -1 {
			quo = quo.Add(quo, big.NewRat(-1, 1))
		}
		return quo, nil

	case "quotient_t":
		arg1 := f.GetArgs()[0]
		arg2 := f.GetArgs()[1]
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
		
	case "quotient_f":
		arg1 := f.GetArgs()[0]
		arg2 := f.GetArgs()[1]
		res1, err1 := termToRat(arg1)
		if err1 != nil {
			return nil, err1
		}
		res2, err2 := termToRat(arg2)
		// Rajout du cas : res2 == 0
		if err2 != nil || (res2.Cmp(zero_rat) == 0) {
			return nil, err2
		}
		quo := res1.Mul(res1, new(big.Rat).Inv(res2))
		return big.NewRat((new(big.Int).Mul(new(big.Int).Quo(quo.Num(), quo.Denom()), quo.Denom())).Int64(), (quo.Denom()).Int64()), nil
		
	case "remainder_e":
		arg1 := f.GetArgs()[0]
		arg2 := f.GetArgs()[1]
		res1, err1 := termToRat(arg1)
		if err1 != nil {
			return nil, err1
		}
		res2, err2 := termToRat(arg2)
		// Rajout du cas : res2 == 0
		if err2 != nil || (res2.Cmp(zero_rat) == 0) {
			return nil, err2
		}
		Rquotient_e := types.MakeFun(types.MakerId("quotient_e"), []types.Term{arg1, arg2}, typing.GetTypeScheme("quotient_e", typing.MkTypeCross(tInt, tInt)))
		quotient_e, _ := FunToRat(Rquotient_e)
		rem := new(big.Rat).Set(new(big.Rat).Add(res1, new(big.Rat).Neg(new(big.Rat).Mul(res2, quotient_e))))
		return rem, nil
		
	case "remainder_t":
		arg1 := f.GetArgs()[0]
		arg2 := f.GetArgs()[1]
		res1, err1 := termToRat(arg1)
		if err1 != nil {
			return nil, err1
		}
		res2, err2 := termToRat(arg2)
		// Rajout du cas : res2 == 0
		if err2 != nil || (res2.Cmp(zero_rat) == 0) {
			return nil, err2
		}
		Rquotient_t := types.MakeFun(types.MakerId("quotient_t"), []types.Term{arg1, arg2}, typing.GetTypeScheme("quotient_t", typing.MkTypeCross(tInt, tInt)))
		quotient_t, _ := FunToRat(Rquotient_t)
		rem := new(big.Rat).Set(new(big.Rat).Add(res1, new(big.Rat).Neg(new(big.Rat).Mul(res2, quotient_t))))
		return rem, nil
		
	case "remainder_f":
		arg1 := f.GetArgs()[0]
		arg2 := f.GetArgs()[1]
		res1, err1 := termToRat(arg1)
		if err1 != nil {
			return nil, err1
		}
		res2, err2 := termToRat(arg2)
		// Rajout du cas : res2 == 0
		if err2 != nil || (res2.Cmp(zero_rat) == 0) {
			return nil, err2
		}
		Rquotient_f := types.MakeFun(types.MakerId("quotient_f"), []types.Term{arg1, arg2}, typing.GetTypeScheme("quotient_f", typing.MkTypeCross(tInt, tInt)))
		quotient_f, _ := FunToRat(Rquotient_f)
		rem := new(big.Rat).Set(new(big.Rat).Add(res1, new(big.Rat).Neg(new(big.Rat).Mul(res2, quotient_f))))
		return rem, nil
	
	// Rien ne change ?
	case "uminus":
		arg1 := f.GetArgs()[0]
		res1, err1 := termToRat(arg1)
		if err1 != nil {
			return zero_rat, err1
		}
		if res1.Num() == big.NewInt(0) {
			res1 = zero_rat
		} else {
			res1 = res1.Neg(res1)
		}

		return res1, nil
	
	}

	return new(big.Rat), nil
}
