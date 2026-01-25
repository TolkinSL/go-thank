package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	strHello := "Привет  Go"
	myBytes, myRunes := calcSymbol(strHello)
	fmt.Println(myBytes, myRunes)
	
}

func calcSymbol(str string) (myBytes int, myRunes int) {
	myBytes = len(str)
	myRunes = utf8.RuneCountInString(str)

	return
}