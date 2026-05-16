package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	count := make(map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		fmt.Println("not have files")
		countWords(os.Stdin, count)
	} else {
		for _, file := range files {
			f, err := os.Open(file)
			if err != nil {
				fmt.Println(err)
				continue
			}
			countWords(f, count)
			f.Close()
		}
	}

	fmt.Println(files)
	fmt.Println(count)
}

func countWords(f *os.File, countArr map[string]int) {
	input := bufio.NewScanner(f)

	for input.Scan() {
		countArr[input.Text()]++
	}
}