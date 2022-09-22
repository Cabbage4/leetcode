package main

import "fmt"

func main() {
	fmt.Println(findDiagonalOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
}

func findDiagonalOrder(mat [][]int) []int {
	if len(mat) == 1 {
		return mat[0]
	}

	directs := [][]int{{-1, 1}, {1, -1}}
	selectForward := 0

	r := []int{mat[0][0]}
	x := 0
	y := 0
	for x != len(mat)-1 || y != len(mat[0])-1 {
		nx := directs[selectForward][0] + x
		ny := directs[selectForward][1] + y

		if 0 <= nx && nx < len(mat) && 0 <= ny && ny < len(mat[0]) {
			x = nx
			y = ny

			r = append(r, mat[x][y])
			continue
		}

		if selectForward == 0 {
			if y == len(mat[0])-1 {
				x++
			} else {
				y++
			}
		} else {
			if x == len(mat)-1 {
				y++
			} else {
				x++
			}
		}

		selectForward = (selectForward + 1) % 2
		r = append(r, mat[x][y])
	}

	return r
}
