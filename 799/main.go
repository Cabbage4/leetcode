package main

import "fmt"

func main() {
	fmt.Println(champagneTower(100000009, 33, 17))
}

func champagneTower(poured int, row int, col int) float64 {
	dp := make([][]float64, 102)
	for i := range dp {
		dp[i] = make([]float64, 102)
	}

	dp[0][0] = float64(poured)
	for i := 0; i < row; i++ {
		for j := 0; j <= i; j++ {
			p := (dp[i][j] - 1) / 2
			if p < 0 {
				continue
			}

			dp[i+1][j] += p
			dp[i+1][j+1] += p
		}
	}

	if dp[row][col] > 1 {
		return 1
	}
	return dp[row][col]
}
