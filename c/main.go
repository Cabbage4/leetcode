package main

import "fmt"

func main() {
	fmt.Println(countDistinctIntegers([]int{2, 2, 2}))
	fmt.Println(countDistinctIntegers([]int{1, 13, 10, 12, 31}))
}

func countDistinctIntegers(nums []int) int {
	set := make(map[int]bool)
	re := func(v int) int {
		var r int
		for v != 0 {
			r = r*10 + v%10
			v = v / 10
		}
		return r
	}

	for _, v := range nums {
		set[v] = true
		set[re(v)] = true
	}

	return len(set)
}
