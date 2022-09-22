package main

import "fmt"

func main() {
	fmt.Println(countArrangement(4))
}

func countArrangement(n int) int {
	match := make([][]int, n+1)
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if i%j == 0 || j%i == 0 {
				match[i] = append(match[i], j)
			}
		}
	}

	var r int
	book := make([]bool, n+1)
	var dfs func(int)
	dfs = func(index int) {
		if index > n {
			r++
			return
		}

		for _, v := range match[index] {
			if book[v] {
				continue
			}
			book[v] = true
			dfs(index + 1)
			book[v] = false
		}
	}

	dfs(1)
	return r
}
