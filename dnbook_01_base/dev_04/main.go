package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	files := os.Args[1:]
	count := make(map[string]int)

	for _, f := range files {
		data, _ := os.ReadFile(f)

		strArr := strings.Split(string(data),"\n")
		for _, item := range strArr {
			count[item]++
		}
	}

	fmt.Println(count)
}
