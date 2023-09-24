package main

import (
	"fmt"
	"math/big"
	"time"
)

func main(){
	start := time.Now()
	
	googol := big.NewInt(int64(10))

	googol.Exp(googol, big.NewInt(int64(100)), nil)

	fmt.Println(googol)
	fmt.Println("Took ", time.Since(start), " to run")
}