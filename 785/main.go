package main

import "fmt"

func main() {
	data := [][]int{{1, 3}, {0, 2}, {1, 3}, {0, 2}}
	fmt.Println(isBipartite(data))
}

func isBipartite(graph [][]int) bool {
	color := make([]int, len(graph))
	r := true

	var dfs func(int, int)
	dfs = func(n int, c int) {
		color[n] = c
		nc := -c

		for _, g := range graph[n] {
			if color[g] == 0 {
				dfs(g, nc)
				if !r {
					return
				}
			} else if color[g] == c {
				r = false
				return
			}
		}
	}

	for i := range graph {
		if color[i] == 0 {
			dfs(i, 1)
		}
		if !r {
			break
		}
	}

	return r
}
