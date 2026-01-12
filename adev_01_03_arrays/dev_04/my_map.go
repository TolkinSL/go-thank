package main

import "fmt"

func main() {
	m1 := make(map[string]int)
	m1["az"] = 3
	m1["cz"] = 7

	fmt.Printf("%#v\n", m1)
	fmt.Println(m1)
	fmt.Println(m1["cz"])
	fmt.Println(m1["cx"])

	for i, v := range m1 {
		fmt.Println(i, v)
	}

	m1["sd"] = 2
	fmt.Printf("%#v\n", m1)
	delete(m1, "cz")
	fmt.Printf("%#v\n", m1)

	str := "az"
	fmt.Println("m1[str]:", m1[str])

	val1, ok1 := m1["xd"]
	fmt.Println("m1[\"xd\"]:", val1, ok1)

	val1, ok1 = m1["sd"]
	fmt.Println("m1[\"sd\"]:", val1, ok1)

	my_map := map[string]int{"ab": 3, "zd": 9}
	fmt.Printf("%#v\n", my_map)

}