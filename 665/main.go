package main

import (
	"fmt"
)

func main() {
	fmt.Println(checkPossibility([]int{4, 2, 3}))
}

func checkPossibility(nums []int) bool {
	count := 0
	for i := 0; i < len(nums)-1; i++ {
		a, b := nums[i], nums[i+1]
		if a <= b {
			continue
		}

		count++
		if count > 1 {
			return false
		}

		if i > 0 && b < nums[i-1] {
			nums[i+1] = a
		}
	}
	return true
}
