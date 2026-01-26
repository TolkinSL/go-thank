package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	x := 10
	y := 30

	fmt.Println("x type:", reflect.TypeOf(x))
	fmt.Println("x bytes:", unsafe.Sizeof(x))

	defer func(val int) {
		fmt.Println("Start 1 x:", val) // 10
	}(x)

	defer func() {
		fmt.Println("Start 2 y:", y) // 300
	}()

	x = 100
	y = 300
	fmt.Println("Main end")
}
