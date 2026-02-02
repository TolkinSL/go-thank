package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	phrase := readString()
	strArr := strings.Fields(phrase)
	var abbr []rune

	for _, s := range strArr {
		myRune := []rune(strings.ToUpper(s))
		if unicode.IsLetter(myRune[0]) {
			abbr = append(abbr, myRune[0])
		}
	}

	fmt.Println(string(abbr))
}

func readString() string {
	rdr := bufio.NewReader(os.Stdin)
	str, _ := rdr.ReadString('\n')

	return str
}
