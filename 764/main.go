package main

func main() {
}

func orderOfLargestPlusSign(n int, mines [][]int) int {
	mp := make(map[int]bool)
	for _, m := range mines {
		mp[n*m[0]+m[1]] = true
	}

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	var r int
	for i := 0; i < n; i++ {
		count := 0
		for j := 0; j < n; j++ {
			if mp[i*n+j] {
				count = 0
			} else {
				count++
			}
			dp[i][j] = count
		}

		count = 0
		for j := n - 1; j >= 0; j-- {
			if mp[i*n+j] {
				count = 0
			} else {
				count++
			}
			dp[i][j] = min(dp[i][j], count)
		}
	}

	for j := 0; j < n; j++ {
		count := 0
		for i := 0; i < n; i++ {
			if mp[i*n+j] {
				count = 0
			} else {
				count++
			}
			dp[i][j] = min(dp[i][j], count)
		}

		count = 0
		for i := n - 1; i >= 0; i-- {
			if mp[i*n+j] {
				count = 0
			} else {
				count++
			}
			dp[i][j] = min(dp[i][j], count)

			r = max(r, dp[i][j])
		}
	}

	return r
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
