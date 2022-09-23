package main

import "fmt"

func main() {
	fmt.Println(minDistance("a", "ab"))
	fmt.Println(minDistance("leetcode", "etco"))
	fmt.Println(minDistance("sea", "eat"))
}

func minDistance(s1 string, s2 string) int {
	dp := make([][]int, len(s1)+1)
	for i := 0; i <= len(s1); i++ {
		dp[i] = make([]int, len(s2)+1)
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for i := 1; i <= len(s1); i++ {
		for j := 1; j <= len(s2); j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return len(s2) - 2*dp[len(s1)][len(s2)] + len(s1)
}
