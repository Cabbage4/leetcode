package main

import (
	"fmt"
	"math"
)

func main() {
	root := &TreeNode{
		Val:  2,
		Left: &TreeNode{Val: 2},
		Right: &TreeNode{
			Val:   5,
			Left:  &TreeNode{Val: 5},
			Right: &TreeNode{Val: 5},
		},
	}
	fmt.Println(findSecondMinimumValue(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findSecondMinimumValue(root *TreeNode) int {
	result := math.MaxInt64
	dfs(root, root.Val, &result)
	if result < math.MaxInt64 {
		return result
	}

	return -1
}

func dfs(root *TreeNode, min int, result *int) {
	if root == nil {
		return
	}

	if root.Val > min && root.Val < *result {
		*result = root.Val
	}
	if root.Val == min {
		dfs(root.Left, min, result)
		dfs(root.Right, min, result)
	}
}
