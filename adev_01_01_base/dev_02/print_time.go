package main

import (
	"fmt"
	"time"
)

func main() {
	s := "1h45m"

	d, _ := time.ParseDuration(s)
	sm := fmt.Sprintf("%.0f", d.Minutes())
	fmt.Println("Конвертация времени в минуты")
	fmt.Println(s,"=",sm,"min")
	fmt.Println(s,"=",sm,"min")
	fmt.Printf("%T\n",d)
}