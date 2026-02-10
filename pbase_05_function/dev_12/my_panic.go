package main

import "fmt"

func main() {
	defer handlePanic()
	panic("My panic 111")
	fmt.Println("Код который Не выполняется после паники")
}

func handlePanic() {
	if r := recover(); r != nil {
		fmt.Println("Обработка паники 111")
		fmt.Printf("Обработка - %#v\n", r)
	}
}