package main

import "fmt"

func main() {
	fmt.Println(averageValue([]int{4, 4, 9, 10}))
}

func averageValue(nums []int) int {
	sum := 0
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i]%3 == 0 && nums[i]%2 == 0 {
			count++
			sum += nums[i]
		}
	}

	if count == 0 {
		return 0
	}
	return sum / count
}
