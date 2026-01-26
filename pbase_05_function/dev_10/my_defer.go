package main

import "fmt"

func main() {
	myFunc()
	fmt.Println("Main Sum:", mySum(2, 3))
	fmt.Println("Main Mult:", myMult(2, 2))
	fmt.Println("Main end")
}

func myFunc() {
	x := 0

	defer fmt.Println("defer1 x:", x)

	x = 3

	defer func(x int) {
		fmt.Println("defer2 x:", x)
	}(x)

	defer func() {
		fmt.Println("defer3 x:", x)
	}()

	x = 5
}

func mySum(a, b int) (sum int) {
	sum = a + b
	defer func() {
		sum += 1
		fmt.Println("deferSum:", sum)
	}()

	return
}

func myMult(a, b int) (int) {
	mult := a * b
	defer func() {
		mult += 1
		fmt.Println("deferMult:", mult)
	}()

	return mult
}
