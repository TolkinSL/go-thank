package main

import "fmt"

type Car struct {
	name string
	speed float64
}

func nitroSpeed(c *Car) {
	c.speed += 50
}

func main() {
	var ford Car
	ford.name = "Focus"
	ford.speed = 215

	nitroSpeed(&ford)
	fmt.Println("Name:", ford.name)
	fmt.Println("Speed:", ford.speed)
}