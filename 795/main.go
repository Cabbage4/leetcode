package main

import (
	"fmt"
)

func main() {
	fmt.Println(numSubarrayBoundedMax([]int{2, 1, 4, 3}, 2, 3))
	fmt.Println(numSubarrayBoundedMax([]int{2, 9, 2, 5, 6}, 2, 8))
}

func numSubarrayBoundedMax(nums []int, left int, right int) int {
	if len(nums) == 0 {
		return 0
	}

	count := func(target int) int {
		var r int
		c := 0
		for _, v := range nums {
			if v <= target {
				c++
			} else {
				c = 0
			}
			r += c
		}
		return r
	}

	return count(right) - count(left-1)
}
