package main

import "fmt"

func main() {
	SayHello("Vasia")
}

func SayHello(user string) {
	fmt.Println("Hello!", user)
}