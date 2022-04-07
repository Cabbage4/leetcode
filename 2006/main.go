package main

import (
	"fmt"
)

func main() {
	fmt.Println(countKDifference([]int{1, 2, 2, 1}, 1))
}

func countKDifference(nums []int, k int) int {
	mp := make([]int, 101)
	for _, v := range nums {
		mp[v]++
	}

	r := 0
	for i := k + 1; i < 101; i++ {
		r += mp[i] * mp[i-k]
	}
	return r
}
