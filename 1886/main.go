package main

import "fmt"

func main() {
	fmt.Println(findRotation([][]int{{1, 1}, {0, 0}}, [][]int{{0, 1}, {0, 1}}))
}

func findRotation(mat [][]int, target [][]int) bool {
	n := len(mat)
	f1 := true
	f2 := true
	f3 := true
	f4 := true
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if f1 && mat[i][j] != target[i][j] {
				f1 = false
			}
			if f2 && mat[i][j] != target[n-i-1][n-j-1] {
				f2 = false
			}
			if f3 && mat[i][j] != target[j][n-i-1] {
				f3 = false
			}
			if f4 && mat[i][j] != target[n-j-1][i] {
				f4 = false
			}

			//if !f1 && !f2 && !f3 && !f4 {
			//	return false
			//}
		}
	}

	return f1 || f2 || f3 || f4
}
