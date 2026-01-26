package main

import "fmt"

func someFn() string {
	i := 1
	fmt.Println("someFn i:", i)
	i += 1
	fmt.Println("someFn i:", i)
	return fmt.Sprintf("someFn Return i: %d\n", i)
}
func main() {
	defer fmt.Println("Start 1")
	defer fmt.Println("Start 2")
	defer func() {
		fmt.Println(someFn())
	}()
	fmt.Println("Start Main")
}
