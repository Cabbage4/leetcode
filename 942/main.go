package main

import "fmt"

func main() {
	fmt.Println(diStringMatch("IDI"))
}

func diStringMatch(s string) []int {
	left := 0
	right := len(s)
	r := make([]int, len(s)+1)
	for i := 0; i < len(s); i++ {
		if s[i] == 'I' {
			r[i] = left
			left++
		} else {
			r[i] = right
			right--
		}
	}

	r[len(s)] = right

	return r
}
