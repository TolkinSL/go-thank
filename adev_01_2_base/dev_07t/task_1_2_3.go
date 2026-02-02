//Запуск Stepik заданий
// $ echo "a 5" | go run ./task_1_2_3.go

package main

import (
	"fmt"
)

func main() {
	var source string
	var times int
	// гарантируется, что значения корректные
	fmt.Scan(&source, &times)

	// возьмите строку `source` и повторите ее `times` раз
	// запишите результат в `result`
  var result string

	for i := 0; i < times; i += 1 {
		result += source
	}

	fmt.Println(result)
}