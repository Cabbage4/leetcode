package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxProfit(2, []int{3, 2, 6, 5, 0, 3}))
}

func maxProfit(k int, prices []int) int {
	hasDP := make([][]int, len(prices))
	noHasDP := make([][]int, len(prices))
	for i := range hasDP {
		hasDP[i] = make([]int, k+1)
		noHasDP[i] = make([]int, k+1)
	}

	hasDP[0][0] = -prices[0]
	for i := 1; i <= k; i++ {
		hasDP[0][i] = math.MinInt
	}

	var r int
	for i := 1; i < len(prices); i++ {
		hasDP[i][0] = max(hasDP[i-1][0], noHasDP[i-1][0]-prices[i])

		r = max(r, noHasDP[i][0])
		for j := 1; j <= k; j++ {
			noHasDP[i][j] = max(noHasDP[i-1][j], hasDP[i-1][j-1]+prices[i])
			hasDP[i][j] = max(hasDP[i-1][j], noHasDP[i-1][j]-prices[i])

			r = max(r, noHasDP[i][j])
		}
	}

	return r
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
