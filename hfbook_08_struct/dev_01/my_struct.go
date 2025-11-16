package main

import "fmt"

func main() {
	var myStruct1 struct {
		number float64
		word   string
		toggle bool
	}
	fmt.Printf("%+v\n", myStruct1)

	myStruct1.number = 3.14
	myStruct1.word = "Go"
	fmt.Println(myStruct1.word)

	var subscriber struct {
		name string
		rate float64
		active bool
	}

	subscriber.name = "Aman"
	subscriber.rate = 4.99
	subscriber.active = true
	fmt.Println("Name:", subscriber.name)
}
