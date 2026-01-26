package main

import "fmt"

type User struct {
	Name string
}

func main() {
	user := User{
		Name: "Vaisa",
	}

	fmt.Println("Main user:", user.Name)
	updateUser(user)
	fmt.Println("Main user:", user.Name)
}

func updateUser(u User) {
	u.Name = "Oleg"

	fmt.Println("UpdtUser:", u.Name)
	resetUser(&u)
	fmt.Println("UpdtUser:", u.Name)
}

func resetUser(u *User) {
	u.Name = "Default user"

	fmt.Println("ResetUser:", u.Name)
}