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

func (u *User) changeName(str string) {
	u.first = str
	fmt.Println("fChangName:", u.first)
}

type Person struct {
	firstName string
	lastName string
}

func (p Person) sayHello() {
	fmt.Println("Person:", p.firstName)
}

// Метод не использующий структуру (Редко используется)
func(Person) hello() string {
	return "Hello"
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
	user1.changeName("Olia")
	user1.hello()

	person1 := Person{}
	person1.hello()

}
