package main

import (
	"fmt"
)

func main() {
	fmt.Println(findErrorNums([]int{1, 2, 2, 4}))
}

func findErrorNums(nums []int) []int {
	mp := make([]int, len(nums)+1)

	dup := -1
	for i := 0; i < len(nums); i++ {
		if mp[nums[i]] != 0 {
			dup = nums[i]
		}
		mp[nums[i]] = -1
	}

	loss := -1
	for i := 1; i <= len(nums); i++ {
		if mp[i] == 0 {
			loss = i
			break
		}
	}

	return []int{dup, loss}
}
