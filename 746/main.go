package main

import "fmt"

func main() {
	fmt.Println(minCostClimbingStairs([]int{10, 15, 20}))
	fmt.Println(minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}))
}

func minCostClimbingStairs(cost []int) int {
	dp := make([]int, len(cost))
	dp[0] = 0
	dp[1] = 0

	for i := 2; i < len(cost); i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}

	ln := len(cost)
	return min(dp[ln-1]+cost[ln-1], dp[ln-2]+cost[ln-2])
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
