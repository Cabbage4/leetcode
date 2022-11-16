package main

import "fmt"

func main() {
	fmt.Println(spiralMatrixIII(1, 4, 0, 0))
}

var directs = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func spiralMatrixIII(rows int, cols int, x int, y int) [][]int {
	r := make([][]int, 0)
	r = append(r, []int{x, y})
	ln := rows * cols

	if ln == 1 {
		return r
	}

	for k := 1; ; k += 2 {
		for i := 0; i < 4; i++ {
			tmp := k + i/2
			for j := 0; j < tmp; j++ {
				x += directs[i][0]
				y += directs[i][1]

				if 0 <= x && x < rows && 0 <= y && y < cols {
					r = append(r, []int{x, y})
					if len(r) == ln {
						return r
					}
				}
			}
		}
	}
}
