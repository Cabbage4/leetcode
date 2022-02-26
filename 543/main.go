package main

import (
	"fmt"
)

func main() {
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}
	//root := &TreeNode{
	//	Val: 1,
	//	Left: &TreeNode{
	//		Val:   2,
	//		Left:  &TreeNode{Val: 4},
	//		Right: &TreeNode{Val: 5},
	//	},
	//	Right: &TreeNode{Val: 3},
	//}
	fmt.Println(diameterOfBinaryTree(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	result := 0
	depth(root, &result)
	return result
}

func depth(root *TreeNode, result *int) int {
	if root == nil {
		return -1
	}

	left := depth(root.Left, result)
	right := depth(root.Right, result)
	*result = max(*result, left+right+2)
	return max(left, right) + 1
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
