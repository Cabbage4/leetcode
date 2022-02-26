package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maximumProduct([]int{-100, -98, -1, 2, 3, 4}))
}

func maximumProduct(nums []int) int {
	ln := len(nums)
	sort.Ints(nums)

	if nums[ln-1] <= 0 || nums[0] >= 0 {
		return nums[ln-1] * nums[ln-2] * nums[ln-3]
	}

	one := nums[0] * nums[1] * nums[ln-1]
	two := nums[ln-1] * nums[ln-2] * nums[ln-3]

	if two > one {
		return two
	}

	return one
}
