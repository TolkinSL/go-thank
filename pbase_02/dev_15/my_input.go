package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var firstName string
	fmt.Println("Input name")
	fmt.Scan(&firstName)
	fmt.Println(firstName)
	fmt.Println()

	fmt.Println("Bufio scan")
	var tempData []string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	tempData = strings.Split(scanner.Text(), " ")
	fmt.Println(tempData)
}