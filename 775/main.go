package main

import (
	"fmt"
)

func main() {
	fmt.Println(isIdealPermutation([]int{1, 0, 2}))
}

func isIdealPermutation(nums []int) bool {
	minValue := len(nums)
	for i := len(nums) - 1; i >= 2; i-- {
		minValue = min(minValue, nums[i])
		if nums[i-2] > minValue {
			return false
		}
	}

	return true
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
