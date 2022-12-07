package main

func main() {
}

func maxProfit(prices []int) int {
	dp := make([][2]int, len(prices))
	dp[0][1] = -prices[0]
	var r int
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
		r = max(r, max(dp[i][0], dp[i][1]))
	}
	return r
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
