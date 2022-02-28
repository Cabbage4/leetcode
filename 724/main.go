package main

import "fmt"

func main() {
	fmt.Println(pivotIndex([]int{1, 7, 3, 6, 5, 6}))
	fmt.Println(pivotIndex([]int{1, 2, 3}))
}

func pivotIndex(nums []int) int {
	left := 0
	right := 0
	for i := 1; i < len(nums); i++ {
		right += nums[i]
	}

	if left == right {
		return 0
	}

	for i := 1; i < len(nums); i++ {
		left += nums[i-1]
		right -= nums[i]
		if left == right {
			return i
		}
	}

	return -1
}
