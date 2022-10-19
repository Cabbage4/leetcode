package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findMaxK([]int{-1, 10, 6, 7, -7, 1}))
	fmt.Println(findMaxK([]int{-1, 2, -3, 3}))
	fmt.Println(findMaxK([]int{-10, 8, 6, 7, -2, -3}))
}

func findMaxK(nums []int) int {
	sort.Ints(nums)
	left := 0
	right := len(nums) - 1
	for left < right {
		if -nums[left] == nums[right] {
			return nums[right]
		}

		if nums[left] >= 0 {
			break
		}

		if -nums[left] > nums[right] {
			left++
		} else {
			right--
		}
	}

	return -1
}
