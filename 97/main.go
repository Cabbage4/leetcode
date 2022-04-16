package main

import "fmt"

func main() {
	fmt.Println(isInterleave("", "", ""))
	fmt.Println(isInterleave("aabcc", "dbbca", "aadbbcbcac"))
	fmt.Println(isInterleave("aabcc", "dbbca", "aadbbbaccc"))
}

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}

	dp := make([][]bool, len(s1)+1)
	dp[0] = make([]bool, len(s2)+1)
	dp[0][0] = true

	for i := 1; i <= len(s1); i++ {
		dp[i] = make([]bool, len(s2)+1)
		dp[i][0] = dp[i-1][0] && (s1[i-1] == s3[i-1])
	}
	for i := 1; i <= len(s2); i++ {
		dp[0][i] = dp[0][i-1] && (s2[i-1] == s3[i-1])
	}

	for i := 1; i <= len(s1); i++ {
		for j := 1; j <= len(s2); j++ {
			dp[i][j] = (dp[i-1][j] && s1[i-1] == s3[i+j-1]) || (dp[i][j-1] && s2[j-1] == s3[i+j-1])
		}
	}

	return dp[len(s1)][len(s2)]
}
