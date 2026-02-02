package main

import "fmt"

func div(x, y int) (int, int) {
	res1 := x / y
	res2 := x % y

	return res1, res2
}

func sum(nums ...int) int {
	var sum int

	for _, v := range nums {
		sum += v
	}

	return sum
}

func main() {
	a, b := div(5, 2)
	fmt.Println("5/4 =", a, b)

	_, b = div(4, 2)

	my_sum := sum(1, 2, 3)
	fmt.Println("sum1 =",my_sum)

	my_num := []int{2, 2, 3}
	my_sum = sum(my_num...)
	fmt.Println("sum1 =",my_sum)
}
