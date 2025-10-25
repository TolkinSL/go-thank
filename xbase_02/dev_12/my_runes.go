package main

import "fmt"

func main() {
	var r1 rune = 'Я'
	fmt.Printf("Символ: %c\n", r1)
	fmt.Println(r1)
	fmt.Println(string(r1))

	var r2 rune = '\u1071'
	fmt.Printf("Символ: %c\n", r2)
	fmt.Println(r2)

	var r3 rune = rune(1071)
	fmt.Printf("Символ: %c\n", r3)
	fmt.Println(r3)

	str1 := "Hello"
	fmt.Printf("Число: %d - Символ: %c\n", str1[0], str1[0])

	str2 := "Привет"
	fmt.Printf("Число: %d - Символ: %c\n", str2[0], str2[0])
	runes2 := []rune(str2)
	fmt.Println(runes2)
	fmt.Println(string(runes2))
	fmt.Printf("Число: %d - Символ: %c\n", runes2[0], runes2[0])
	fmt.Printf("Число: %d - Символ: %c\n", runes2[1], runes2[1])
}