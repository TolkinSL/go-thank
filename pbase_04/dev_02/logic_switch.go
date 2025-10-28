package main

import "fmt"

func main() {
	day := 3

	switch day {
	case 1:
		fmt.Println("Mon")
	case 2:
		fmt.Println("Tue")
	case 3:
		fmt.Println("Wed")
		fmt.Println("Test")
		fallthrough //Проваливание на следующий case
	case 4:
		fmt.Println("Thu")
	case 5:
		fmt.Println("Fri")
	default:
		fmt.Println("Unknown day")
	}

	switch {
	case day <= 3:
		fmt.Println("Первая половина недели")
	case day <= 5:
		fmt.Println("Вторая половина недели")
	default:
		fmt.Println("Неизвестно что")
	}

	switch day {
	case 1, 2, 3, 4, 5:
		fmt.Println("Рабочий день")
		//fallthrough
	case 6, 7:
		fmt.Println("Выходной день")
	default:
		fmt.Println("Неизвестный день")
	}

	var x any = "zzz "

	switch s := x.(type) {
	case int:
		fmt.Println("Type int -", s)
	case string:
		fmt.Println("Type string -", s)
	default:
		fmt.Println("Unknown Type -", s)
	}
}
