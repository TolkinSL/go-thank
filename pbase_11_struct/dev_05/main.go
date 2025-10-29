package main

import (
	"develop/pbase_11_struct/dev_05/user"
	"fmt"
)

func main() {
	name := user.User{
		Name: user.Name{First: "Vasia"},
	}

	fmt.Println(name)
	name.ChangeYear(1940)
	fmt.Println(name)
	name.ChangeYear(1990)
	fmt.Println(name)

	// user.isGoodYear(1990)
}
