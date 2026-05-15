package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)

	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		if input.Text() != "" {
			counts[input.Text()]++
		}
	}

	for item, val := range counts {
		if val > 1 {
			fmt.Printf("%s:\t%d\n", item, val)
		}
	}
}
