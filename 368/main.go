package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(largestDivisibleSubset([]int{3, 4, 16, 8}))
	fmt.Println(largestDivisibleSubset([]int{5, 9, 18, 54, 108, 540, 90, 180, 360, 720}))
	fmt.Println(largestDivisibleSubset([]int{4, 8, 10, 240}))
	fmt.Println(largestDivisibleSubset([]int{2, 3}))
	fmt.Println(largestDivisibleSubset([]int{1, 2, 4, 7}))
	fmt.Println(largestDivisibleSubset([]int{1, 2, 3}))
}

func largestDivisibleSubset(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	sort.Ints(nums)

	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}

	mv := 0
	mc := 1
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i]%nums[j] == 0 {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}

		if dp[i] > mc {
			mc = dp[i]
			mv = nums[i]
		}
	}

	r := make([]int, 0)

	for i := len(nums) - 1; i >= 0 && mc > 0; i-- {
		if dp[i] == mc && mv%nums[i] == 0 {
			r = append(r, nums[i])
			mv = nums[i]
			mc--
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
