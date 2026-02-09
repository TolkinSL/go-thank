package main

import (
	"fmt"
	"strings"
)

type lookupError struct {
	str string
	sub string
}

func (e lookupError) Error() string {
	return fmt.Sprintf("%s not found in: %s", e.sub, e.str)
}

func indexOf(str string, sub string) (int, error) {
	idx := strings.Index(str, sub)
	if idx == -1 {
		return -1, lookupError{str: str, sub: sub}
	}

	return idx, nil
}

func main() {
	str := "go is awesome"
	
	sub := "c++"
	res, err := indexOf(str, sub)
	fmt.Println(sub, res, err)
	fmt.Printf("%#v\n\n", err)

	sub = "awes"
	res, err = indexOf(str, sub)
	fmt.Println(sub, res, err)
	fmt.Printf("%#v\n\n", err)

	sub = "js"
	res, err = indexOf(str, sub)
	fmt.Println(sub, ":", res, err)
	fmt.Printf("%#v\n\n", err)

	if err, ok := err.(lookupError); ok {
		fmt.Println("Inside Error")
		fmt.Println(err.str)
		fmt.Println(err.sub)
	}
}