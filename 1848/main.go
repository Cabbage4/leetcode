package main

import "math"

func main() {
}

func getMinDistance(nums []int, target int, start int) int {
	r := math.MaxInt32
	for i, v := range nums {
		if v != target {
			continue
		}
		tr := abs(i, start)
		if tr < r {
			r = tr
		}
	}
	return r
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
