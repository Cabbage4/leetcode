package main

import "fmt"

func main() {
	fmt.Println(maxChunksToSorted([]int{1, 2, 0, 3}))
	fmt.Println(maxChunksToSorted([]int{2, 0, 1}))
	fmt.Println(maxChunksToSorted([]int{1, 0, 2, 3, 4}))
	fmt.Println(maxChunksToSorted([]int{4, 3, 2, 1, 0}))
}

func maxChunksToSorted(arr []int) int {
	r := 0
	maxValue := -1
	for i := 0; i < len(arr); i++ {
		maxValue = max(maxValue, arr[i])
		if maxValue == i {
			r++
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
