package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(deleteAndEarn([]int{2, 2, 3, 3, 3, 4}))
}

func deleteAndEarn(nums []int) int {
	var r int

	sort.Ints(nums)
	list := make([]int, 0)
	list = append(list, nums[0])

	for i := 1; i < len(nums); i++ {
		if val := nums[i]; val == nums[i-1] {
			list[len(list)-1] += val
		} else if val-nums[i-1] == 1 {
			list = append(list, val)
		} else {
			r += rob(list)
			list = []int{val}
		}
	}
	r += rob(list)

	return r
}

func rob(list []int) int {
	if len(list) == 1 {
		return list[0]
	}

	dp2, dp1 := list[0], max(list[0], list[1])
	for i := 2; i < len(list); i++ {
		dp2, dp1 = dp1, max(dp2+list[i], dp1)
	}

	return dp1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
