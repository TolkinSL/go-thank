package main

import "fmt"

type Part struct {
	description string
	count       int
}

func showInfo(p Part) {
	fmt.Println("Desc:", p.description)
	fmt.Println("Count:", p.count)
}

func minimumOrder(description string) Part {
	var p Part
	p.description = description
	p.count = 100
	return p
}

func main() {
	var bolt Part
	bolt.description = "Hex bolts"
	bolt.count = 24
	showInfo(bolt)

	// var order1 Part
	// order1 = minimumOrder("Bolts")
	order1 := minimumOrder("Bolts")
	fmt.Println("Min order:", order1.description, order1.count)
}
