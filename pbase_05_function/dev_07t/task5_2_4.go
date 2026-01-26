package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(isIntSqrt(4))
	fmt.Println(isIntSqrt(5))

}

func isIntSqrt(num int) bool {
	mySqrt := int(math.Sqrt(float64(num)))
	
	if mySqrt * mySqrt != num {
		return false
	}

	return true

}