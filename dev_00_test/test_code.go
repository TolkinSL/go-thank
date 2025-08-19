package main

import (
	"fmt"
)

func main() {
	x := 0b0101
	y := x << 1

	fmt.Printf("%8b", y)
}