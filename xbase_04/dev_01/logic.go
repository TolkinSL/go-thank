package main

import (
	"fmt"
	"math/rand"
)

func main() {
	a := 2

	if a < 5 {
		fmt.Println("a меньше 5")
	} else if a == 5 {
		fmt.Println("a равно 5")
	} else {
		fmt.Println("a больше 5")
	}

	num := rand.Intn(100)

	if num > 50 {
		fmt.Printf("Выпало Большое число %d\n", num)
	} else {
		fmt.Printf("Выпало Маленькое число %d\n", num)
	}

	//Если переменная не нужна или временная (обработка ошибок)
	//if num := rand.Intn(100); num > 50 {}
	//if err := doError; err != nil {}
}