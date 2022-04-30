package main

import "fmt"

func main() {
	fmt.Println(combinationSum3(4, 1))
	fmt.Println(combinationSum3(3, 9))
	fmt.Println(combinationSum3(3, 7))
}

func combinationSum3(k int, n int) [][]int {
	r := make([][]int, 0)
	dfs(1, &r, make([]int, 0), k, n)
	return r
}

func dfs(ci int, r *[][]int, cr []int, k int, n int) {
	if k == 0 {
		if n == 0 {
			ncr := make([]int, len(cr))
			copy(ncr, cr)
			*r = append(*r, ncr)
		}
		return
	}

	for i := ci; i <= 9 && i <= n; i++ {
		dfs(i+1, r, append(cr, i), k-1, n-i)
	}
}
