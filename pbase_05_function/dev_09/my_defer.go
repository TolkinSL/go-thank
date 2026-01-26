package main

import "fmt"

func main() {
	fmt.Println("Main", myDiv(4, 2))
}

func myDiv(a, b int) int {
	defer fmt.Println("myDiv defer 1")
	defer fmt.Println("myDiv defer 2")
	x := a / b
	fmt.Println("myDiv x:", x)
	return x
}