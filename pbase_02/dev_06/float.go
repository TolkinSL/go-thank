package main

import (
	"fmt"
	"math"
	"math/rand/v2"
)

func main() {
	var xInt int = 1
	// var divInt int = 0
	// test := xInt / divInt // Деление Int на 0 получается panic
	// fmt.Println(test)

	var xFlt float64 = 1
	var divFlt float64 = 0
	test := xFlt / divFlt // Деление float64 на 0 получается +Inf
	fmt.Println(test)

	test = 0 / divFlt // Деление 0 на 0 float64 получается NaN
	fmt.Println(test)

	myNaN := math.NaN()
	fmt.Println("NaN -", myNaN)
	fmt.Println("NaN is -", math.IsNaN(myNaN))

	xFlt = float64(xInt)
	fmt.Println(xFlt)

	xFlt = 3.6
	fmt.Println("Round", int(xFlt))
	fmt.Println("Round", int(math.Round(xFlt)))

	randomFlt := rand.Float64()
	fmt.Println("Random [0 to 1)", randomFlt)

	min := 5.0
	max := 10.0 
	randomFlt = rand.Float64()*(max + 1 - min) + min
	fmt.Println("Random [5.0 to 10]", randomFlt) //до 10 включительно
}
