package main

import "fmt"

func main() {
	fmt.Println(findNumberOfLIS([]int{1, 3, 5, 4, 7}))
}

func findNumberOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	cdp := make([]int, len(nums))

	var r int
	maxLen := 0
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		cdp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
					cdp[i] = cdp[j]
				} else if dp[j]+1 == dp[i] {
					cdp[i] += cdp[j]
				}
			}
		}

		if dp[i] > maxLen {
			maxLen = dp[i]
			r = cdp[i]
		} else if dp[i] == maxLen {
			r += cdp[i]
		}
	}

	return r
}
