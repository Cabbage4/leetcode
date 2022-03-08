package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(largestPerimeter([]int{9, 5, 3, 3, 4, 5, 19, 3}))
	fmt.Println(largestPerimeter([]int{1, 4, 18, 3, 8, 4, 4}))
	fmt.Println(largestPerimeter([]int{2, 1, 2}))
}

func largestPerimeter(a []int) int {
	sort.Ints(a)
	for i := len(a) - 3; i >= 0; i-- {
		if a[i]+a[i+1] > a[i+2] {
			return a[i] + a[i+1] + a[i+2]
		}
	}

	return 0
}
