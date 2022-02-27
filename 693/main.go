package main

import (
	"fmt"
)

func main() {
	fmt.Println(hasAlternatingBits(5))
	fmt.Println(hasAlternatingBits(7))
	fmt.Println(hasAlternatingBits(11))
}

func hasAlternatingBits(n int) bool {
	last := n & 1
	n = n >> 1
	for n > 0 {
		cur := n & 1
		if cur == last {
			return false
		}
		last = cur
		n = n >> 1
	}

	return true
}
