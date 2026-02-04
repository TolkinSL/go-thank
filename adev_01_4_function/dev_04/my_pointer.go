package main

import (
	"fmt"
	"sort"
)

func addptr(xp *int, y int) {
	*xp += y
}

func sortSlice(nums []int) {
	sort.Ints(nums)
}

func appendVal(nums []int, x int) {
	fmt.Printf("Func &nums[0] - %p\n", &nums[0])
	nums = append(nums, x)
	fmt.Printf("Func append &nums[0] - %p\n", &nums[0])
}

func appendValPtr(nums *[]int, x int) {
	fmt.Printf("Func appendPtr *nums - %p\n", &(*nums)[0])
	*nums = append(*nums, x)
	fmt.Printf("Func appendPtr val nums[0] - %v\n", (*nums)[0])
	fmt.Printf("Func appendPtr ptr nums[0] - %p\n", &(*nums)[0])
}

func appendReturn(nums []int, x int) []int {
	nums = append(nums, x)
	return nums
}

func mapChange(myMap map[string]int) {
	myMap["c"] = 7
	myMap["d"] = 5
}

func mapReturn(myMap map[string]int) map[string]int {
	m := myMap
	delete(m, "a")
	delete(m, "c")
	return m
}

func main() {
	var ip *int
	fmt.Println("ip---")
	fmt.Printf("%#v\n", ip)
	// fmt.Printf("%#v\n", *ip) //panic: runtime error: invalid memory address

	num := 37
	ip = &num
	fmt.Printf("%#v\n", ip)
	fmt.Printf("%#v\n", *ip)

	fmt.Println("num---")
	addptr(&num, 5)
	fmt.Printf("%#v\n", num)
	fmt.Printf("%#v\n", ip)
	fmt.Printf("%#v\n", *ip)

	// Slice in function
	fmt.Println("Slice---")
	myArr1 := []int{6, 3, 1, 5, 9}
	sortSlice(myArr1)
	fmt.Printf("%#v\n", myArr1)

	fmt.Println("Slice---val")
	fmt.Printf("Main &myArr1[0] - %p\n", &myArr1[0])
	appendVal(myArr1, 8)
	fmt.Printf("Main %#v\n", myArr1)

	fmt.Println("Slice---ptr")
	fmt.Printf("Main &myArr1[0] - %p\n", &myArr1[0])
	appendValPtr(&myArr1, 8)
	fmt.Printf("Main %#v\n", myArr1)

	fmt.Println("Slice---return")
	fmt.Printf("Main &myArr1[0] - %p\n", &myArr1[0])
	fmt.Printf("Main %#v\n", myArr1)
	myArr1 = appendReturn(myArr1, 12)
	fmt.Printf("Main &myArr1[0] - %p\n", &myArr1[0])
	fmt.Printf("Main %#v\n", myArr1)

	fmt.Println("Map---change")
	myMap := map[string]int{
		"a": 2,
		"b": 4,
	}

	fmt.Printf("Main %#v\n", myMap)
	mapChange(myMap)
	fmt.Printf("Main %#v\n", myMap)
	myMap = mapReturn(myMap)
	fmt.Printf("Main %#v\n", myMap)

}