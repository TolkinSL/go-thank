package main

import "fmt"

func main() {
	s1 := make([]string, 3)
	fmt.Println(len(s1))
	fmt.Println(cap(s1))
	fmt.Printf("%#v\n", s1)

	s2 := []string{"a", "b", "c"}
	s2 = append(s2, "z")
	fmt.Printf("%#v\n", s2)

	dst := make([]string, len(s2))
	copy(dst, s2)
	fmt.Printf("%#v\n", s2)
	fmt.Printf("%#v\n", s2[1:3])
	fmt.Printf("%#v\n", s2[:2])
	fmt.Printf("%#v\n", s2[2:])

	var s3 []string
	fmt.Printf("%#v\n", s3)
	fmt.Println(len(s3))
	fmt.Println(cap(s3))
	s3 = append(s3, "abc")
	fmt.Printf("%#v\n", s3)
}
