package main

import (
	"fmt"
)

func main() {
	fmt.Println(numFriendRequests([]int{73, 106, 39, 6, 26, 15, 30, 100, 71, 35, 46, 112, 6, 60, 110}))
	fmt.Println(numFriendRequests([]int{8, 85, 24, 85, 69}))
	fmt.Println(numFriendRequests([]int{20, 30, 100, 110, 120}))
	fmt.Println(numFriendRequests([]int{16, 17, 18}))
	fmt.Println(numFriendRequests([]int{16, 16}))
}

func numFriendRequests(ages []int) int {
	var dp [121]int
	for _, a := range ages {
		dp[a]++
	}

	var r int
	for i := 1; i <= 120; i++ {
		if dp[i] == 0 {
			continue
		}

		if i/2+7 < i {
			r += (dp[i] - 1) * dp[i]
		}

		for j := i - 1; j >= 1; j-- {
			if j <= (i/2 + 7) {
				break
			}

			if dp[j] == 0 {
				continue
			}

			if j > 100 && i < 100 {
				continue
			}

			r += dp[j] * dp[i]
		}
	}

	return r
}
