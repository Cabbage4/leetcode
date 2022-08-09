package main

import (
	"fmt"
)

func main() {
	l1 := lexicalOrder(2000)
	fmt.Println(len(l1))
	fmt.Println(l1)
}

func lexicalOrder(n int) []int {
	var r []int
	dfs(&r, n, 1)
	return r
}

func dfs(r *[]int, n int, cur int) {
	if cur > n {
		return
	}
	*r = append(*r, cur)
	dfs(r, n, cur*10)
	if cur%10 < 9 {
		dfs(r, n, cur+1)
	}
}
