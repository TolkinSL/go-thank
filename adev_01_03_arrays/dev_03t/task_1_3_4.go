//Запуск Stepik заданий
//echo "Eyjafjallajokull 6" | go run task_1_3_4.go
//echo "12345 5" | go run task_1_3_4.go

package main

import "fmt"

func main() {
	var text string
	var width int
	var myStr string

	fmt.Scan(&text, &width)
	fmt.Printf("%#v\n", text)
	fmt.Printf("%#v\n", width)

	if len([]rune(text)) > width {
		r := []rune(text)
		myStr = string(r[:width]) + "..."
	} else {
		myStr = text
	}

	fmt.Printf("%#v\n", myStr)
}