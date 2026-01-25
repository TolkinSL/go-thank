package main

import "fmt"

func main() {
	value := 2
	
	myCopy(value)
	fmt.Println("Copy param:", value)

	myPointer(&value)
	fmt.Println("Pointer param:", value)

}

func myCopy(a int) {
	a += 2
}

func myPointer(a *int) {
	*a += 2
}