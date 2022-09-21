package main

import "fmt"

func main() {
	fmt.Println(findMaxForm([]string{"10", "0001", "111001", "1", "0"}, 5, 3))
}

func findMaxForm(strList []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i < len(strList); i++ {
		zeroCount := 0
		oneCount := 0
		for _, c := range strList[i] {
			if c == '0' {
				zeroCount++
			} else {
				oneCount++
			}
		}

		for j := m; j >= zeroCount; j-- {
			for k := n; k >= oneCount; k-- {
				dp[j][k] = max(dp[j][k], dp[j-zeroCount][k-oneCount]+1)
			}
		}
	}

	return dp[m][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
