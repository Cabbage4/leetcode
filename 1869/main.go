package main

import "fmt"

func main() {
	fmt.Println(checkZeroOnes("1101"))
	fmt.Println(checkZeroOnes("10"))
	fmt.Println(checkZeroOnes("1"))
}

func checkZeroOnes(s string) bool {
	l0 := 0
	l1 := 0

	c0 := 0
	c1 := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			c0++
			c1 = 0
			l0 = max(l0, c0)
		} else {
			c1++
			c0 = 0
			l1 = max(l1, c1)
		}
	}

	return l1 > l0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
