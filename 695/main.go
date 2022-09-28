package main

import "fmt"

func main() {
	data := [][]int{{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}, {0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0}, {0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0}, {0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0}}
	fmt.Println(maxAreaOfIsland(data))
}

func maxAreaOfIsland(grid [][]int) int {
	directs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	var r int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 0 {
				continue
			}
			grid[i][j] = 0

			tr := 0
			q := make([][]int, 1)
			q[0] = []int{i, j}
			for len(q) != 0 {
				cur := q[0]
				q = q[1:]
				tr++

				for _, d := range directs {
					nx, ny := cur[0]+d[0], cur[1]+d[1]
					if 0 <= nx && nx < len(grid) && 0 <= ny && ny < len(grid[i]) && grid[nx][ny] == 1 {
						grid[nx][ny] = 0
						q = append(q, []int{nx, ny})
					}
				}
			}

			if tr > r {
				r = tr
			}
		}
	}

	return r
}
