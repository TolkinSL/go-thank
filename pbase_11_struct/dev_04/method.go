package main

import "fmt"

type Name struct {
	first string
	last  string
}

type User struct {
	Name
	birthYear int
}

func (u User) hello() {
	fmt.Println("Hello -", u.first)
}


func main() {
	user1 := User{
		Name: Name{
			first: "Vasia",
			last:  "Ivanov",
		}, 
		birthYear: 1987,
	}

	user1.hello()

}
