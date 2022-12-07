package main

import "fmt"

func main() {
	fmt.Println(shortestBridge([][]int{{0, 1, 0}, {0, 0, 0}, {0, 0, 1}}))
	//fmt.Println(shortestBridge([][]int{{0, 1}, {1, 0}}))
}

var directs = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func shortestBridge(grid [][]int) int {
	book := make([][]bool, len(grid))
	for i := range book {
		book[i] = make([]bool, len(grid[i]))
	}

	bfs := func(x, y int) [][]int {
		var r [][]int
		r = append(r, []int{x, y})

		q := make([][]int, 1)
		q[0] = []int{x, y}

		for len(q) > 0 {
			node := q[0]
			q = q[1:]
			for _, direct := range directs {
				nx, ny := node[0]+direct[0], node[1]+direct[1]

				if 0 <= nx && nx < len(grid) && 0 <= ny && ny < len(grid[0]) && !book[nx][ny] && grid[nx][ny] == 1 {
					q = append(q, []int{nx, ny})
					book[nx][ny] = true
					r = append(r, []int{nx, ny})
				}
			}
		}

		return r
	}

	var list1, list2 [][]int
tag:
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 && !book[i][j] {
				list := bfs(i, j)
				if len(list1) == 0 {
					list1 = list
				} else if len(list2) == 0 {
					list2 = list
				} else {
					break tag
				}
			}
		}
	}

	r := 100
	for _, item1 := range list1 {
		for _, item2 := range list2 {
			tmp := abs(item1[0], item2[0]) + abs(item1[1], item2[1]) - 1
			if tmp < r {
				r = tmp
			}
		}
	}

	return r
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
