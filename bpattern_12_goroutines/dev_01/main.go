package main

import (
	"fmt"
	"time"
)

func count() chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 0; i < 5; i++ {
			ch <- i
		}
		
	}()

	return ch
}

func main() {
	ch := count()

	for val := range ch {
		fmt.Println(val)
	}

	time.Sleep(time.Second)
}
