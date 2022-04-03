package main

import "fmt"

func main() {
	fmt.Println(check([]int{3, 4, 5, 4, 2}))
}

func check(nums []int) bool {
	c := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] < nums[(i+1)%len(nums)] {
			c++
		}
	}

	return c <= 1
}
