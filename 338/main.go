package main

import (
	"fmt"
)

func main() {
	r := countBits(30)
	for i := range r {
		fmt.Printf("%d: %d\n", i, r[i])
	}
}

func countBits(n int) []int {
	dp := make([]int, n+1)

	dp[0] = 0
	offset := 1
	for i := 1; i <= n; i++ {
		if i == offset*2 {
			offset *= 2
		}
		dp[i] = dp[i-offset] + 1
	}

	return dp
}
