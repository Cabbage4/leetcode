package main

func main() {
}

func findMiddleIndex(nums []int) int {
	ln := len(nums)
	dp := make([]int, ln)
	dp[0] = nums[0]
	for i := 1; i < ln; i++ {
		dp[i] = dp[i-1] + nums[i]
	}

	for i := 0; i < ln; i++ {
		l := dp[i] - nums[i]
		r := dp[ln-1] - dp[i]

		if l == r {
			return i
		}
	}

	return -1
}
