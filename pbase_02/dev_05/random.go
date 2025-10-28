package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	min := 5
	max := 10
	myRand := rand.IntN(max - min) + min // [5, 10)
	fmt.Println(myRand)
}