package main

import (
	// "errors"
	"errors"
	"fmt"
	"os"
	"strconv"
	// "strconv"
)

func check(a, b int) error {
	if a == 0 && b == 0 {
		return errors.New("a and b = 0")
	}
	return nil
}

func formattedCheck(a, b int) error {
	if a == 0 && b == 0 {
		return fmt.Errorf("a = %d and b = %d, UserID = %d\n", a, b, os.Getuid())
	}
	return nil
}

func main() {
	fmt.Println("")

	err1 := formattedCheck(0, 1)
	if err1 != nil {
		fmt.Printf("%+v\n", err1)
	} else {
		fmt.Println("Ok")
	}

	err2 := formattedCheck(0, 0)
	if err2 != nil {
		fmt.Printf("%#v\n", err2)
		fmt.Println(err2)
	} else {
		fmt.Println("Ok")
	}

	err3 := check(0, 0)
	if err3 != nil {
		fmt.Printf("%#v\n", err3)
		fmt.Printf("%#v\n", err3.Error())
	} else {
		fmt.Println("Ok")
	}

	_, err4 := strconv.Atoi("A123")
	if err4 != nil {
		fmt.Printf("%#v\n", err4)
		fmt.Println(err4)
	}
}
