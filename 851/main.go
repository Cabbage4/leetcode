package main

import (
	"fmt"
)

func main() {
	richer := [][]int{{1, 0}, {2, 1}, {3, 1}, {3, 7}, {4, 3}, {5, 3}, {6, 3}}
	quiet := []int{3, 2, 5, 4, 6, 1, 7, 0}
	fmt.Println(loudAndRich(richer, quiet))
}

func loudAndRich(richer [][]int, quiet []int) []int {
	dig := make([][]int, len(quiet))
	for _, r := range richer {
		dig[r[1]] = append(dig[r[1]], r[0])
	}

	r := make([]int, len(quiet))
	for i := range r {
		r[i] = -1
	}

	var dfs func(int)
	dfs = func(index int) {
		if r[index] != -1 {
			return
		}

		r[index] = index
		for _, d := range dig[index] {
			dfs(d)
			if quiet[r[d]] < quiet[r[index]] {
				r[index] = r[d]
			}
		}
	}

	for i := range quiet {
		dfs(i)
	}

	return r
}
