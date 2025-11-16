package main

import "fmt"

type Subscriber2 struct {
	name   string
	rate   float64
	active bool
}

func printInfo(s Subscriber2) {
	fmt.Println(s.name, s.rate, s.active)
}

func defaultSub(name string) Subscriber2 {
	var s Subscriber2
	s.name = name
	s.rate = 5.9
	s.active = true

	return s
}

func main() {
	sub1 := defaultSub("Aman")
	sub1.rate = 4
	printInfo(sub1)

	sub2 := defaultSub("Beth")
	printInfo(sub2)
}
