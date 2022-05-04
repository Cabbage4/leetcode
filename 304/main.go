package main

import "fmt"

func main() {
	c := Constructor([][]int{{1}, {-7}})
	fmt.Println(c.SumRegion(1, 0, 1, 0))
}

type NumMatrix struct {
	dp [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	dp := make([][]int, len(matrix))

	dp[0] = make([]int, len(matrix[0]))
	dp[0][0] = matrix[0][0]
	for i := 1; i < len(matrix[0]); i++ {
		dp[0][i] = dp[0][i-1] + matrix[0][i]
	}

	for i := 1; i < len(matrix); i++ {
		dp[i] = make([]int, len(matrix[i]))
		dp[i][0] = dp[i-1][0] + matrix[i][0]

		for j := 1; j < len(matrix[i]); j++ {
			dp[i][j] = dp[i][j-1] + dp[i-1][j] + matrix[i][j] - dp[i-1][j-1]
		}
	}
	return NumMatrix{dp: dp}
}

func (n *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	dp := n.dp
	if row1 == 0 && col1 == 0 {
		return dp[row2][col2]
	}

	if row1 == 0 {
		return dp[row2][col2] - dp[row2][col1-1]
	}

	if col1 == 0 {
		return dp[row2][col2] - dp[row1-1][col2]
	}

	return dp[row2][col2] - dp[row1-1][col2] - dp[row2][col1-1] + dp[row1-1][col1-1]
}
