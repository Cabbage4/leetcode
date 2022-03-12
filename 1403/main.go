package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minSubsequence([]int{4, 4, 7, 6, 7}))
}

func minSubsequence(nums []int) []int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	dp := make([]int, len(nums))
	sum := 0
	for i, v := range nums {
		sum += v
		dp[i] = sum
	}

	left := 0
	right := len(nums)
	for left < right {
		mid := (left + right) / 2
		if dp[mid]*2 <= sum {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return nums[0 : left+1]
}
