package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}

func maxProfit(prices []int) int {
	var r int
	min := math.MaxInt
	for i := range prices {
		if prices[i] < min {
			min = prices[i]
		}

		if prices[i]-min > r {
			r = prices[i] - min
		}
	}
	return r
}
