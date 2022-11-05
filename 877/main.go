package main

import "fmt"

func main() {
	fmt.Println(stoneGame([]int{3, 2, 10, 4}))
	fmt.Println(stoneGame([]int{5, 3, 4, 5}))
	fmt.Println(stoneGame([]int{3, 7, 2, 3}))
}

func stoneGame(piles []int) bool {
	dp := make([][]int, len(piles))
	for i := range dp {
		dp[i] = make([]int, len(piles))
		dp[i][i] = piles[i]
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for i := 0; i < len(piles)-1; i++ {
		for j := i + 1; j < len(piles); j++ {
			dp[i][j] = max(piles[i]-dp[i+1][j], piles[j]-dp[i][j-1])
		}
	}

	return dp[0][len(piles)-1] > 0
}

//
//func stoneGame(piles []int) bool {
//	return true
//}
