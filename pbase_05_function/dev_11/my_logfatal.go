package main

import (
	"fmt"
	"log"
)

func main() {
	dbConnect()
	fmt.Println("Main end!")
}

func dbConnect() {
	fmt.Println("Db start!")
	defer func() {
		fmt.Println("Close db connection!")
	}()

	err := fmt.Errorf("Some Wrong!")
	if err != nil {
		log.Fatal(err) //os.Exit(1)
	}

	fmt.Println("Db close!")
}
