package main

import "fmt"

func main() {
	fmt.Println(countHillValley([]int{57, 57, 57, 57, 57, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 85, 85, 85, 86, 86, 86}))
	fmt.Println(countHillValley([]int{2, 4, 1, 1, 6, 5}))
}

func countHillValley(nums []int) int {
	r := 0
	j := 0
	for i := 1; i < len(nums)-1; i++ {
		if (nums[j] < nums[i] && nums[i] > nums[i+1]) ||
			(nums[j] > nums[i] && nums[i] < nums[i+1]) {
			j = i
			r++
		}
	}
	return r
}
