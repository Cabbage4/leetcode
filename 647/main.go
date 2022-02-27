package main

import "fmt"

func main() {
	fmt.Println(findLengthOfLCIS([]int{1, 3, 5, 4, 7}))
	fmt.Println(findLengthOfLCIS([]int{1, 3, 5, 7}))
}

func findLengthOfLCIS(nums []int) int {
	max := 1

	for i := 0; i < len(nums); {
		j := i + 1
		for j < len(nums) && nums[j-1] < nums[j] {
			j++
		}

		if j-i > max {
			max = j - i
		}
		i = j
	}

	return max
}
