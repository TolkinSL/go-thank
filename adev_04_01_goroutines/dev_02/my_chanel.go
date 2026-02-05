package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	message := make(chan string)

	fmt.Println("Goroutines run")

	go func() {
		defer wg.Done()
		fmt.Println("go func1 Отправила сообщение")
		message <- "go func1"
	}()

	go func() {
		defer wg.Done()
		var msg string

		msg = <-message
		fmt.Println("go func2 Приняла сообщение:", msg)
	}()

	wg.Wait()
	fmt.Println("Main complete")
}
