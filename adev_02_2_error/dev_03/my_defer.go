package main

import (
	"fmt"
	"os"
)

func createFile(name string) (*os.File, error) {
	fmt.Println("Creating file...")
	
	f, err := os.Create(name)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func closeFile(f *os.File) {
	fmt.Println("Closing file...")
	err := f.Close()
	if err != nil {
		fmt.Println("Error closing file:", err)
	}
}

func writeFile(f *os.File) error {
	fmt.Println("Writing to file...")

	if _, err := f.Write([]byte("Hello Go!")); err != nil {
		return err
	}

	//return nil
	return fmt.Errorf("Writing file error !")
}

func main() {
	// f, err := createFile("/my_temp.txt") // witch error
	f, err := createFile("./my_temp.txt")
	
	if err != nil {
		fmt.Println("Main Error creating file:", err)
		return
	}

	defer closeFile(f)

	if err := writeFile(f); err != nil {
		fmt.Println("Main Error Writing file:", err)
		return
	}

	fmt.Println("Success!")
}
