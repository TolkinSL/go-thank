package main

import (
	"errors"
	"fmt"
)

var errNotFound error = errors.New("not found")
type languages map[string]string

func getValue(m languages, key string) (string, error) {
	val, ok := m[key]
	if !ok {
		return "", errNotFound
	}

	return  val, nil
}

func (l languages) describe(lang string) (string, error) {
	res, err := getValue(l, lang)
	if err != nil {
		// return "", err
		
		// Возврат оригинальной ошибки через %w
		return "", fmt.Errorf("Error desc() %s: %w", lang, err)
	}

	return res, nil
}


func main() {
	langs := languages{
		"go": "number 1",
		"python": "number 2",
		"javascript": "number 3",
	}

	res, err := getValue(langs, "go")
	fmt.Println("getV", res, err)

	res, err = getValue(langs, "java")
	fmt.Println("getV", res, err)

	res, err = langs.describe("java")
	fmt.Println("langs.desc", res, err)
}
