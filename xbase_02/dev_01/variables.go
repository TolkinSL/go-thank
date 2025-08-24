package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x int
	fmt.Println(x)

	var a, b int = 5, 10
	fmt.Println(a, b)

	const (
		t = "azz"
		r = 12
		p // Сохраняется предыдущее если не указано 12
	)
	fmt.Println(t, r, p)

	fmt.Println(reflect.TypeOf(a), a) //Определить тип Переменной
}