package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	str := "Привет!\u00A9\n"
	fmt.Println(str)
	fmt.Println("Байт", len(str))
	fmt.Println("Символов", utf8.RuneCountInString(str))
	fmt.Println("Символов в байтах", utf8.RuneCount([]byte(str)))

	fmt.Println("----------")
	str2 := "\u00A9"
	fmt.Println(str2)
	fmt.Println("Байт", len(str2))
	fmt.Println("Символов", utf8.RuneCountInString(str2))
	fmt.Println("Символов в байтах", utf8.RuneCount([]byte(str2)))

	RuneError := '\uFFFD'
	fmt.Println(string(RuneError))
	
	fmt.Println("----------")
	str3 := "   Привет!\u00A9  "
	fmt.Println(str3)
	fmt.Println(strings.TrimSpace(str3))
	fmt.Println(strings.ToUpper(str3))
	fmt.Println(strings.ToLower(str3))

	fmt.Println("----------")
	str4 := "Hello Go in 2025"
	fmt.Println(strings.Split(str4, " "))
	fmt.Println("Строка содержит Go", strings.Contains(str4, "Go"))
	fmt.Println("Строка содержит go", strings.Contains(str4, "go"))

	fmt.Println(strings.HasPrefix(str4, "He"))

	str5 := "HeLLo Go IN 2025"
	fmt.Println("Сравнение строк без учета регистра", strings.EqualFold(str4, str5))
}