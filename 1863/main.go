package main

import "fmt"

func main() {
	fmt.Println(subsetXORSum([]int{1, 3}))
	fmt.Println(subsetXORSum([]int{5, 1, 6}))
}

func subsetXORSum(nums []int) int {
	r := 0
	for l := 1; l <= len(nums); l++ {
		dfs(nums, 0, &r, 0, l, 0)
	}
	return r
}

func dfs(list []int, index int, r *int, tr int, c int, tc int) {
	if c == tc {
		*r += tr
		return
	}

	for i := index; i < len(list); i++ {
		dfs(list, i+1, r, tr^list[i], c, tc+1)
	}
}
