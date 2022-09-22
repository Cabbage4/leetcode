package main

import "fmt"

func main() {
	fmt.Println(findMaxLength([]int{0, 1}))
}

func findMaxLength(nums []int) int {
	var r int

	mp := make(map[int]int)
	mp[0] = -1
	c := 0
	for i, v := range nums {
		if v == 1 {
			c++
		} else {
			c--
		}

		if preIndex, ok := mp[c]; ok {
			r = max(r, i-preIndex)
		} else {
			mp[c] = i
		}
	}
	return r
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
