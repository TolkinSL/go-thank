package main

import "fmt"

func main() {
	s := "Привет"
	for j := range s {
		fmt.Println(j)
	}
}