package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

// func say(wg *sync.WaitGroup, id int, phrase string) {

func say(id int, phrase string) {
	for _, word := range strings.Fields(phrase) {
		fmt.Printf("%d Computer say: %s\n", id, word)
		time.Sleep(500 * time.Millisecond)	
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	
	go func() {
		defer wg.Done()
		// say(&wg, 1, "Hello Go World!")
		say(1, "Hello Go World!")
	}()

	go func() {
		defer wg.Done()
		// say(&wg, 2, "Go Is Awesome!")
		say(2, "Go Is Awesome!")
	}()

	wg.Wait()
}