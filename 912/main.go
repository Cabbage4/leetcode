package main

import "fmt"

func main() {
	fmt.Println(sortArray([]int{5, 2, 3, 1}))
	fmt.Println(sortArray([]int{5, 1, 1, 2, 0, 0}))
}

func sortArray(nums []int) []int {
	var diff int = 1e4 * 5
	list := make([]int, 1e5+1)
	for _, v := range nums {
		list[v+diff]++
	}

	index := 0
	for i := range list {
		for list[i] > 0 {
			nums[index] = i - diff
			index++
			list[i]--
		}
	}
	return nums
}
