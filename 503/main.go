package main

import "fmt"

func main() {
	fmt.Println(nextGreaterElements([]int{5, 4, 3, 2, 1}))
	//fmt.Println(nextGreaterElements([]int{1, 2, 3, 4, 3}))
	//fmt.Println(nextGreaterElements([]int{1, 2, 1}))
}

func nextGreaterElements(nums []int) []int {
	maxIndex := 0
	for i, v := range nums {
		if nums[maxIndex] < v {
			maxIndex = i
		}
	}

	ln := len(nums)
	dp := make([]int, ln)
	dp[maxIndex] = -1

	for i := (maxIndex - 1 + ln) % ln; i != maxIndex; i = (i - 1 + ln) % ln {
		j := (i + 1 + ln) % ln
		for j != -1 && nums[i] >= nums[j] {
			j = dp[j]
		}

		dp[i] = j
	}

	for i, v := range dp {
		if v == -1 {
			continue
		}

		dp[i] = nums[dp[i]]
	}

	return dp
}
