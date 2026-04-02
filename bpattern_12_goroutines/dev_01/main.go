package main

import (
	"context"
	"fmt"
	"time"
)

func count(ctx context.Context) chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 0; i < 10; i++ {
			select {
			case ch <- i:
			case <-ctx.Done():
				fmt.Println("<-ctx.Done()")
				return
			}
		}

	}()

	return ch
}

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	ch := count(ctx)

	for val := range ch {
		if val > 5 {
			cancel()
			break
		}
		fmt.Println(val)
	}

	time.Sleep(time.Second)
}
