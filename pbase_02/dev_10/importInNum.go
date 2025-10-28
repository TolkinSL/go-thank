package main

import (
	"fmt"
	"strconv"
)

func main() {
	str1 := "43765"

	num1, err := strconv.Atoi(str1)
	if err != nil {
		fmt.Println(str1, "Ошибка преобразования в число int", err)
	}
	fmt.Println("int Atoi", num1)

	num2, err := strconv.ParseInt(str1, 10, 64)
	if err != nil {
		fmt.Println(str1, "Ошибка преобразования в число int", err)
	}
	fmt.Println("int ParseInt", num2)

	str2 := "0101"
	num3, err := strconv.ParseInt(str2, 2, 64)
	if err != nil {
		fmt.Println(str1, "Ошибка преобразования в число int", err)
	}
	fmt.Println("int8(byte) ParseInt", num3)

	str4 := "3.14326"
	num4, err := strconv.ParseFloat(str4, 64)
	if err != nil {
		fmt.Println(str1, "Ошибка преобразования в число int", err)
	}
	fmt.Println("float64 ParseFloat", num4)
}