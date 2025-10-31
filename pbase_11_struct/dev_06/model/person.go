package model

import "errors"

type Person struct {
	First string
	Last string
	Age   int
}

func NewPerson(first string, last string, age int) (*Person, error) {
	if age <10 {
		return nil, errors.New("Fatal Age")
	}
	return &Person{
		First: first,
		Last: last,
		Age:   age,
	}, nil
}
