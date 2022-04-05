package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxProductDifference([]int{5, 6, 2, 7, 4}))
}

func maxProductDifference(nums []int) int {
	max1, max2 := -1, -1
	min1, min2 := math.MaxInt32, math.MaxInt32
	for _, v := range nums {
		if v > max1 {
			max2 = max1
			max1 = v
		} else if v > max2 {
			max2 = v
		}

		if v < min1 {
			min2 = min1
			min1 = v
		} else if v < min2 {
			min2 = v
		}
	}
	return max2*max1 - min1*min2
}
