package main

import "fmt"

func main() {
	str1 := "Hello"
	b1 := []byte(str1)
	fmt.Printf("%#v\n", b1)

	str2 := "Привет"
	b2 := []rune(str2)
	fmt.Printf("%#v\n", b2)
	b3 := []byte(str2)
	fmt.Printf("%#v\n", b3)
	fmt.Printf("%#v\n", string(b2))
	fmt.Printf("%#v\n", string(b3))
}
