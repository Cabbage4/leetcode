package main

import (
	"fmt"
	"math"
)

func main() {
	root := &TreeNode{
		Val: 27,
		Right: &TreeNode{
			Val: 34,
			Right: &TreeNode{
				Val: 58,
				Left: &TreeNode{
					Val:  50,
					Left: &TreeNode{Val: 44},
				},
			},
		},
	}

	fmt.Println(minDiffInBST(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDiffInBST(root *TreeNode) int {
	result := math.MaxInt32
	open := false
	cur := 0
	dfs(root, &open, &cur, &result)

	return result
}

func dfs(root *TreeNode, open *bool, cur *int, result *int) {
	if root.Left != nil {
		dfs(root.Left, open, cur, result)
	}

	if *open && root.Val-*cur < *result {
		*result = root.Val - *cur
	}

	*cur = root.Val
	*open = true

	if root.Right != nil {
		dfs(root.Right, open, cur, result)
	}
}
