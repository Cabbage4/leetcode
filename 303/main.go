package main

import "fmt"

func main() {
	c := Constructor([]int{-2, 0, 3, -5, 2, -1})

	fmt.Println(c.SumRange(2, 5))
}

type NumArray struct {
	dp []int
}

func Constructor(nums []int) NumArray {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = dp[i-1] + nums[i]
	}
	return NumArray{dp: dp}
}

func (n *NumArray) SumRange(left int, right int) int {
	if left == 0 {
		return n.dp[right]
	}

	return n.dp[right] - n.dp[left-1]
}
