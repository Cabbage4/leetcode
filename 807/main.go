package main

import "fmt"

func main() {
	//data := [][]int{{3, 0, 8, 4}, {2, 4, 5, 7}, {9, 2, 6, 3}, {0, 3, 1, 0}}
	data := [][]int{{59, 88, 44}, {3, 18, 38}, {21, 26, 51}}
	fmt.Println(maxIncreaseKeepingSkyline(data))
}

func maxIncreaseKeepingSkyline(grid [][]int) int {
	n := len(grid)
	rl := make([]int, n)
	cl := make([]int, n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			rl[i] = max(grid[i][j], rl[i])
			cl[i] = max(grid[j][i], cl[i])
		}
	}

	var r int
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			t := min(rl[i], cl[j]) - grid[i][j]
			r += t
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

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
