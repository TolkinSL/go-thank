package main

import "fmt"

func main() {
	const testSt string = "test1"
	fmt.Printf("%#v - %T\n", testSt, testSt) //Вывод внутреннего полного Го значения
	
	const testSu = "test2"
	fmt.Printf("%v - %T\n", testSu, testSu)
	
	const ch = 't'
	fmt.Printf("%v - %T\n", ch, ch)
}