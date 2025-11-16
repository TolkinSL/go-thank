package main

import "fmt"

type Subscriber struct {
	name   string
	rate   float64
	active bool
}

func defaulSubscriber(name string) *Subscriber {
	var s Subscriber
	s.name = name
	s.rate = 5.85
	s.active = true
	return &s
}

func printInfo(s *Subscriber) {
	fmt.Println("Name:", s.name)
}

func applyDiscount(num float64, s *Subscriber) {
	s.rate += num
}

func main() {
	sub1 := Subscriber{
		name:   "Vasia",
		rate:   3.6,
		active: true,
	}

	fmt.Printf("%+v\n", sub1)

	sub2 := defaulSubscriber("Olya")
	fmt.Printf("%+v\n", sub2)
	fmt.Printf("%+v\n", *sub2)

	applyDiscount(1.5, &sub1)
	fmt.Printf("%+v\n", sub1)

	applyDiscount(2, sub2)
	fmt.Printf("%+v\n", sub2)
}
