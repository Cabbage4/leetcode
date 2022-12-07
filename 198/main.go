package main

import "fmt"

func main() {
	fmt.Println(rob([]int{1, 2, 3, 1}))
	fmt.Println(rob([]int{2, 7, 9, 3, 1}))
}

func rob(nums []int) int {
	if len(nums) < 3 {
		var r int
		for i := 0; i < len(nums); i++ {
			r = max(r, nums[i])
		}
		return r
	}

	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(dp[0], nums[1])
	dp[2] = max(nums[1], nums[0]+nums[2])

	r := max(dp[0], max(dp[1], dp[2]))
	for i := 3; i < len(nums); i++ {
		dp[i] = max(nums[i]+dp[i-2], dp[i-1])
		r = max(r, dp[i])
	}
	return r
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
