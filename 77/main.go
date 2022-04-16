package main

import "fmt"

func main() {
	fmt.Println(combine(4, 2))
}

func combine(n int, k int) [][]int {
	r := make([][]int, 0)
	dfs(k, &r, make([]int, 0), n, 1)
	return r
}

func dfs(k int, r *[][]int, cr []int, n int, index int) {
	if len(cr) == k {
		tr := make([]int, len(cr))
		copy(tr, cr)
		*r = append(*r, tr)
		return
	}

	for i := index; i <= n; i++ {
		dfs(k, r, append(cr, i), n, i+1)
	}
}
