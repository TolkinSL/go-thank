package main

import "fmt"

func main() {
	var p *int
	fmt.Printf("%#v\n", p)
	// fmt.Printf("%#v\n", *p) // panic: runtime error: invalid memory address

	if p == nil {
		fmt.Println("p = nil")
	}
	fmt.Println()

	p1 := new(int)
	fmt.Printf("%#v\n", p1)
	fmt.Printf("%#v\n", *p1)
}
