package main

import "math"

func main() {

}

func smallestRangeI(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}

	max := math.MinInt32
	min := math.MaxInt32

	for _, v := range nums {
		if v > max {
			max = v
		}

		if v < min {
			min = v
		}
	}

	if max-min > 2*k {
		return max - min - 2*k
	}

	return 0
}
