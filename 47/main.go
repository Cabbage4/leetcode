package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(permuteUnique([]int{1, 1, 2}))
}

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	r := make([][]int, 0)
	dfs(nums, make([]bool, len(nums)), &r, make([]int, 0))
	return r
}

func dfs(list []int, book []bool, result *[][]int, keyList []int) {
	if len(keyList) == len(list) {
		r := make([]int, len(list))
		for i, key := range keyList {
			r[i] = list[key]
		}
		*result = append(*result, r)
		return
	}

	for i := 0; i < len(list); i++ {
		if book[i] {
			continue
		}

		if i > 0 && !book[i-1] && list[i] == list[i-1] {
			continue
		}

		book[i] = true
		dfs(list, book, result, append(keyList, i))
		book[i] = false
	}
}
