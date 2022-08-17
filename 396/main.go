package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxRotateFunction([]int{100}))
	fmt.Println(maxRotateFunction([]int{4, 3, 2, 6}))
}

func maxRotateFunction(nums []int) int {
	sum := 0
	avgSum := 0
	for i := range nums {
		avgSum += i * nums[i]
		sum += nums[i]
	}

	r := avgSum
	for i := len(nums) - 1; i >= 0; i-- {
		avgSum = avgSum + sum - nums[i]*len(nums)
		if avgSum > r {
			r = avgSum
		}
	}

	return r
}
