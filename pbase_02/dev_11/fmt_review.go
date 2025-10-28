package main

import "fmt"

func main() {
	fmt.Printf("%#v\n", "Hello!")
	fmt.Printf("%#v\n", 7)

	str1 := fmt.Sprintf("Sprintf %#v\n", "Hello!")
	fmt.Print(str1)

	fmt.Printf("float %.2f\n", 3.14567)
	fmt.Printf("float %T\n", 3.14567)
	fmt.Printf("bool %t\n", true)
	fmt.Printf("bool %T\n", true)
	fmt.Printf("float %.2f%%\n", 3.14567)
}