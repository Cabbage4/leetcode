package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(largestSumAfterKNegations([]int{4, 2, 3}, 1))
	fmt.Println(largestSumAfterKNegations([]int{3, -1, 0, 2}, 3))
	fmt.Println(largestSumAfterKNegations([]int{2, -3, -1, 5, -4}, 2))
	fmt.Println(largestSumAfterKNegations([]int{-8, 3, -5, -3, -5, -2}, 6))
}

func largestSumAfterKNegations(nums []int, k int) int {
	r := 0
	sort.Slice(nums, func(i, j int) bool {
		return abs(nums[i]) < abs(nums[j])
	})

	for i := len(nums) - 1; i >= 0; i-- {
		if k <= 0 {
			r += nums[i]
			continue
		}

		if nums[i] < 0 {
			r -= nums[i]
			k--
			continue
		}

		r += nums[i]
	}

	if k > 0 && k%2 != 0 {
		r -= abs(nums[0]) * 2
	}

	return r
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
