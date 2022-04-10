package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(threeSumClosest([]int{1, 2, 5, 10, 11}, 12))
	fmt.Println(threeSumClosest([]int{0, 0, 0}, 1))
	fmt.Println(threeSumClosest([]int{-1, 2, 1, 4}, 1))
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	r := 0
	ab := math.MaxInt32
	for t3 := 0; t3 < len(nums)-2; t3++ {
		i := t3 + 1
		j := len(nums) - 1

		for i < j {
			if nums[i]+nums[j]+nums[t3] == target {
				return nums[i] + nums[j] + nums[t3]
			}

			if tab := abs(nums[i]+nums[j]+nums[t3], target); tab < ab {
				ab = tab
				r = nums[i] + nums[j] + nums[t3]
			}

			if nums[i]+nums[j]+nums[t3] > target {
				j--
			} else {
				i++
			}
		}
	}

	return r
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
