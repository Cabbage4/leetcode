package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(combinationSum4([]int{1, 2}, 2))
	fmt.Println(combinationSum4([]int{1, 2, 3}, 35))
}

func combinationSum4(nums []int, target int) int {
	sort.Ints(nums)
	mp := make([]int, target+1)
	mp[0] = 1

	for i := 1; i <= target; i++ {
		for _, item := range nums {
			if item > i {
				break
			}

			mp[i] += mp[i-item]
		}
	}

	return mp[target]
}
