package main

import "fmt"

func main() {
	fmt.Println(wiggleMaxLength([]int{0, 0}))
	fmt.Println(wiggleMaxLength([]int{1, 1, 7, 4, 9, 2, 5}))
	fmt.Println(wiggleMaxLength([]int{1, 7, 4, 9, 2, 5}))
	fmt.Println(wiggleMaxLength([]int{1, 17, 5, 10, 13, 15, 10, 5, 16, 8}))
}

func wiggleMaxLength(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	up := make([]int, len(nums))
	down := make([]int, len(nums))
	up[0] = 1
	down[0] = 1

	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			up[i] = max(up[i], down[i-1]+1)
			down[i] = down[i-1]
		} else if nums[i] < nums[i-1] {
			up[i] = up[i-1]
			down[i] = max(down[i], up[i-1]+1)
		} else {
			up[i] = up[i-1]
			down[i] = down[i-1]
		}
	}

	return max(up[len(nums)-1], down[len(nums)-1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
