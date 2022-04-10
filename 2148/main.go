package main

import "math"

func main() {
}

func countElements(nums []int) int {
	min := math.MaxInt32
	max := math.MinInt32
	for _, v := range nums {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}

	r := 0
	for _, v := range nums {
		if min < v && v < max {
			r++
		}
	}
	return r
}
