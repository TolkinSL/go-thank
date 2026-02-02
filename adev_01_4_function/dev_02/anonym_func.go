package main

import (
	"fmt"
	"sort"
)

func myInc() func() int {
	i := 0
	
	return func() int {
		i++

		return i
	}
}

func main() {
	
	inc1 := myInc()

	fmt.Println("inc1()")
	fmt.Println(inc1())
	fmt.Println(inc1())
	fmt.Println(inc1())

	inc2 := myInc()
	fmt.Println("inc2()")
	fmt.Println(inc2())

	arr := []int{2, 4, 6, 7, 9}
	xs := 5

	closest := sort.Search(len(arr), func(i int) bool {
		fmt.Println("i", i)
		if arr[i] >= xs {
			return true
		}

		return false
	})

	fmt.Println("closest", closest)
}
