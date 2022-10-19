package main

import (
	"fmt"
)

func main() {
	var data = [][]int{{1, 2}, {2, 3}, {5}, {0}, {5}, {}, {}}
	fmt.Println(eventualSafeNodes(data))
}

func eventualSafeNodes(graph [][]int) []int {
	var r []int

	color := make(map[int]int)
	var dfs func(int) bool
	dfs = func(index int) bool {
		if color[index] != 0 {
			return color[index] == 2
		}

		color[index] = 1
		for _, v := range graph[index] {
			if !dfs(v) {
				return false
			}
		}

		color[index] = 2
		return true
	}

	for i := range graph {
		if dfs(i) {
			r = append(r, i)
		}
	}
	return r
}
