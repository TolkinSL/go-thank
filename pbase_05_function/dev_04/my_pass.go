package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	password := "Sdew321da"
	
	isSecured := securedPass(password)
	fmt.Println("Secured:", isSecured)
}

func securedPass(pass string) bool {
	
	if utf8.RuneCountInString(pass) < 8 {
		return false
	}

	if pass == strings.ToLower(pass) {
		return false
	}

	if pass == strings.ToUpper(pass) {
		return false
	}

	return true
}