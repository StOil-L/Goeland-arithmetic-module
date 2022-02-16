package main

import (
	"fmt"
	"math/big"
)

func main() {
	x,_:=new(big.Rat).SetString("-2")
	y,_:=new(big.Rat).SetString("-2.02")
	fmt.Println(x.IsInt())
	fmt.Println(y.IsInt())

}
