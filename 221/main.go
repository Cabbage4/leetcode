package main

import "fmt"

func main() {
	fmt.Println(maximalSquare([][]byte{
		{'0'},
	}))
	fmt.Println(maximalSquare([][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}))
	fmt.Println(maximalSquare([][]byte{
		{'0', '1'},
		{'1', '0'},
	}))
}

func maximalSquare(matrix [][]byte) int {
	r := 0
	dp := make([][]int, len(matrix))
	for i := range dp {
		dp[i] = make([]int, len(matrix[0]))
	}
	for i := 0; i < len(matrix[0]); i++ {
		dp[0][i] = int(matrix[0][i] - '0')
		r = max(r, dp[0][i])
	}
	for i := 0; i < len(matrix); i++ {
		dp[i][0] = int(matrix[i][0] - '0')
		r = max(r, dp[i][0])
	}

	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][j] == '1' {
				dp[i][j] = min(min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
				r = max(r, dp[i][j])
			}
		}
	}

	return r * r
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
