package main

import (
	"develop/hfbook_08_struct/dev_06/magazine"
	"fmt"
)


func main() {
	s := magazine.Subscriber{
		Active: true,
	}

	fmt.Printf("%+v", s)

	s.Active = false
	fmt.Printf("%+v", s)

	// s.name = "Test"
}
