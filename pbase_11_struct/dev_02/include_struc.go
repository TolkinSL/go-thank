package main

import "fmt"

type Name struct {
	firstName string
	lastName string
}

// композиция
type User struct {
	userName Name
	birthYear int
}

// встраивание (embedding)
type UserAdmin struct {
	Name
	birthYear int
}

func main() {
	user1 := User{
		userName: Name{
			firstName: "Vasia",
			lastName: "Ivanov",
		},
		birthYear: 1980,
	}

	fmt.Printf("%+v\n", user1)
	fmt.Println(user1.birthYear)
	fmt.Println(user1.userName.firstName)

	user2 := UserAdmin{
		Name: Name{
			firstName: "Petia",
			lastName: "Petrov",
		},
		birthYear: 1985,
	}

	fmt.Printf("%+v\n", user2)
	fmt.Println(user2.birthYear)
	fmt.Println(user2.firstName)

	user2.firstName = "Olia"
	user2.Name.lastName = "Ivanova"
	fmt.Printf("%+v\n", user2)
}