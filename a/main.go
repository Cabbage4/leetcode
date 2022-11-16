package main

import "fmt"

func main() {
	fmt.Println(applyOperations([]int{0, 1}))
	fmt.Println(applyOperations([]int{1, 2, 2, 1, 1, 0}))
}

func applyOperations(nums []int) []int {
	var r []int
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			nums[i-1] *= 2
			nums[i] = 0
		}

		if nums[i-1] != 0 {
			r = append(r, nums[i-1])
		}
	}

	if nums[len(nums)-1] != 0 {
		r = append(r, nums[len(nums)-1])
	}

	limit := len(nums) - len(r)
	for i := 0; i < limit; i++ {
		r = append(r, 0)
	}
	return r
}
