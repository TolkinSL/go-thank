package main

import "fmt"

func main() {
	var b bool = true
	fmt.Println(b)

	var s string = "hello"
	fmt.Println(s)

	var i int = 42
	fmt.Println(i)

	var f float64 = 12.3
	fmt.Println(f)

	var one, two int = 3, 5
	fmt.Println(one, two)

	var myS = "myString"
	fmt.Printf("%v - %T\n", myS, myS)
}

