package main

import (
	"fmt"
)

func main() {
	fmt.Println(minEatingSpeed([]int{3, 6, 7, 11}, 8))
	fmt.Println(minEatingSpeed([]int{30, 11, 23, 4, 20}, 5))
}

func minEatingSpeed(piles []int, h int) int {
	left := 1
	right := 0
	for _, v := range piles {
		if right < v {
			right = v
		}
	}

	for left < right {
		mid := (left + right) / 2
		c := 0
		for _, v := range piles {
			c += (v + mid - 1) / mid
		}

		if c > h {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left
}
