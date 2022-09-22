package main

import (
	"fmt"
)

func main() {
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:  1,
			Left: &TreeNode{Val: 0},
			Right: &TreeNode{
				Val:   2,
				Right: &TreeNode{Val: 3},
			},
		},
		Right: &TreeNode{
			Val:  6,
			Left: &TreeNode{Val: 5},
			Right: &TreeNode{
				Val:   7,
				Right: &TreeNode{Val: 8},
			},
		},
	}

	root = convertBST(root)
	fmt.Println(root)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func convertBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	sum := 0
	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}

		dfs(root.Right)
		root.Val, sum = sum+root.Val, sum+root.Val
		dfs(root.Left)

	}

	dfs(root)

	return root
}
