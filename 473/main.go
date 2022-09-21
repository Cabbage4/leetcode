package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(makesquare([]int{10, 6, 5, 5, 5, 3, 3, 3, 2, 2, 2, 2}))
	fmt.Println(makesquare([]int{1, 1, 2, 2, 2}))
	fmt.Println(makesquare([]int{3, 3, 3, 3, 4}))
}

func makesquare(matchsticks []int) bool {
	ln := 0
	for _, v := range matchsticks {
		ln += v
	}
	if ln%4 != 0 {
		return false
	}
	ln = ln / 4

	sort.Sort(sort.Reverse(sort.IntSlice(matchsticks)))

	var edges [4]int
	var dfs func(index int) bool
	dfs = func(index int) bool {
		if index == len(matchsticks) {
			return true
		}

		for i := range edges {
			edges[i] += matchsticks[index]
			if edges[i] <= ln && dfs(index+1) {
				return true
			}
			edges[i] -= matchsticks[index]
		}

		return false
	}

	return dfs(0)
}
