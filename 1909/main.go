package main

import "fmt"

func main() {
	fmt.Println(canBeIncreasing([]int{105, 924, 32, 968}))
	fmt.Println(canBeIncreasing([]int{1, 2, 10, 5, 7}))
	fmt.Println(canBeIncreasing([]int{2, 3, 1, 2}))
}

func canBeIncreasing(nums []int) bool {
	flag := false
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] < nums[i+1] {
			continue
		}

		if flag {
			return false
		}

		if i == 0 || nums[i-1] < nums[i+1] {
			flag = true
		} else if i+2 == len(nums) || nums[i] < nums[i+2] {
			flag = true
			i++
		} else {
			return false
		}
	}

	return true
}
