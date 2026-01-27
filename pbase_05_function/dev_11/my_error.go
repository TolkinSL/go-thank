package main

import (
	"errors"
	"fmt"
)

func main() {
	err := someFunc()
	if err != nil {
		fmt.Printf("%#v\n", err)
	}

	myDiv, err := div(4, 2)
	fmt.Println("div:", myDiv, err)

	myDiv, err = div(4, 0)
	fmt.Println("div:", myDiv, err)
}

func someFunc() error {
	// return errors.New("My error 111")
	return fmt.Errorf("Func: %s, Process ID: %d", "myErrorFunc", 1023)
}

func div(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("Divide by 0 111")
	}

	return x / y, nil
}
