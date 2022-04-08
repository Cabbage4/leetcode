package main

import "fmt"

func main() {
	fmt.Println(maximumDifference([]int{1, 5, 2, 10}))
	//fmt.Println(maximumDifference([]int{9, 4, 3, 2}))
	//fmt.Println(maximumDifference([]int{7, 1, 5, 4}))
}

func maximumDifference(nums []int) int {
	r := -1

	last := nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		if last <= nums[i] {
			last = nums[i]
		} else if last-nums[i] > r {
			r = last - nums[i]
		}
	}

	return r
}
