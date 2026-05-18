package main

import "fmt"

func main() {
	fmt.Println(test1())
	fmt.Println(test1())

	//------

	a := 3
	fmt.Println(inc(&a))
	inc(&a)
	fmt.Println(a)

}

func test1() *int {
	v := 1

	return &v
}

func inc(p *int) int {
	*p += 1
	return *p
}