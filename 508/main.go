package main

import (
	"fmt"
	"math"
)

func main() {
	root := &TreeNode{
		Val:   5,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: -5},
	}
	fmt.Println(findFrequentTreeSum(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findFrequentTreeSum(root *TreeNode) []int {
	maxValue := math.MinInt
	mp := make(map[int]int)

	var dfs func(*TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		r := root.Val + dfs(root.Left) + dfs(root.Right)
		mp[r]++
		if mp[r] > maxValue {
			maxValue = mp[r]
		}

		return r
	}

	dfs(root)

	var r []int

	for k, v := range mp {
		if v == maxValue {
			r = append(r, k)
		}
	}

	return r
}
