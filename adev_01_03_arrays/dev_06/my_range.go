package main

import "fmt"

func main() {

	myNums := []int{2, 4, 7}
	var mySum int

	for _, num := range myNums {
		mySum += num
	}
	fmt.Println("Sum []int :", mySum)

	for i, num := range myNums {
		if num == 7 {
			fmt.Println("Found 7 on idx:", i)
		}
	}

	myMap := map[string]string{
		"a": "a1",
		"b": "b1",
		"c": "c3",
	}

	for key, val := range myMap {
		fmt.Println(key, val)
	}

	for key := range myMap {
		fmt.Println(key)
	}

	myStr := "Привет!"

	for idx, char := range myStr {
		fmt.Println(idx, char, string(char))
	}
	fmt.Println()
	
	for i := 0; i < len(myStr); i++ {
		fmt.Println("byte ", myStr[i], " ", string(myStr[i]), " ")
	}
	fmt.Println()

	myStr = "Hello"

	for i := 0; i < len(myStr); i++ {
		fmt.Println("byte ", myStr[i], " ", string(myStr[i]), " ")
	}
	fmt.Println()
}
