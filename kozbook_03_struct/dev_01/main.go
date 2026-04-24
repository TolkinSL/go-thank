package main

import "fmt"

type Animal struct {
	name string
}

func (a Animal) speak() string{
	switch a.name {
	case "cat":
		return "Meow!"
	case "dog":
		return "Gav!"
	default:
		return "?..."
	}
}

func main() {
	// Анонимная структура
	anoAnimal := struct {
		name string
		speak func() string
	}{
		name: "cat",
		speak: func() string{
			return "Meow!"
		},
	}

	fmt.Println(anoAnimal.name, anoAnimal.speak())

	pet1 := Animal{
		name: "cat",
	}

	fmt.Printf("%s say %s\n", pet1.name, pet1.speak())

	pet2 := Animal{
		name: "mouse",
	}

	fmt.Printf("%s say %s\n", pet2.name, pet2.speak())
}
