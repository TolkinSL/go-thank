package main

import "fmt"

func main() {
	var x int8 = 127
	x += 1
	fmt.Println(x) //-128

	var z uint8 = 255
	z += 1
	fmt.Println(z) //0

	var t = 200
	var p64 int64
	p64 = int64(t)
	fmt.Println(p64)

	var p8 int8
	p8 = int8(t)
	fmt.Println(p8) //-56
}