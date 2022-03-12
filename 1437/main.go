package main

import "fmt"

func main() {
	fmt.Println(kLengthApart([]int{1, 0, 0, 1, 0, 1}, 2))
	fmt.Println(kLengthApart([]int{1, 0, 0, 0, 1, 0, 0, 1}, 2))
}

func kLengthApart(nums []int, k int) bool {
	for i := 0; i < len(nums); {
		if nums[i] != 1 {
			i++
			continue
		}

		j := i + 1
		for j < len(nums) && nums[j] == 0 {
			j++
		}

		if j >= len(nums) {
			break
		}

		if j-i-1 < k {
			return false
		}

		i = j
	}

	return true
}
