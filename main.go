package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(fourSum([]int{-3, -2, -1, 0, 0, 1, 2, 3}, 0))
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
}

func fourSum(nums []int, target int) [][]int {
	set := make(map[string]bool)
	sort.Ints(nums)
	r := make([][]int, 0)
	for l1 := 0; l1 < len(nums)-4+1; l1++ {
		for l2 := l1 + 1; l2 < len(nums)-3+1; l2++ {
			t := nums[l1] + nums[l2]

			l3 := l2 + 1
			l4 := len(nums) - 1
			for l3 < l4 {
				t2 := t + nums[l3] + nums[l4]
				if t2 == target {
					k := fmt.Sprintf("%d:%d:%d:%d", nums[l1], nums[l2], nums[l3], nums[l4])
					if !set[k] {
						r = append(r, []int{nums[l1], nums[l2], nums[l3], nums[l4]})
						set[k] = true
					}
					l3++
				} else if t2 > target {
					l4--
				} else {
					l3++
				}
			}
		}
	}
	return r
}
