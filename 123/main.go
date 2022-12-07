package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxProfit([]int{3, 3, 5, 0, 0, 3, 1, 4}))
}

func maxProfit(prices []int) int {
	hasDP := make([][]int, len(prices))
	noHasDP := make([][]int, len(prices))

	for i := 0; i < len(prices); i++ {
		hasDP[i] = make([]int, 3)
		noHasDP[i] = make([]int, 3)
	}

	hasDP[0][0] = -prices[0]
	hasDP[0][1] = math.MinInt / 2
	hasDP[0][2] = math.MinInt / 2

	var r int
	for i := 1; i < len(prices); i++ {
		hasDP[i][0] = max(hasDP[i-1][0], noHasDP[i-1][0]-prices[i])
		for j := 1; j <= 2; j++ {
			hasDP[i][j] = max(hasDP[i-1][j], noHasDP[i-1][j]-prices[i])
			noHasDP[i][j] = max(noHasDP[i-1][j], hasDP[i-1][j-1]+prices[i])

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
