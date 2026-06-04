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
	// return errors.New("my error 111")
	return fmt.Errorf("func: %s, process id: %d", "myErrorFunc", 1023)
}

func div(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("divide by 0 111")
	}

	return x / y, nil
}
