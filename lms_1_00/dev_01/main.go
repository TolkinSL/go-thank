package main

import "fmt"

func main() {
	var name string

	fmt.Print("Input name: ")
	fmt.Scanln(&name)

	result := fmt.Sprintf("Hello %s", name)
	fmt.Println(result)
}