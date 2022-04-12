package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
}

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	book := make([]bool, len(candidates))
	r := make([][]int, 0)
	dfs(candidates, book, 0, &r, target)
	return r
}

func dfs(list []int, book []bool, index int, result *[][]int, sum int) {
	if sum == 0 {
		r := make([]int, 0)
		for i := 0; i < len(book); i++ {
			if book[i] {
				r = append(r, list[i])
			}
		}
		*result = append(*result, r)
		return
	}

	for i := index; i < len(list); i++ {
		if sum-list[i] < 0 {
			break
		}

		if i > 0 && list[i] == list[i-1] && !book[i-1] {
			continue
		}

		book[i] = true
		dfs(list, book, i+1, result, sum-list[i])
		book[i] = false
	}
}
