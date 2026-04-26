package main

import "fmt"

type character struct {
	name string
}

func (c *character) fixName(name string) {
	c.name = name
}

func main() {
	user1 := new(character)
	(*user1).name = "test"
	user1.name = "test"
	fmt.Println(user1.name)

	user1.fixName("Vasia")
	fmt.Println(user1.name)
	
	user2 := character{
		name: "test",
	}
	name := "Petia"

	user2.fixName(name)
	fmt.Println(user2.name)
}