package main

import "fmt"

func main() {
	fmt.Println(rangeBitwiseAnd(5, 7))
}

func rangeBitwiseAnd(left int, right int) int {
	if left == 0 || right == 0 {
		return 0
	}

	c := 0
	for left != right {
		left = left >> 1
		right = right >> 1
		c++
	}

	return left << c
}
