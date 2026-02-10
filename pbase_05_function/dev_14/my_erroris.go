package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	a, b := 10, 0
	res, err := calculate(a, b)
	if err != nil {
		log.Fatalf("Unable calculate: %s", err)
	}

	fmt.Printf("Result calculate: %v\n", res)
	errors.New("Simpe run")
}

func calculate(a, b int) (int, error) {
	if err := logicX(); err != nil {
		return 0, fmt.Errorf("Error Calculate in LogicX: %w", err)
	}

	result, err := divide(a, b)
	if err != nil {
		return 0, fmt.Errorf("Error Calculate in Divide: %w", err)
	}

	return result, nil

}

func logicX() error {
	if err := logicY(); err != nil {
		return err
	}

	// return errors.New("failed by X")
	return nil
}

func logicY() error {
	// return errors.New("failed by YB")
	// return errors.New("failed by YC")

	return nil
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("Divide by 0!")
	}

	return a / b, nil
}
