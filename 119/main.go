package main

import "fmt"

func main() {
	fmt.Println(getRow(3))
}

func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}
	if rowIndex == 1 {
		return []int{1, 1}
	}

	dp := make([][]int, rowIndex+1)
	for i := 0; i <= rowIndex; i++ {
		dp[i] = make([]int, i+1)
	}

	dp[0][0] = 1
	dp[1][0] = 1
	dp[1][1] = 1

	for i := 2; i <= rowIndex; i++ {
		dp[i][0] = 1
		dp[i][len(dp[i])-1] = 1

		for j := 1; j < len(dp[i])-1; j++ {
			dp[i][j] = dp[i-1][j] + dp[i-1][j-1]
		}
	}

	return dp[rowIndex]
}
