package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(countPrimeSetBits(6, 10))
	fmt.Println(countPrimeSetBits(10, 15))
}

func countPrimeSetBits(left int, right int) int {
	result := 0
	mp := genMp()
	dp := genDp(right)
	for i := left; i <= right; i++ {
		target := dp[i]
		if mp[target] {
			result++
		}
	}

	return result
}

func genDp(n int) []int {
	dp := make([]int, n+1)
	dp[0] = 0

	for i := 1; i <= n; i++ {
		dp[i] = dp[i>>1] + i&1
	}
	return dp
}

func genMp() []bool {
	mp := make([]bool, 128)
	for i := 2; i <= 32; i++ {
		flag := true
		for j := 2; j <= int(math.Sqrt(float64(i))); j++ {
			if i%j == 0 {
				flag = false
				break
			}
		}
		if flag {
			mp[i] = true
		}
	}
	return mp
}
