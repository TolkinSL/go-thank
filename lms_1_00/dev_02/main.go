package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	fmt.Println("Current time:", currentTime)

	workTime := time.Date(2025, 8, 17, 9, 30, 0, 0, time.UTC) 
	fmt.Println("Work time:", workTime)

	diffTime := currentTime.Sub(workTime)
	fmt.Println("Diff time hours:", diffTime.Hours())

	fmt.Println("Format time:", workTime.Format("15:04:05 02-01-2006"))
}
