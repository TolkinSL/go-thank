package main

import "fmt"

func main() {
	var firstName, lastName string
	var age int

	fmt.Print("Введите ваше имя и фамилию: ")
	fmt.Scanln(&firstName, &lastName)

	fmt.Print("Введите ваш возраст: ")
	fmt.Scanln(&age)

	fmt.Printf("Приятно познакомиться %s %s, ваш возраст через 5 лет %d\n", firstName, lastName, age + 5)

	var str string
	fmt.Scanln(&str)
	fmt.Println(str)
}