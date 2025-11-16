package main

import "fmt"

type MyStruct struct {
	count int
}

func applyCount(s *MyStruct) {
	s.count += 10
}

func main() {
	var myCount MyStruct
	myCount.count = 3
	var pointer *MyStruct = &myCount
	fmt.Println(pointer)
	fmt.Println(pointer.count)
	// fmt.Println(*pointer.count) Ошибка нету указателяв свойстве count !
	fmt.Println((*pointer).count)
	pointer.count = 4
	fmt.Println(myCount)

	applyCount(&myCount)
	fmt.Println(myCount)

	applyCount(pointer)
	fmt.Println(myCount)

}