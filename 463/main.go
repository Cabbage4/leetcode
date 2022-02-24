package main

import "fmt"

func main() {
	test := [][]int{{0, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 0, 0}}
	fmt.Println(islandPerimeter(test))
}

var direct = [][]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

func islandPerimeter(grid [][]int) int {
	result := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				result += helper(grid, i, j)
			}
		}
	}

	return result
}

func helper(grid [][]int, i, j int) int {
	result := 4
	for _, v := range direct {
		ni := i + v[0]
		nj := j + v[1]
		if ni >= 0 && nj >= 0 && ni < len(grid) && nj < len(grid[0]) {
			result = result - grid[ni][nj]
		}
	}
	return result
}
