package main

import (
	"fmt"
	"strconv"
)

func main() {
	i1 := 1977
	s1 := strconv.Itoa(i1)
	fmt.Println(s1)
	fmt.Println(strconv.FormatInt(int64(i1), 10))
	fmt.Println(strconv.FormatInt(int64(i1), 2))
	fmt.Println(strconv.FormatInt(int64(i1), 16))

	fmt.Println("----------")
	f2 := 10.3745
	s2 := strconv.FormatFloat(f2, 'f', 3, 64)
	fmt.Println(s2)

	s3 := fmt.Sprintf("Возврат строки %0.2f", f2)
	fmt.Println(s3)
	fmt.Printf("Переменная %#v\n", f2)
	fmt.Printf("Строка %#v\n", s3)
}