package main

import (
	"fmt"
)

func main() {
	fmt.Println(minFlipsMonoIncr("100000001010000"))
	//fmt.Println(minFlipsMonoIncr("00110"))
	//fmt.Println(minFlipsMonoIncr("010110"))
}

func minFlipsMonoIncr(s string) int {
	preSum := make([]int, len(s))
	if s[0] == '1' {
		preSum[0] = 1
	}

	for i := 1; i < len(s); i++ {
		preSum[i] = preSum[i-1]
		if s[i] == '1' {
			preSum[i]++
		}
	}

	r := len(s) - preSum[len(s)-1]
	for i := 0; i < len(s); i++ {
		tmp := 2*preSum[i] + len(s) - i - 1 - preSum[len(s)-1]
		r = min(r, tmp)
	}

	return r
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}
