package main

import (
	"fmt"
)

func main() {
	test := []int{1, 3, 2}
	reverse(test, 2, len(test))
	fmt.Println(test)
}

func nextPermutation(nums []int) {
	i := len(nums) - 2
	for ; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			break
		}
	}

	if i < 0 {
		reverse(nums, 0, len(nums))
		return
	}

	j := len(nums) - 1
	for ; j >= 0; j-- {
		if nums[i] < nums[j] {
			break
		}
	}

	nums[i], nums[j] = nums[j], nums[i]
	reverse(nums, i+1, len(nums))
}

func reverse(list []int, s, e int) {
	for i := s; i < s+(e-s)/2; i++ {
		list[i], list[e-1-(i-s)] = list[e-1-(i-s)], list[i]
	}
}
