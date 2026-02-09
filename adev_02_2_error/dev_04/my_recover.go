package main

import "fmt"

func getChar(str string, idx int) byte {
	return str[idx]
}

func main() {

	// recover() останвливает ту функцию где был вызван !

	safeCall := func() {
		defer func() {
			if err := recover(); err != nil {
			fmt.Println("safeCall Recover ERROR:", err)
			}	 else {
			fmt.Println("All Ok")
			}
		}()

		_ = getChar("Hi", 2)

		// не выполнится после recover()
		fmt.Println("safeCall complete")
	}
	
	safeCall()

	fmt.Println("Main work")

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recover ERROR:", err)
		} else {
			fmt.Println("All Ok")
		}
	}()

	_ = getChar("Hi", 2)

	fmt.Println("Main complete")
}
