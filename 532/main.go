package main

import "fmt"

func main() {
	fmt.Println(findPairs([]int{1, 3, 1, 5, 4}, 0))
	fmt.Println(findPairs([]int{1, 2, 3, 4, 5}, 1))
	fmt.Println(findPairs([]int{3, 1, 4, 1, 5}, 2))
}

func findPairs(nums []int, target int) int {
	set := make(map[int]int)
	for _, v := range nums {
		set[v]++
	}

	var r int
	if target == 0 {
		for _, v := range set {
			if v >= 2 {
				r++
			}
		}

		return r
	}

	for k := range set {
		if set[k+target] > 0 {
			r++
		}
	}

	return r
}
