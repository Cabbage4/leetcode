package main

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{1, 2, 3, 0, 2}))
}

// 0->什么都没有，1->有股票，2->冷冻期中
func maxProfit(prices []int) int {
	dp := make([][3]int, len(prices))
	dp[0][1] = -prices[0]
	var r int
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][2])
		dp[i][1] = max(dp[i-1][0]-prices[i], dp[i-1][1])
		dp[i][2] = dp[i-1][1] + prices[i]

		r = max(r, max(dp[i][0], max(dp[i][1], dp[i][2])))
	}
	return r
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
