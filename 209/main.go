package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minSubArrayLen(11, []int{1, 2, 3, 4, 5}))
	fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
}

func minSubArrayLen(target int, nums []int) int {
	ans := math.MaxInt32

	s := 0
	l := 0
	for r := 0; r < len(nums); r++ {
		s += nums[r]
		for s >= target {
			if r-l+1 < ans {
				ans = r - l + 1
			}

			s -= nums[l]
			l++
		}
	}

	if ans == math.MaxInt32 {
		return 0
	}
	return ans
}
