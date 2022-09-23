package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(optimalDivision([]int{2, 3, 4}))
	fmt.Println(optimalDivision([]int{1000, 100, 10, 2}))
}

func optimalDivision(nums []int) string {
	if len(nums) == 1 {
		return fmt.Sprintf("%d", nums[0])
	} else if len(nums) == 2 {
		return fmt.Sprintf("%d/%d", nums[0], nums[1])
	}
	r := fmt.Sprintf("%d/(", nums[0])
	for i := 1; i < len(nums)-1; i++ {
		r += strconv.Itoa(nums[i]) + "/"
	}
	r = r + strconv.Itoa(nums[len(nums)-1]) + ")"
	return r
}
