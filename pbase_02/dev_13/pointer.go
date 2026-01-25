package main

import "fmt"

func main() {
	x := 14
	fmt.Printf("&x %p\n", &x)

	p := &x
	fmt.Printf("p %p\n", p)

}
