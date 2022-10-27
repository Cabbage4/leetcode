package main

import "fmt"

func main() {
	data := [][]int{
		{3, 10, 2, 3, 4},
		{4, 5, 6, 8, 1},
		{8, 8, 1, 6, 8},
		{1, 3, 5, 7, 1},
		{9, 4, 9, 2, 9},
	}
	//data := [][]int{{4, 3, 8, 4}, {9, 5, 1, 9}, {2, 7, 6, 2}}
	fmt.Println(numMagicSquaresInside(data))
}

func numMagicSquaresInside(g [][]int) int {
	if len(g) < 0 || len(g[0]) < 3 {
		return 0
	}

	var r int
	//rowDp := make([]int, len(g[0]))
	//colDp := make([]int, len(g))
	for i := 0; i < len(g)-2; i++ {
		for j := 0; j < len(g[i])-2; j++ {
			if i == 2 && j == 1 {
				fmt.Println()
			}
			flag := false

			set := make(map[int]bool)
			for ii := i; ii <= i+2; ii++ {
				for jj := j; jj <= j+2; jj++ {
					if set[g[ii][jj]] || g[ii][jj] > 9 || g[ii][jj] == 0 {
						flag = true
						break
					}

					set[g[ii][jj]] = true
				}
			}
			if flag {
				continue
			}

			row1 := g[i][j] + g[i][j+1] + g[i][j+2]
			row2 := g[i+1][j] + g[i+1][j+1] + g[i+1][j+2]
			row3 := g[i+2][j] + g[i+2][j+1] + g[i+2][j+2]
			col1 := g[i][j] + g[i+1][j] + g[i+2][j]
			col2 := g[i][j+1] + g[i+1][j+1] + g[i+2][j+1]
			col3 := g[i][j+2] + g[i+1][j+2] + g[i+2][j+2]
			ltr := g[i][j] + g[i+1][j+1] + g[i+2][j+2]
			rtl := g[i][j+2] + g[i+1][j+1] + g[i+2][j]
			if row1 == row2 && row2 == row3 && row3 == col1 && col1 == col2 && col2 == col3 && col3 == ltr && ltr == rtl {
				r++
			}
		}
	}
	return r
}
