package main

import "fmt"

func main() {
	str := "My String"
	b := []byte(str)

	fmt.Println(string(b[0]))

	var re rune = '€'
	fmt.Printf("%[1]c\n", re)
	fmt.Printf("As a string: %s and as a character: %c\n", re, re)
}
