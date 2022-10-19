package main

import "fmt"

func main() {
	fmt.Println(countSubarrays([]int{1, 3, 5, 2, 7, 5}, 1, 5))
	//fmt.Println(countSubarrays([]int{1, 1, 1, 1}, 1, 1))
}

func countSubarrays(nums []int, minK int, maxK int) int64 {
	var r int64
	maxI, minI, i0 := -1, -1, -1
	for i, v := range nums {
		if v == minK {
			minI = i
		}
		if v == maxK {
			maxI = i
		}
		if maxK < v || v < minK {
			i0 = i
		}

		r += int64(max(min(maxI, minI)-i0, 0))
	}
	return r
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
