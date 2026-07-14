package main

import "fmt"

func main() {
	s := "Привет Go"

	fmt.Println(hasPrefix(s, "Прив"))
	fmt.Println(hasPrefix(s, "Прие"))

	fmt.Println()
	fmt.Println(contain(s, "вет"))

}

func hasPrefix(s, prefix string) bool {
	fmt.Printf("st: %s\npr: %s\n", s, prefix)
	return len(s) >= len(prefix) && s[0:len(prefix)] == prefix
}

func contain(s, substr string) bool {
	for i := 0; i < len(s) - len(substr); i++ {
		if hasPrefix(s[i:], substr) {
			return	true
		}		
	}

	return false
}