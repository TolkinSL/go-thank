package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	mySum := sum(2, 3)
	fmt.Println(mySum)

	strHello := "Привет Мир"
	myBytes, myRunes := countStr(strHello)
	fmt.Println("Bytes:", myBytes, "Runes:", myRunes)
}
func countStr(str string) (int, int) {
	myBytes := len(str)
	myRunes := utf8.RuneCountInString(str)

	return myBytes, myRunes
}

func sum(x, y int) int {
	result := x + y

	return result
}
