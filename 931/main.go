package main

import (
	"fmt"
	"math"
)

func main() {
	matrix := [][]int{{-19, 57}, {-40, -5}}
	fmt.Println(minFallingPathSum(matrix))
}

func minFallingPathSum(matrix [][]int) int {
	dp := make([][]int, len(matrix))
	for i := range dp {
		dp[i] = make([]int, len(matrix[0]))
	}

	for i := 0; i < len(matrix[0]); i++ {
		dp[0][i] = matrix[0][i]
	}

	for i := 1; i < len(matrix); i++ {
		dp[i][0] = min(dp[i-1][0], dp[i-1][1]) + matrix[i][0]
		dp[i][len(matrix[i])-1] = min(dp[i-1][len(matrix)-1], dp[i-1][len(matrix)-2]) + matrix[i][len(matrix)-1]

		for j := 1; j < len(matrix[i])-1; j++ {
			dp[i][j] = min(dp[i-1][j-1], dp[i-1][j]) + matrix[i][j]
			dp[i][j] = min(dp[i][j], dp[i-1][j+1]+matrix[i][j])
		}
	}

	r := math.MaxInt
	for i := 0; i < len(matrix[0]); i++ {
		r = min(r, dp[len(matrix)-1][i])
	}

	return r
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
