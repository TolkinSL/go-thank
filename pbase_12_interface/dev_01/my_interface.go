package main

import "fmt"

type Dog struct {}
func (Dog) Speak() string {
	return "Gaff"
}
func (Dog) Growl() string {
	return "Rrrr"
}

type Cat struct{}
func (Cat) Speak() string {
	return "Meow"
}

type Speaker interface {
	Speak() string
}

// Полиморфизм
func makeSound(s Speaker) string {
	return s.Speak()
}

func main() {
	dog := Dog{}
	cat := Cat{}

	fmt.Println(makeSound(cat))
	fmt.Println(makeSound(dog))
}