package main

import "fmt"

func main() {
	data := [][]int{{520626, 685427, 788912, 800638, 717251, 683428}, {23602, 608915, 697585, 957500, 154778, 209236}, {287585, 588801, 818234, 73530, 939116, 252369}}
	fmt.Println(maxSum(data))
}

var directs = [][]int{
	{-1, -1},
	{-1, 0},
	{-1, 1},

	{1, -1},
	{1, 0},
	{1, 1},
}

func maxSum(grid [][]int) int {
	var r int
	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[i])-1; j++ {
			tmp := grid[i][j]
			for _, d := range directs {
				nx := i + d[0]
				ny := j + d[1]
				if 0 <= nx && nx < len(grid) && 0 <= ny && ny < len(grid[i]) {
					tmp += grid[nx][ny]
				}
			}

			if tmp > r {
				r = tmp
			}
		}
	}
	return r
}
