package main

import "fmt"

func main() {
	fmt.Println(minimumDeleteSum("sea", "tea"))
}

func minimumDeleteSum(s1 string, s2 string) int {
	dp := make([][]int, len(s1)+1)
	for i := range dp {
		dp[i] = make([]int, len(s2)+1)
	}

	for i := 0; i < len(s1); i++ {
		dp[i+1][0] = dp[i][0] + int(s1[i])
	}

	for i := 0; i < len(s2); i++ {
		dp[0][i+1] = dp[0][i] + int(s2[i])
	}

	for i := 0; i < len(s1); i++ {
		for j := 0; j < len(s2); j++ {
			if s1[i] == s2[j] {
				dp[i+1][j+1] = dp[i][j]
			} else {
				dp[i+1][j+1] = min(dp[i+1][j]+int(s2[j]), dp[i][j+1]+int(s1[i]))
			}
		}
	}

	return dp[len(s1)][len(s2)]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
