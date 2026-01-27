package main

import (
	"fmt"
	"runtime/debug"
)

func main() {
	fmt.Println("Main run")
	fn1()
	fmt.Println("Main end")
}

func fn1() {
	fn2()
	fmt.Println("Fn1")
}

func fn2() {
	fmt.Println("Fn2")
	debug.PrintStack()
}
