package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxCount(3, 3, [][]int{{2, 2}, {3, 3}}))
}

func maxCount(m int, n int, ops [][]int) int {
	x := math.MaxInt32
	y := math.MaxInt32
	for _, v := range ops {
		if v[0] < x {
			x = v[0]
		}
		if v[1] < y {
			y = v[1]
		}
	}
	if x > m {
		x = m
	}
	if y > n {
		y = n
	}

	return x * y
}
