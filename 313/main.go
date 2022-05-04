package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(nthSuperUglyNumber(12, []int{2, 7, 13, 19}))
}

func nthSuperUglyNumber(n int, primes []int) int {
	pi := make([]int, len(primes))
	dp := make([]int, n)
	dp[0] = 1
	for i := 1; i < n; i++ {
		dp[i] = math.MaxInt32
		mini := -1

		for j := 0; j < len(primes); {
			cTmp := primes[j] * dp[pi[j]]
			if cTmp <= dp[i-1] {
				pi[j]++
				continue
			}

			if cTmp < dp[i] {
				dp[i] = cTmp
				mini = j
			}
			j++
		}

		pi[mini]++
	}
	return dp[n-1]
}
