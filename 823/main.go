package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(numFactoredBinaryTrees([]int{2, 4}))
	fmt.Println(numFactoredBinaryTrees([]int{2, 4, 5, 10}))
}

func numFactoredBinaryTrees(arr []int) int {
	var mod int = 1e9 + 7
	sort.Ints(arr)
	imp := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		imp[arr[i]] = i
	}

	var r int
	dp := make([]int, len(arr))
	for i, v := range arr {
		dp[i] = 1
		for _, a := range arr[:i] {
			if a > int(math.Sqrt(float64(v))) || v%a != 0 {
				continue
			}

			ai := imp[a]
			bi, ok := imp[v/a]
			if !ok {
				continue
			}

			if bi == ai {
				dp[i] = (dp[i] + dp[ai]*dp[bi]) % mod
			} else {
				dp[i] = (dp[i] + dp[ai]*dp[bi]*2) % mod
			}
		}

		r = (r + dp[i]) % mod
	}

	return r
}
