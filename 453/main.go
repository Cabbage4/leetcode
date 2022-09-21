package main

import "fmt"

func main() {
	fmt.Println(minMoves([]int{1, 1, 1000000000}))
}

func minMoves(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	min := nums[0]
	for _, v := range nums {
		if v < min {
			min = v
		}
	}

	var r int
	for _, v := range nums {
		r += v - min
	}

	return r
}
