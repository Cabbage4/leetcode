package main

func main() {
}

func construct2DArray(original []int, m int, n int) [][]int {
	ln := len(original)
	if ln != m*n {
		return make([][]int, 0)
	}

	r := make([][]int, m)
	oi := 0
	for i := 0; i < m; i++ {
		r[i] = make([]int, n)
		for j := 0; j < n; j++ {
			r[i][j] = original[oi]
			oi++
		}
	}

	return r
}
