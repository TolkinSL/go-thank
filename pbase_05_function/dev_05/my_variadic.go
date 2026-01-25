package main

import "fmt"

func main() {
	fmt.Println(sum(1, 3, 5))
}

// func sum(a int, b int, value ...int) {}

func sum(values ...int) int{
	var sum int

	for _, v := range values {
		sum += v
	}

	return sum
}
