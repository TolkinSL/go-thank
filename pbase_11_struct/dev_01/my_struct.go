package main

import "fmt"

type User struct {
	firstName string
	lastName string
	birthYear int
}

func main() {
	user1 := User{
		firstName: "Vasia",
		lastName: "Ivanov",
		birthYear: 1980,
	}

	fmt.Println(user1)
	fmt.Printf("%+v\n", user1)

	var user2 User
	fmt.Printf("%+v\n", user2)

	user3 := &User{}
	fmt.Printf("%+v\n", user3)
	fmt.Printf("%+v\n", *user3)
}