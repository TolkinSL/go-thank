// go run main.go -s "-" -n a1 b1 c1
// go run main.go -s ", " a1 b1 c1

package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "пропуск перехода на новую строку")
var s = flag.String("s", " ", "разделитель элементов")

func main() {
	flag.Parse()
	fmt.Printf("%#v\n", *n)
	fmt.Printf("%#v\n", *s)
	fmt.Printf("%#v\n", flag.Args())

	fmt.Print(strings.Join(flag.Args(), *s))
	if !*n {
		fmt.Print("\n")
	}
}
