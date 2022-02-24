package main

import "fmt"

func main() {
	fmt.Println(findMaxConsecutiveOnes([]int{1, 1, 0, 1, 1, 1}))
}

func findMaxConsecutiveOnes(nums []int) int {
	result := 0
	count := 0
	for _, v := range nums {
		if v == 1 {
			count++
			if count > result {
				result = count
			}
			continue
		}
		count = 0
	}

	return result
}
