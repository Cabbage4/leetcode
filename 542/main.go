package main

import "fmt"

func main() {
	fmt.Println(updateMatrix([][]int{{0, 1, 0}, {0, 1, 0}, {0, 1, 0}}))
}

func updateMatrix(mat [][]int) [][]int {
	r := make([][]int, len(mat))
	book := make([][]bool, len(mat))
	for i := range book {
		r[i] = make([]int, len(mat[i]))
		book[i] = make([]bool, len(mat[i]))
	}

	q := make([][]int, 0)
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			if mat[i][j] == 0 {
				q = append(q, []int{i, j})
				book[i][j] = true
			}
		}
	}

	directs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	for len(q) != 0 {
		cur := q[0]
		q = q[1:]

		x := cur[0]
		y := cur[1]

		book[x][y] = true

		for _, d := range directs {
			nx := x + d[0]
			ny := y + d[1]
			if 0 <= nx && nx < len(mat) && 0 <= ny && ny < len(mat[0]) {
				if book[nx][ny] {
					continue
				}
				book[nx][ny] = true
				r[nx][ny] = r[x][y] + 1
				q = append(q, []int{nx, ny})
			}
		}
	}

	return r
}
