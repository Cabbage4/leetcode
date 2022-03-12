package main

import "fmt"

func main() {
	fmt.Println(sumOddLengthSubarrays([]int{1, 4, 2, 5, 3}))
	fmt.Println(sumOddLengthSubarrays([]int{1, 2}))
}

func sumOddLengthSubarrays(arr []int) int {
	r := 0

	dp := make([][]int, len(arr)+1)
	for i := 0; i < len(arr)+1; i++ {
		dp[i] = make([]int, len(arr)+1)
	}

	for i, v := range arr {
		dp[1][i] = v
		r += v
	}

	for l := 3; l <= len(arr); l = l + 2 {
		for i := 0; i < len(arr)-l+1; i++ {
			dp[l][i] = dp[l-2][i] + arr[i+l-2] + arr[i+l-1]
			r += dp[l][i]
		}
	}

	return r
}
