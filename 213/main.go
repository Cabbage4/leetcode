package main

import "fmt"

func main() {
	fmt.Println(rob([]int{2, 7, 7, 7, 4}))
	fmt.Println(rob([]int{2, 7, 9, 3, 1}))
	fmt.Println(rob([]int{200, 3, 140, 20, 10}))
	fmt.Println(rob([]int{2, 1, 1, 2}))
	fmt.Println(rob([]int{1, 2, 3, 1}))
	fmt.Println(rob([]int{2, 3, 2}))
}

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	ln := len(nums)
	return max(rob1(nums[:ln-1]), rob1(nums[1:]))
}

func rob1(nums []int) int {
	r := 0
	pre := 0
	for i := 0; i < len(nums); i++ {
		r, pre = max(r, pre+nums[i]), r
	}
	return r
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
