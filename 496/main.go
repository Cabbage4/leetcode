package main

import "fmt"

func main() {
	fmt.Println(nextGreaterElement([]int{4, 1, 2}, []int{1, 3, 4, 2}))
	fmt.Println(nextGreaterElement([]int{1, 3, 5, 2, 4}, []int{6, 5, 4, 3, 2, 1, 7}))
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	mp := make(map[int]int)
	for i, v := range nums2 {
		mp[v] = i
	}

	stack := make([]int, 0)
	maxRList := make([]int, len(nums2))
	maxRList[len(nums2)-1] = -1
	stack = append(stack, len(nums2)-1)
	for i := len(nums2) - 2; i >= 0; i-- {
		for len(stack) > 0 && nums2[stack[len(stack)-1]] < nums2[i] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) == 0 {
			maxRList[i] = -1
		} else {
			maxRList[i] = stack[len(stack)-1]
		}

		stack = append(stack, i)
	}

	result := make([]int, len(nums1))

	for i, v := range nums1 {
		numsI := mp[v]
		if maxRList[numsI] != -1 {
			result[i] = nums2[maxRList[numsI]]
		} else {
			result[i] = -1
		}
	}

	return result
}
