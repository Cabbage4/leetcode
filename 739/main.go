package main

import "fmt"

func main() {
	fmt.Println(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))
}

func dailyTemperatures(list []int) []int {
	dp := make([]int, len(list))
	dp[len(list)-1] = -1

	for i := len(list) - 2; i >= 0; i-- {
		j := i + 1
		for j != -1 && list[i] >= list[j] {
			j = dp[j]
		}

		dp[i] = j
	}

	for i := 0; i < len(dp); i++ {
		if dp[i] == -1 {
			dp[i] = 0
			continue
		}
		dp[i] = dp[i] - i
	}

	return dp
}
