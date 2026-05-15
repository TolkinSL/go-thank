package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	myParam1 := ""
	myParam2 := ""

	sep := ""
	param := os.Args[1:]

	for _, p := range param {
		myParam1 += sep + p
		sep = " "
	}

	myParam2 = strings.Join(param, " ")
	
	fmt.Println("Program params:", myParam1)
	fmt.Println("Program params:", myParam2)
}