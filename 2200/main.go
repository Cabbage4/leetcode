package main

import "fmt"

func main() {
	fmt.Println(findKDistantIndices([]int{2, 2, 2, 2, 2}, 2, 2))
	fmt.Println(findKDistantIndices([]int{3, 4, 9, 1, 3, 9, 5}, 9, 1))
}

func findKDistantIndices(nums []int, key int, k int) []int {
	r := make([]int, 0)
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != key {
			continue
		}

		for j = max(j, i-k); j < i+k+1 && j < len(nums); j++ {
			r = append(r, j)
		}
	}
	return r
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
