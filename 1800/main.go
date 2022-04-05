package main

import "fmt"

func main() {
	fmt.Println(maxAscendingSum([]int{10, 20, 30, 5, 10, 50}))
}

func maxAscendingSum(nums []int) int {
	r := 0
	for i := 0; i < len(nums); {

		last := nums[i]
		tr := last
		j := i + 1
		for ; j < len(nums); j++ {
			if nums[j] > last {
				last = nums[j]
				tr += nums[j]
			} else {
				break
			}
		}

		i = j
		if tr > r {
			r = tr
		}
	}
	return r
}
