package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func main() {
	if num := rand.IntN(10); num < 5 {
		fmt.Println("Small number", num)
	} else {
		fmt.Println("Big number", num)
	}

	switch i := 2; i {
	case 1:
		fmt.Println("1-", i)
	case 2:
		fmt.Println("2-", i)
		fallthrough
	case 3:
		fmt.Println("3-", i)
	default:
		fmt.Println("Неизвестное число")
	}

	switch day := time.Now().Weekday(); day {
	case time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday:
		fmt.Println("Рабочие дни", day)
	case time.Saturday, time.Sunday:
		fmt.Println("Выходные дни", day)
	}

	myT := time.Now().Hour()
	switch {
	case myT > 22 || myT < 6:
		fmt.Println("Ночь", myT)
	default:
		fmt.Println("Рабочее время", myT)
	}
}