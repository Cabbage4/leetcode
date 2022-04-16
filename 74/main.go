package main

import "fmt"

func main() {
	//test := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	test := [][]int{{1}}
	fmt.Println(searchMatrix(test, 0))
}

func searchMatrix(matrix [][]int, target int) bool {
	rln := len(matrix)
	cln := len(matrix[0])

	left := 0
	right := rln * cln
	for left < right {
		mid := (left + right) / 2
		i := mid / cln
		j := mid - i*cln
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return false
}
