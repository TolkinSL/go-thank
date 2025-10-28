package main

import "fmt"

type User struct {
	firstName string
}

func main() {
	user1 := User{
		firstName: "Vasia",
	}

	copyValue(user1)
	fmt.Println("copy:", user1.firstName)
	modifyValue(&user1)
	fmt.Println("modify:", user1.firstName)
}

func copyValue(u User) {
	u.firstName = "Olia"
}

func modifyValue(u *User) {
	u.firstName = "Petia"
}