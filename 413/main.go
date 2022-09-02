package main

import "fmt"

func main() {
	fmt.Println(numberOfArithmeticSlices([]int{1, 2, 3}))
	fmt.Println(numberOfArithmeticSlices([]int{1, 2, 3, 4}))
}

func numberOfArithmeticSlices(nums []int) int {
	var r int
	for i := 0; i < len(nums)-1; {
		j := i + 1
		diff := nums[i] - nums[j]

		for ; j < len(nums)-1; j++ {
			if diff != nums[j]-nums[j+1] {
				break
			}
		}

		c := j - i + 1
		if c >= 3 {
			r += c*(c-3)/2 + 1
		}

		i = j
	}
	return r
}
