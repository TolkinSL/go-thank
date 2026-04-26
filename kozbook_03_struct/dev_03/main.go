package main

import (
	"fmt"
	"reflect"
)

type Animal struct {
	name string `help:"the name cat or dog"`
}

func main() {
	a := Animal{name: "Fish"}
	field, _ := reflect.TypeOf(a).FieldByName("name")
	fmt.Printf("%#v\n", field)
	fmt.Println(field.Tag.Get("help"))
	fmt.Println(field.Name)
}
