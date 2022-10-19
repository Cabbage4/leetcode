package main

import "fmt"

func main() {
	//fmt.Println(largestSumOfAverages([]int{1, 2, 3, 4, 5, 6, 7}, 4))
	fmt.Println(largestSumOfAverages([]int{9, 1, 2, 3, 9}, 3))
}

func largestSumOfAverages(nums []int, k int) float64 {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = dp[i-1] + nums[i]
	}

	var mp [100][100]float64
	var dfs func(int, int) float64
	dfs = func(left, count int) (ans float64) {
		if v := mp[left][count]; v != 0 {
			return v
		}

		defer func() {
			mp[left][count] = ans
		}()

		if count == 1 {
			if left == 0 {
				return float64(dp[len(nums)-1]) / float64(len(nums))
			}

			return float64(dp[len(nums)-1]-dp[left-1]) / float64(len(nums)-left)
		}

		var r float64
		for i := left; i < len(nums)-1; i++ {
			var tr float64
			if i == left {
				tr = float64(nums[i]) + dfs(i+1, count-1)
			} else {
				if left == 0 {
					tr = float64(dp[i])/float64(i+1) + dfs(i+1, count-1)
				} else {
					tr = float64(dp[i]-dp[left-1])/float64(i-left+1) + dfs(i+1, count-1)
				}
			}

			if tr > r {
				r = tr
			}
		}

		return r
	}

	return dfs(0, k)
}
