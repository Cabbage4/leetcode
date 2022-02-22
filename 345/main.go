package main

import (
	"fmt"
)

func main() {
	fmt.Println(reverseVowels("hello"))
	fmt.Println(reverseVowels("leetcode"))
}

func reverseVowels(s string) string {
	mp := make([]bool, 256)
	mp['a'] = true
	mp['e'] = true
	mp['i'] = true
	mp['o'] = true
	mp['u'] = true
	mp['A'] = true
	mp['E'] = true
	mp['I'] = true
	mp['O'] = true
	mp['U'] = true

	b := []byte(s)

	left := 0
	right := len(b) - 1

	for left < right {
		for left < right && !mp[b[left]] {
			left++
		}
		for left < right && !mp[b[right]] {
			right--
		}

		b[left], b[right] = b[right], b[left]
		left++
		right--
	}

	return string(b)
}
