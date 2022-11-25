package main

import "fmt"

func main() {
	fmt.Println(maxSubarraySumCircular([]int{-3, -2, -3}))
	fmt.Println(maxSubarraySumCircular([]int{1, -2, 3, -2}))
	fmt.Println(maxSubarraySumCircular([]int{5, -3, 5}))
}

func maxSubarraySumCircular(nums []int) int {
	r := nums[0]
	curSum := nums[0]
	for i := 1; i < len(nums); i++ {
		curSum = nums[i] + max(curSum, 0)
		r = max(r, curSum)
	}

	rightSum := make([]int, len(nums))
	rightSum[len(nums)-1] = nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		rightSum[i] = rightSum[i+1] + nums[i]
	}

	maxRight := make([]int, len(nums))
	maxRight[len(nums)-1] = nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		maxRight[i] = max(maxRight[i+1], rightSum[i])
	}

	leftSum := 0
	for i := 0; i < len(nums)-2; i++ {
		leftSum += nums[i]
		r = max(r, leftSum+maxRight[i+2])
	}

	return r
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
