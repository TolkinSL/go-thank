//Запуск Stepik заданий
// $ echo "Eyjafjallajokull 6" | go run ./dev_01_10t
// $ echo "12345 5" | go run ./dev_01_10t

package main

import (
	"fmt"
)

func main() {
	var text string
	var width int
	fmt.Scanf("%s %d", &text, &width)

	// Возьмите первые `width` байт строки `text`,
	// допишите в конце `...` и сохраните результат
	// в переменную `res`
	// ...
	// res := (make([]string, 0, width + 3))
	// res = append(res, text[: width] + "...")
	// fmt.Println(res[0])

	var res string
	if len(text) <= width {
		res = text
	} else {
		res = text[:width] + "..."
	}

	fmt.Println(res)
	// fmt.Printf("%#v\n", res)
}
