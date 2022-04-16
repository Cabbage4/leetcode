package main

import (
	"fmt"
)

func main() {
	test := []int{-1, 0, 0, 0, 0, 3, 3}
	fmt.Println(removeDuplicates(test))
	fmt.Println(test)
}

func removeDuplicates(nums []int) int {
	r := 0
	for _, v := range nums {
		if r < 2 || v > nums[r-2] {
			nums[r] = v
			r++
		}
	}
	return r
}
