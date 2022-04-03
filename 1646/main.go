package main

func main() {
}

func getMaximumGenerated(n int) int {
	if n <= 1 {
		return n
	}

	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	r := 0

	for i := 2; i <= n; i++ {
		if i%2 == 0 {
			dp[i] = dp[i/2]
		} else {
			dp[i] = dp[i/2] + dp[i/2+1]
		}

		if dp[i] > r {
			r = dp[i]
		}
	}

	return r
}
