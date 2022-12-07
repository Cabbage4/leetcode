package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(countSubarrays([]int{2, 3, 1}, 3))
	fmt.Println(countSubarrays([]int{3, 2, 1, 4, 5}, 4))
}

func countSubarrays(nums []int, k int) int {
	sNums := make([]int, len(nums))
	copy(sNums, nums)
	sort.Ints(sNums)

	index := -1
	mp := make(map[int]int)
	for i := range nums {
		mp[nums[i]] = sort.SearchInts(sNums, nums[i])

		if nums[i] == k {
			index = i
		}
	}

	if index == -1 {
		return 0
	}

	r := 1
	left, right := index, index
	direct := [][]int{{0, 1}, {-1, 0}}
	for left >= 0 && right < len(nums) {
		for i := range direct {
			left += direct[i][0]
			right += direct[i][1]

			if left < 0 || right >= len(nums) {
				break
			}

			tt := 0
			if mp[left] < index || mp[right] < index {
				tt--
			}
			if mp[left] > index || mp[right] > index {
				tt++
			}
			if tt == 0 {
				r++
			}
		}
	}

	return r
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
