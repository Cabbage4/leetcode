package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minimumCost([]int{5, 5, 5, 5}))
	fmt.Println(minimumCost([]int{6, 5, 7, 9, 2, 2}))
	fmt.Println(minimumCost([]int{1, 2, 3}))
}

func minimumCost(cost []int) int {
	r := 0
	sort.Ints(cost)
	for i := len(cost) - 1; i >= 0; i = i - 3 {
		if i == 0 {
			r += cost[i]
			continue
		}

		r += cost[i-1] + cost[i]
	}

	return r
}
