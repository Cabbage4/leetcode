package main

import (
	"fmt"
	"sort"
)

func main() {
	//fmt.Println(frequencySort([]int{1, 1, 2, 2, 2}))
	fmt.Println(frequencySort([]int{2, 3, 1, 2, 1, 3, 2}))
}

func frequencySort(nums []int) []int {
	mp := make([]int, 201)
	for _, v := range nums {
		mp[v+100]++
	}

	sort.Slice(nums, func(i, j int) bool {
		a := mp[nums[i]+100]
		b := mp[nums[j]+100]
		if a == b {
			return nums[i] > nums[j]
		}
		return a < b
	})

	return nums
}
