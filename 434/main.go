package main

import "fmt"

func main() {
	fmt.Println(countSegments("Hello, my name is John"))
	fmt.Println(countSegments("hello"))
}

func countSegments(s string) int {
	result := 0
	i := 0
	for i < len(s) {
		for i < len(s) && s[i] == ' ' {
			i++
		}

		if i >= len(s) {
			break
		}

		result++

		for i < len(s) && s[i] != ' ' {
			i++
		}
	}

	return result
}
