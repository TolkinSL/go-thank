package main

import "fmt"

func main() {
	i := 1

	fmt.Println("i")
	for i <= 3 {
		fmt.Println("i-", i)
		i += 1
	}

	fmt.Println("j")
	for j := 5; j < 10; j += 1 {
		fmt.Println("j-", j)
	}

	fmt.Println("range")
	n := 10
	for i := range n {
		fmt.Println("range-", i)
	}

	myB := 0
	fmt.Println("break")
	for {
		fmt.Println("break-", myB)
		myB += 1

		if myB > 4 {
			break
		}
	}

	fmt.Println("continue")
	for n := 0; n < 6; n += 1 {
		if n%2 != 0 {
			continue
		}
		fmt.Println("continue", n)
	}
}
