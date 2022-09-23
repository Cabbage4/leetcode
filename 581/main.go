package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findUnsortedSubarray([]int{1, 2, 3, 4}))
	fmt.Println(findUnsortedSubarray([]int{2, 6, 4, 8, 10, 9, 15}))
}

func findUnsortedSubarray(nums []int) int {
	min := math.MaxInt
	max := math.MinInt
	left := -1
	right := -1
	for i := 0; i < len(nums); i++ {
		if min >= nums[len(nums)-i-1] {
			min = nums[len(nums)-i-1]
		} else {
			left = len(nums) - i - 1
		}

		if max <= nums[i] {
			max = nums[i]
		} else {
			right = i
		}
	}

	if right == -1 {
		return 0
	}

	return right - left + 1
}
