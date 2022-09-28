package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(knightProbability(8, 30, 6, 4))
	fmt.Println(knightProbability(3, 2, 0, 0))
}

func knightProbability(n int, k int, row int, column int) float64 {
	directs := [][]int{
		{-1, -2},
		{-1, 2},
		{-2, -1},
		{-2, 1},
		{1, -2},
		{1, 2},
		{2, -1},
		{2, 1},
	}

	mp := make([][][]*float64, n)
	for i := range mp {
		mp[i] = make([][]*float64, n)
		for j := range mp[i] {
			mp[i][j] = make([]*float64, k)
		}
	}

	var dfs func(int, int, int) float64
	dfs = func(c int, x int, y int) float64 {
		if c == k {
			return 1
		}

		if v := mp[x][y][c]; v != nil {
			return *v
		}

		var r float64
		for _, d := range directs {
			nx, ny := x+d[0], y+d[1]
			if 0 <= nx && nx < n && 0 <= ny && ny < n {
				r += dfs(c+1, nx, ny)
			}
		}

		mp[x][y][c] = &r
		return r
	}
	r := dfs(0, row, column)

	return r / math.Pow(8, float64(k))
}
