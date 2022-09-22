package main

import "fmt"

func main() {
	fmt.Println(findTargetSumWays([]int{1, 0}, 1))
	fmt.Println(findTargetSumWays([]int{1, 1, 1, 1, 1}, 3))
}

func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}

	if sum-target < 0 || (sum-target)%2 != 0 {
		return 0
	}

	neg := (sum - target) / 2

	dp := make([][]int, len(nums)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, neg+1)
	}

	dp[0][0] = 1
	for i := 0; i < len(nums); i++ {
		for j := 0; j <= neg; j++ {
			dp[i+1][j] = dp[i][j]
			if nums[i] <= j {
				dp[i+1][j] += dp[i][j-nums[i]]
			}
		}
	}

	return dp[len(nums)][neg]
}
