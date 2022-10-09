package main

import "fmt"

func main() {
	fmt.Println(numTilings(5))
}

func numTilings(n int) int {
	var mod int = 1e9 + 7
	dp := [4]int{1, 0, 0, 0}
	for i := 0; i < n; i++ {
		var ndp [4]int
		ndp[0b00] = (dp[0b00] + dp[0b11]) % mod
		ndp[0b10] = (dp[0b01] + dp[0b00]) % mod
		ndp[0b01] = (dp[0b10] + dp[0b00]) % mod
		ndp[0b11] = (dp[0b00] + dp[0b01] + dp[0b10]) % mod
		dp = ndp
	}
	return dp[0]
}
