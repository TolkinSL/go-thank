package main

import "fmt"

type MyStruct struct {
	count int
}

func applyCount(s *MyStruct) {
	s.count += 10
}

func applyDiscount(s *MyStruct) {
	s.count -= 3
}

func (ms *MyStruct) applyDiscountMethod(num int) {
	ms.count -= num
}

func main() {
	var myCount MyStruct
	myCount.count = 3
	fmt.Printf("\n%+v\n", myCount)

	var pointer *MyStruct = &myCount
	fmt.Println(pointer)
	fmt.Println(pointer.count)
	// fmt.Println(*pointer.count) Ошибка нету указателяв свойстве count !
	fmt.Println((*pointer).count) // Получает значение структуры потом обращается с свойству
	pointer.count = 4
	fmt.Println(myCount)

	applyCount(&myCount)
	fmt.Println(myCount)

	applyCount(pointer)
	fmt.Println(myCount)

	applyDiscount(&myCount)
	fmt.Println("applyDiscount", myCount)

	myCount.applyDiscountMethod(2)
	fmt.Println("myCount.applyDiscountMethod(2)", myCount)

}