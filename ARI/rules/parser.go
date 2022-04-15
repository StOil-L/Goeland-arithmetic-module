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
	"math"
	"strconv"
	"strings"
)

var zero_rat = &big.Rat{}

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

func FormToPred(f types.Form) /*types.Pred*/ int{

	fmt.Printf("ici \n")

	return 0
}


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
			return strconv.Atoi(ttypes.GetID().GetName())
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


	case "quotient":
		fmt.Printf("erreur quotient pas défini sur les entiers \n")
		var err error
		return 0, err

	case "quotient_e":
		res1, err1 := TermToInt(arg1)
		if err1 != nil {
			return 0, err1
		}
		res2, err2 := TermToInt(arg2)
		// Rajout du cas : res2 == 0
		if err2 != nil || res2 == 0 {
			return 0, err2
		}
		quo := res1/res2
		if(res1/res2 < 0){
			quo= res1/res2 -1 
		}
		return quo,nil


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
		
		return res1/res2,nil


	case "quotient_f":
		res1, err1 := TermToInt(arg1)
		if err1 != nil {
			return 0, err1
		}
		res2, err2 := TermToInt(arg2)
		// Rajout du cas : res2 == 0
		if err2 != nil || res2 == 0 {
			return 0, err2
		}
		
		return int(math.Floor(float64((res1*1.)/res2))),nil


	case "remainder_e":
		res1, err1 := TermToInt(arg1)
		if err1 != nil {
			return 0, err1
		}
		res2, err2 := TermToInt(arg2)
		// Rajout du cas : res2 == 0
		if err2 != nil || res2 == 0 {
			return 0, err2
		}

		remain:= res1%res2

		if(res1/res2<0){
			remain+=res2
		}

		return remain, nil
	

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

		return res1%res2, nil

	case "remainder_f":
		res1, err1 := TermToInt(arg1)
		if err1 != nil {
			return 0, err1
		}
		res2, err2 := TermToInt(arg2)
		// Rajout du cas : res2 == 0
		if err2 != nil || res2 == 0 {
			return 0, err2
		}

		return /*res1/(res2*quotient_f(res1,res2))*/res1, nil
/*
	case "uminus":
		arg := f.GetArgs()[0]
		res1, err1 := TermToInt(arg1)
		if err1 != nil {
			return 0, err1
		}
		if res1 == 0{
			res1 = 0
		} 
		else{
			res1 = res1*-1
		} 

		return res1,nil

	case "floor":
		arg := f.GetArgs()[0]
		res1, err1 := TermToInt(arg1)
		if err1 != nil {
			return 0, err1
		}

		return int(math.Floor(res1)),nil

	case "ceiling":
		arg := f.GetArgs()[0]
		res1, err1 := TermToInt(arg1)
		if err1 != nil {
			return 0, err1
		}

		return int(math.Ceil(res1)), res1
	
	case "truncate":
		arg := f.GetArgs()[0]
		res1, err1 := TermToInt(arg1)
		if err1 != nil {
			return 0, err1
		}
		if(res1 > 0){
			res1 = int(math.Floor(res1))
		} 
		else{
			res1 = int(math.Ceil(res1))
		} 

	case "round":
		arg := f.GetArgs()[0]
		res1, err1 := TermToInt(arg1)
		if err1 != nil {
			return 0, err1
		}

		if(res1 - int(math.Floor(res1)) >= 0.5){
			res1 = int(math.Floor(res1)+1)
		} 
		else{
			res1 = int(math.Floor(res1))
		} 

		return res1,nil
*/	
	}
	return 0, nil
}

// TODO : tout
func TermToRat(t types.Term) (*big.Rat, error) {
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

			fmt.Printf("Error TermToRat")
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

func funToRat(f types.Fun) (*big.Rat, error) {
	arg1 := f.GetArgs()[0]
	arg2 := f.GetArgs()[1]
	switch f.GetID().GetName() {
	case "sum":
		res1, err1 := TermToRat(arg1)
		if err1 != nil {
			return nil, err1
		}
		res2, err2 := TermToRat(arg2)
		if err2 != nil {
			return nil, err2
		}
		// On créé un nouveau rationnel ou on on met le resultat dans rat1 ?
		return res1.Add(res1, res2), nil
	case "difference":
		res1, err1 := TermToRat(arg1)
		if err1 != nil {
			return nil, err1
		}
		res2, err2 := TermToRat(arg2)
		if err2 != nil {
			return nil, err2
		}
		return res1.Add(res1, res2.Mul(res2, big.NewRat(-1, 1))), nil

	case "product":
		res1, err1 := TermToRat(arg1)
		if err1 != nil {
			return nil, err1
		}
		res2, err2 := TermToRat(arg2)
		if err2 != nil {
			return nil, err2
		}
		return res1.Mul(res1, res2), nil

	case "quotient_t":
		res1, err1 := TermToRat(arg1)
		if err1 != nil {
			return nil, err1
		}
		res2, err2 := TermToRat(arg2)
		// Rajout du cas : res2 == 0
		if err2 != nil || (res2.Cmp(zero_rat) == 0) {
			return nil, err2
		}
		return res1.Mul(res1, new(big.Rat).Inv(res2)), nil
	}

	return new(big.Rat), nil
}
