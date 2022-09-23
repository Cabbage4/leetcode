package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findPaths(36, 5, 50, 15, 3)) // 390153306
	fmt.Println(findPaths(1, 3, 3, 0, 1))
	fmt.Println(findPaths(2, 2, 2, 0, 0))
}

func findPaths(m int, n int, maxMove int, startRow int, startColumn int) int {
	mod := int(math.Pow10(9)) + 7
	directs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	q := make([][]int, 1)
	q[0] = []int{startRow, startColumn, 1}

	var r int
	for maxMove > 0 {
		size := len(q)

		cmp := make([][]int, m)

		for i := 0; i < size; i++ {
			cur := q[i]
			for _, d := range directs {
				nx := cur[0] + d[0]
				ny := cur[1] + d[1]
				c := cur[2]

				if 0 <= nx && nx < m && 0 <= ny && ny < n {
					if cmp[nx] == nil {
						cmp[nx] = make([]int, n)
					}

					cmp[nx][ny] += c % mod
					continue
				}

				r = (r + c) % mod
			}
		}

		for i := 0; i < m; i++ {
			if cmp[i] == nil {
				continue
			}

			for j := 0; j < n; j++ {
				if cmp[i][j] != 0 {
					q = append(q, []int{i, j, cmp[i][j]})
				}
			}
		}

		maxMove--
		q = q[size:]
	}

	return r
}
