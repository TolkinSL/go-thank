package main

import "fmt"

func main() {
	var val1 interface{} = "Hello"
	fmt.Printf("%#T\n", val1)
	val1 = 12
	fmt.Printf("%#T\n", val1)
	val1 = "12"
	fmt.Printf("%#T\n", val1)

	if str, ok := val1.(string); ok {
		fmt.Println("This String: ", str)
	} else {
		fmt.Println("Not String")
	}
}
