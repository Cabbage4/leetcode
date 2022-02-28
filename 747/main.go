package main

import "fmt"

func main() {
	fmt.Println(dominantIndex([]int{1, 0}))
	fmt.Println(dominantIndex([]int{0, 0, 1, 2}))
	fmt.Println(dominantIndex([]int{3, 6, 1, 0}))
}

func dominantIndex(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	m := 0
	nm := -1

	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[m] {
			nm = m
			m = i
			continue
		}

		if nm == -1 || nums[i] > nums[nm] {
			nm = i
		}
	}

	if nums[m] >= 2*nums[nm] {
		return m
	}

	return -1
}
