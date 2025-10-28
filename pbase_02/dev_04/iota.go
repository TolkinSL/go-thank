package main

import "fmt"

func main() {
	const (
		Red = iota
		Green
		Blue
	)

	const (
		mon = iota
		tue
		wed
		thu
		fri
	)

	fmt.Println(Red, Green, Blue)
	fmt.Println(tue, fri)

	myDay := tue
	fmt.Println(2 == tue)
	fmt.Println(1 == tue)
	fmt.Println(myDay == tue)

	const (
		a = iota + 7
		b
		c
	)
	fmt.Println(a, b, c)

	const (
		_ = iota + 7 // Игнорирование 0
		f 
		g
		h
	)
	fmt.Println(f, g, h)
}