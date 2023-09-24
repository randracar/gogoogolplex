package main

import (
	"fmt"
	"time"
	"math/big"
)

func main(){
	start := time.Now()

	googol := big.NewInt(int64(10))
	googol.Exp(googol, big.NewInt(int64(100)), nil)

	fmt.Println("Googol: ", googol)
	fmt.Println("Calculating googolplex...")

	var googolplex big.Int
	googolplex.Exp(big.NewInt(int64(10)), googol, nil)
	
	fmt.Println(googolplex)
	fmt.Println("Took ", time.Since(start), " to calculate googolplex")
}