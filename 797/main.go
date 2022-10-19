package main

import "fmt"

func main() {
	fmt.Println(allPathsSourceTarget([][]int{{1, 2}, {3}, {3}, {}}))
}

func allPathsSourceTarget(graph [][]int) [][]int {
	var r [][]int

	var dfs func(int, string)
	dfs = func(index int, path string) {
		if index == len(graph)-1 {
			list := make([]int, len(path))
			for i := 0; i < len(path); i++ {
				list[i] = int(path[i] - '0')
			}
			r = append(r, list)
			return
		}

		if len(graph[index]) == 0 {
			return
		}

		for _, g := range graph[index] {
			dfs(g, path+string(g+'0'))
		}
	}

	dfs(0, "0")

	return r
}
