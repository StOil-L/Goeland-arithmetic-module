/*****************/
/* parser_int.go */
/*****************/
/**
* This file provides function to parse a arithmetic term into an int.
**/

package ari

import (
	typing "ARI/polymorphism"
	"ARI/types"
	"errors"
	"fmt"
	"math"
	"strconv"
)

/* Transform a term into an int */
func termToInt(t types.Term) (int, error) {
	switch ttypes := t.(type) {
	case types.Fun:
		// Constante
		if len(ttypes.GetArgs()) == 0 {
			return strconv.Atoi(ttypes.GetID().GetName())
		} else {
			if res, err := FunToInt(ttypes); err == nil {
				return res, nil
			} else {
				return 0, err
			}
		}
	default:
		fmt.Println("Error in conversion : format not found")
		return 0, errors.New("Error")
	}
}

/* Transform a fun into an int */
/**
* TODO : vous devez definir arg1 et agr2 dans chaque sous fonction.
* En effet, votre functino peut ête uniare (genre uminus), et dans ce cas ça va planter si vous faites f.GetArgs()[1]
* Parce que dans le cas d'une fonction unaire, g.GetArgs est de taille 1 !
*	===> Fais par Margaux
* Les 2 returns sur la dernière ligne, c'est pour qu'ils se sentent moins seuls ?
**/
func FunToInt(f types.Fun) (int, error) {
	switch f.GetID().GetName() {

	case "sum":
		arg1 := f.GetArgs()[0]
		arg2 := f.GetArgs()[1]
		res1, err1 := termToInt(arg1)
		if err1 != nil {
			return 0, err1
		}
		res2, err2 := termToInt(arg2)
		if err2 != nil {
			return 0, err2
		}
		return res1 + res2, nil

	case "difference":
		arg1 := f.GetArgs()[0]
		arg2 := f.GetArgs()[1]
		res1, err1 := termToInt(arg1)
		if err1 != nil {
			return 0, err1
		}
		res2, err2 := termToInt(arg2)
		if err2 != nil {
			return 0, err2
		}
		return res1 - res2, nil

	case "product":
		arg1 := f.GetArgs()[0]
		arg2 := f.GetArgs()[1]
		res1, err1 := termToInt(arg1)
		if err1 != nil {
			return 0, err1
		}
		res2, err2 := termToInt(arg2)
		if err2 != nil {
			return 0, err2
		}
		return res1 * res2, nil

		//à voir`
	case "quotient":
		fmt.Printf("erreur quotient pas défini sur les entiers \n")
		var err error
		return 0, err

	case "quotient_e":
		arg1 := f.GetArgs()[0]
		arg2 := f.GetArgs()[1]
		res1, err1 := termToInt(arg1)
		if err1 != nil {
			return 0, err1
		}
		res2, err2 := termToInt(arg2)
		// Rajout du cas : res2 == 0
		if err2 != nil || res2 == 0 {
			return 0, err2
		}
		quo := res1 / res2
		if res1/res2 < 0 {
			quo = res1/res2 - 1
		}
		return quo, nil

	case "quotient_t":
		arg1 := f.GetArgs()[0]
		arg2 := f.GetArgs()[1]
		res1, err1 := termToInt(arg1)
		if err1 != nil {
			return 0, err1
		}
		res2, err2 := termToInt(arg2)
		// Rajout du cas : res2 == 0
		if err2 != nil || res2 == 0 {
			return 0, err2
		}

		return res1 / res2, nil

	case "quotient_f":
		arg1 := f.GetArgs()[0]
		arg2 := f.GetArgs()[1]
		res1, err1 := termToInt(arg1)
		if err1 != nil {
			return 0, err1
		}
		res2, err2 := termToInt(arg2)
		// Rajout du cas : res2 == 0
		if err2 != nil || res2 == 0 {
			return 0, err2
		}
		fmt.Printf("%d \n", int(math.Floor(float64((res1*1.)/res2))))
		return int(math.Floor(float64((res1 * 1.) / res2))), nil

	case "remainder_e":
		arg1 := f.GetArgs()[0]
		arg2 := f.GetArgs()[1]
		res1, err1 := termToInt(arg1)
		if err1 != nil {
			return 0, err1
		}
		res2, err2 := termToInt(arg2)
		// Rajout du cas : res2 == 0
		if err2 != nil || res2 == 0 {
			return 0, err2
		}

		remain := res1 % res2

		if res1/res2 < 0 {
			remain += res2
		}

		return remain, nil

	case "remainder_t":
		arg1 := f.GetArgs()[0]
		arg2 := f.GetArgs()[1]
		res1, err1 := termToInt(arg1)
		if err1 != nil {
			return 0, err1
		}
		res2, err2 := termToInt(arg2)
		// Rajout du cas : res2 == 0
		if err2 != nil || res2 == 0 {
			return 0, err2
		}

		return res1 % res2, nil

	case "remainder_f":
		arg1 := f.GetArgs()[0]
		arg2 := f.GetArgs()[1]
		res1, err1 := termToInt(arg1)
		if err1 != nil {
			return 0, err1
		}
		res2, err2 := termToInt(arg2)
		// Rajout du cas : res2 == 0
		if err2 != nil || res2 == 0 {
			return 0, err2
		}
		//peut être faire ça sur chaque remainder ?
		Rquotient_f := types.MakeFun(types.MakerId("quotient_f"), []types.Term{arg1, arg2}, typing.GetTypeScheme("quotient_f", typing.MkTypeCross(tInt, tInt)))
		quotient_f, _ := FunToInt(Rquotient_f)
		return res1 - (res2 * quotient_f), nil

	case "uminus":
		arg1 := f.GetArgs()[0]
		res1, err1 := termToInt(arg1)
		if err1 != nil {
			return 0, err1
		}
		if res1 == 0 {
			res1 = 0
		} else {
			res1 = res1 * -1
		}

		return res1, nil

	case "floor":
		arg1 := f.GetArgs()[0]
		res1, err1 := termToInt(arg1)
		if err1 != nil {
			return 0, err1
		}

		return int(math.Floor(float64(res1))), nil

	case "ceiling":
		arg1 := f.GetArgs()[0]
		res1, err1 := termToInt(arg1)
		if err1 != nil {
			return 0, err1
		}

		return int(math.Ceil(float64(res1))), err1

	case "truncate":
		arg1 := f.GetArgs()[0]
		res1, err1 := termToInt(arg1)
		if err1 != nil {
			return 0, err1
		}
		if res1 > 0 {
			res1 = int(math.Floor(float64(res1)))
		} else {
			//pas sur de ça
			res1 = int(math.Ceil(float64(res1)))
		}

	case "round":
		arg1 := f.GetArgs()[0]
		res1, err1 := termToInt(arg1)
		if err1 != nil {
			return 0, err1
		}
		if float64(res1-int(math.Floor(float64(res1)))) >= 0.5 {
			res1 = int(math.Ceil(float64(res1)))
		} else {
			res1 = int(math.Floor(float64(res1)))
		}
		return res1, nil

	}

	return 0, nil
}
