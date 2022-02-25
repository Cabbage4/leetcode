package main

import (
	"fmt"
)

func main() {
	fmt.Println(reverseStr("abcdefg", 2))
	fmt.Println(reverseStr("abcd", 2))
	fmt.Println(reverseStr("abcdefg", 1))
}

func reverseStr(s string, k int) string {
	r := make([]uint8, 0)
	i := 0
	for ; i+2*k < len(s); i = i + 2*k {
		for j := i + k - 1; j >= i; j-- {
			r = append(r, s[j])
		}
		for j := i + k; j < i+2*k; j++ {
			r = append(r, s[j])
		}
	}

	limit := len(s) - 1
	if i+k-1 < len(s)-1 {
		limit = i + k - 1
	}

	for j := limit; j >= i; j-- {
		r = append(r, s[j])
	}

	for j := limit + 1; j < len(s); j++ {
		r = append(r, s[j])
	}

	return string(r)
}
