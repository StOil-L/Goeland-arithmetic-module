/******************************************************************************

                            Online Go Lang Compiler.
                Code, Compile, Run and Debug Go Lang program online.
Write your code in this editor and press "Run" button to execute it.

*******************************************************************************/

package main

import (
    "fmt"
    "strings"
    "regexp"
    "strconv"
)

type Symbole interface {
    toString() string
}

type Sup struct {}
type Inf struct {}
type Eq struct {}
type SupEq struct {}
type InfEq struct {}

func (s Sup) toString() string {
    return ">"
}

func (i Inf) toString() string {
    return "<"
}

func (e Eq) toString() string {
    return "="
}
func (se SupEq) toString() string {
    return ">="
}
func (ie InfEq) toString() string {
    return "<="
}

type Equation struct {
    leftArray  []float64
    rightArray []float64
    constexpr  float64
    symbol     Symbol

}

/*func MakeEquationFromT(t T) Equation {
    
}*/

func main() {
}

