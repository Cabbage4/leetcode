package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(advantageCount([]int{2, 7, 11, 15}, []int{1, 10, 4, 11}))
}

func advantageCount(nums1 []int, nums2 []int) []int {
	idx1 := make([]int, len(nums1))
	idx2 := make([]int, len(nums1))
	for i := range nums1 {
		idx1[i] = i
		idx2[i] = i
	}

	sort.Slice(idx1, func(i, j int) bool {
		return nums1[idx1[i]] < nums1[idx1[j]]
	})
	sort.Slice(idx2, func(i, j int) bool {
		return nums2[idx2[i]] < nums2[idx2[j]]
	})

	left := 0
	right := len(nums1) - 1
	r := make([]int, len(nums1))
	for i := range nums1 {
		if nums1[idx1[i]] > nums2[idx2[left]] {
			r[idx2[left]] = nums1[idx1[i]]
			left++
		} else {
			r[idx2[right]] = nums1[idx1[i]]
			right--
		}
	}
	return r
}
