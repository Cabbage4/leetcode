package main

func main() {
}

func new21Game(n int, k int, maxPts int) float64 {
	if k == 0 {
		return 1
	}

	dp := make([]float64, k+maxPts)
	for i := k; i <= n && i < k+maxPts; i++ {
		dp[i] = 1.0
	}

	dp[k-1] = 1.0 * float64(min(n-k+1, maxPts)) / float64(maxPts)
	for i := k - 2; i >= 0; i-- {
		dp[i] = dp[i+1] - (dp[i+maxPts+1]-dp[i+1])/float64(maxPts)
	}
	return dp[0]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
