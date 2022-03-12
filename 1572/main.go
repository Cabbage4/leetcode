package main

func main() {
}

func diagonalSum(mat [][]int) int {
	r := 0
	ln := len(mat)
	for i := 0; i < ln; i++ {
		r = r + mat[i][i] + mat[i][ln-i-1]
	}

	if ln%2 != 0 {
		r = r - mat[ln/2][ln/2]
	}

	return r
}
