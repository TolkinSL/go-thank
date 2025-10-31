package main

import (
	"develop/pbase_11_struct/dev_06/model"
	"fmt"
	// "log"
)

func main() {
	fmt.Printf("Hello\n")

	p1, err := model.NewPerson("Vasia", "Ivanov", 35)
	if err != nil {
		fmt.Printf("Error %v\n", err)
	}
	fmt.Printf("%+v\n", p1)

	p2, err := model.NewPerson("Olia", "Petrova", 8)
	if err != nil {
		// log.Fatalf("Error in data %v", err)
		fmt.Printf("Error in data: %v\n", err)
	}
	fmt.Printf("%+v\n", p2)
}
