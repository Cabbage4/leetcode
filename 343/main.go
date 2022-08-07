package main

import (
	"fmt"
)

func main() {
	fmt.Println(integerBreak(27))
	fmt.Println(integerBreak(6))
	//fmt.Println(integerBreak(5))
	//fmt.Println(integerBreak(2))
	//fmt.Println(integerBreak(10))
}
func integerBreak(n int) int {
	if n == 3 {
		return 2
	}
	if n == 2 {
		return 1
	}

	c3 := n / 3
	if n%3 == 0 {
		return pow(c3)
	} else if n%3 == 1 {
		return pow(c3-1) * 2 * 2
	} else {
		return pow(c3) * 2
	}
}

func pow(n int) int {
	if n == 0 {
		return 1
	}

	if n == 1 {
		return 3
	}

	if n%2 == 0 {
		v := pow(n / 2)
		return v * v
	}
	v := pow(n / 2)
	return v * v * 3
}
