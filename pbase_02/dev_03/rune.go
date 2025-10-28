package main

import "fmt"

func main() {
	var b byte = 196
	fmt.Println(b,string(b))

	s := "Go"
	fmt.Println(s[0], string(s[0]))

	// rune == int32

	r := '®' 
	fmt.Println(r, string(r))

	var rx rune = 128539
	fmt.Println(rx, string(rx))

	rz := '😛'
	fmt.Println(rz, string(rz)[0])
}