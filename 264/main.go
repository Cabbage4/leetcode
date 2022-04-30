package main

import (
	"fmt"
)

func main() {
	fmt.Println(nthUglyNumber(203))
	fmt.Println(nthUglyNumber(103))
	fmt.Println(nthUglyNumber(10))
}

func nthUglyNumber(n int) int {
	if n == 1 {
		return 1
	}
	ct2, ct3, ct5 := 0, 0, 0
	dp := make([]int, n)
	dp[0] = 1
	for i := 1; i < n; i++ {
		dp[i] = min(min(dp[ct2]*2, dp[ct3]*3), dp[ct5]*5)
		if dp[i] == dp[ct2]*2 {
			ct2++
		}
		if dp[i] == dp[ct3]*3 {
			ct3++
		}
		if dp[i] == dp[ct5]*5 {
			ct5++
		}
	}
	return dp[n-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
