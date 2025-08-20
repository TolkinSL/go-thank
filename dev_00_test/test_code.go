package main

import (
	"fmt"
)

func main() {
	x := 0b0101
	y := x << 1

	fmt.Printf("%8b\n", y)

	var b byte = 169
	fmt.Println(b,string(b))

	var b1 byte = 71
	fmt.Println(b1,string(b1))

	s := "hello"
	fmt.Println(s[0],string(s[0]))

	r := '©'
	fmt.Println(r,string(r))
}