package main

func main() {
}

func checkValid(matrix [][]int) bool {
	n := len(matrix)
	for i := 0; i < n; i++ {
		r := 0
		c := 0
		mpr := make([]int, n+1)
		mpc := make([]int, n+1)
		for j := 0; j < n; j++ {
			mpr[matrix[i][j]]++
			if mpr[matrix[i][j]] == 1 {
				r++
			}

			mpc[matrix[j][i]]++
			if mpc[matrix[j][i]] == 1 {
				c++
			}
		}

		if r != n || c != n {
			return false
		}
	}

	return true
}
