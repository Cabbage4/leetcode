package main

import (
	"fmt"
)

func main() {
	fmt.Println(findShortestSubArray([]int{2, 1, 1, 2, 1, 3, 3, 3, 1, 3, 1, 3, 2}))
	//fmt.Println(findShortestSubArray([]int{1, 2, 2, 3, 1}))
	//fmt.Println(findShortestSubArray([]int{1, 2, 2, 3, 1, 4, 2}))
}

func findShortestSubArray(nums []int) int {
	mp := make(map[int][]int)
	result := 1
	max := 1
	for i, v := range nums {
		if mp[v] == nil {
			mp[v] = []int{i, i, 1}
			continue
		}

		list := mp[v]
		list[2]++
		if list[0] > i {
			list[0] = i
		}
		if list[1] < i {
			list[1] = i
		}

		if list[2] == max {
			result = min(result, list[1]-list[0]+1)
		} else if list[2] > max {
			max = list[2]
			result = list[1] - list[0] + 1
		}
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
