package main

import (
	"fmt"
)

func main() {
	fmt.Println(findMaxAverage([]int{5}, 1))
	fmt.Println(findMaxAverage([]int{1, 12, -5, -6, 50, 3}, 4))
}

func findMaxAverage(nums []int, k int) float64 {
	target := 0
	for i := 0; i < k; i++ {
		target += nums[i]
	}

	result := target

	for i := k; i < len(nums); i++ {
		target += nums[i]
		target -= nums[i-k]

		if result < target {
			result = target
		}
	}

	return float64(result) / float64(k)
}
