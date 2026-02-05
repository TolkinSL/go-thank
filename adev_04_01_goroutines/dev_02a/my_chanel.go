package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	message := make(chan int)

	fmt.Println("Goroutines run")

	go func() {
		defer wg.Done()
		for x := 0; x < 10; x++ {
			message <- x
		}
		close(message)

	}()

	go func() {
		defer wg.Done()

		fmt.Printf("%#v\n", message)

		// for {
		// 	val := <-message
		// 	fmt.Println("go func2 Приняла сообщение:", val)
		// 	if val == 9 {
		// 		break
		// 	}
		// }

		for val := range message {
			fmt.Println("go func2 Приняла сообщение:", val)
		}
	}()

	wg.Wait()
	fmt.Println("Main complete")
}
