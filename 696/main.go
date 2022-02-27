package main

import "fmt"

func main() {
	fmt.Println(countBinarySubstrings("00110011"))
	fmt.Println(countBinarySubstrings("10101"))
}

func countBinarySubstrings(s string) int {
	if len(s) <= 1 {
		return 0
	}

	result := 0
	cur := 1
	pre := 0
	for i := 1; i < len(s); i++ {
		if s[i-1] != s[i] {
			result += min(cur, pre)
			pre = cur
			cur = 1
		} else {
			cur++
		}
	}

	result += min(cur, pre)

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
