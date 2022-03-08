package main

import "fmt"

func main() {
	fmt.Println(allCellsDistOrder(1, 2, 0, 0))
}

type co []int

var dl = [][]int{
	{1, 0},
	{0, 1},
	{-1, 0},
	{0, -1},
}

func allCellsDistOrder(rows int, cols int, rCenter int, cCenter int) [][]int {
	book := make([][]int, rows)
	for i := 0; i < len(book); i++ {
		book[i] = make([]int, cols)
	}

	r := make([][]int, 0)

	q := make([]co, 0)
	book[rCenter][cCenter] = 1
	q = append(q, []int{rCenter, cCenter})
	r = append(r, []int{rCenter, cCenter})
	for len(q) != 0 {
		c := q[0]
		q = q[1:]

		x := c[0]
		y := c[1]
		book[x][y] = 2
		for _, d := range dl {
			nx := x + d[0]
			ny := y + d[1]
			if nx >= 0 && nx < rows &&
				ny >= 0 && ny < cols &&
				book[nx][ny] == 0 {
				book[nx][ny] = 1
				nn := []int{nx, ny}
				r = append(r, nn)
				q = append(q, nn)
			}
		}
	}

	return r
}
