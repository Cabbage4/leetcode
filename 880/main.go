package main

import "fmt"

func main() {
	fmt.Print(decodeAtIndex("ha22", 5))
}

func decodeAtIndex(s string, k int) string {
	size := 0
	for i := 0; i < len(s); i++ {
		if '0' <= s[i] && s[i] <= '9' {
			size *= int(s[i] - '0')
		} else {
			size++
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		k = k % size
		if k == 0 && 'a' <= s[i] && s[i] <= 'z' {
			return string(s[i])
		}

		if '0' <= s[i] && s[i] <= '9' {
			size = size / int(s[i]-'0')
		} else {
			size--
		}
	}

	return ""
}
