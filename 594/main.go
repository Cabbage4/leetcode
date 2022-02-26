package main

import (
	"fmt"
)

func main() {
	fmt.Println(findLHS([]int{1, 1, 1, 1}))
	fmt.Println(findLHS([]int{1, 3, 2, 2, 5, 2, 3, 7}))
}

func findLHS(nums []int) int {
	result := 0
	mp := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		mp[nums[i]]++
	}

	for k, v := range mp {
		if vp, has := mp[k+1]; has {
			result = max(result, v+vp)
		}
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
