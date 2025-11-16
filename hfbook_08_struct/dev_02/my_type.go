package main

import "fmt"

type Subscriber struct {
	name string
	rate float64
	active bool
}

func main() {
	var sub1 Subscriber
	sub1.name = "Aman"
	fmt.Println("Name:", sub1.name)

	var sub2 Subscriber
	sub2.name = "Beth"
	fmt.Println("Name:", sub2.name)	
}
