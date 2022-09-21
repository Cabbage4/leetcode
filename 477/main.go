package main

import "fmt"

func main() {
	fmt.Println(totalHammingDistance([]int{4, 14, 2}))
}

func totalHammingDistance(nums []int) int {
	var r int
	for i := 0; i < 32; i++ {
		var c int
		for _, v := range nums {
			c += (v >> i) & 1
		}
		r += c * (len(nums) - c)
	}

	return r
}
