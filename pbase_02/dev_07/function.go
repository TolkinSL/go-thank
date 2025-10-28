package main

import "fmt"

func main() {
	test := []int{1, 2, 3, 4}
	fmt.Println("Sum -", sum(test...))

	printer := func(s string) {
		fmt.Println(s)
	}
	printer("Test func")

	type myStringFunc func(string)

	// Callback
	worker := func(callback myStringFunc) {
		callback("Выполнение callback")
	}

	worker(printer)

	// Замыкание
	logConnect := func(srv string) myStringFunc {
		return func(msg string) {
			fmt.Println("Server -", srv, "| message -", msg)
		}
	}

	srvMoscow := logConnect("Moscow")
	srvMoscow("Full success")

	func(s string) {
		fmt.Println("Анонимная функция и ее вызов",s)
	}("test")

	panicTest()
}

func sum(x ...int) int {
	var mySum int
	for _, v := range x {
		mySum += v
	}

	return mySum
}

func panicTest() {
	// Отлавливание паники для завершения работы программы
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Какое-то выполнение действий перед паникой -", err)
		}
	}()

	fmt.Println("Выполнение функции с паникой")
	panic("Ошибка в выполнении функции, Крах программы")
}