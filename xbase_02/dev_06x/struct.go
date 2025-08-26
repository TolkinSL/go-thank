package main

import "fmt"

type Human struct {
	name string
	age int
}

func (a Human) isOld(n string) bool {
	fmt.Println("Hello -",n)
	if a.age > 18 {
		return true
	} 

	return false
}

func main() {
	myTest1 := Human{
		age: 19,
		name: "Vasia",
	}

	myTest2 := Human{
		age: 16,
		name: "Olia",
	}

	isOld1 := myTest1.isOld(myTest1.name)
	fmt.Println(isOld1)

	isOld2 := myTest2.isOld(myTest2.name)
	fmt.Println(isOld2)
}