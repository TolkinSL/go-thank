package main

import (
	"errors"
	"fmt"
)

func main() {
	res, mod, err := div(17, 3)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res, mod)

	res, mod, err = div(17, 0)
	if err != nil {
		fmt.Println(err)
		fmt.Println(err.Error())
		return
	}
}

func div(num, denum int) (int, int, error) {
	if denum == 0 {
		myError := errors.New("divide by 0")

		fmt.Printf("%#v\n",myError)
		return 0, 0, fmt.Errorf("divide by 0")
	}

	return num / denum, num % denum, nil
}