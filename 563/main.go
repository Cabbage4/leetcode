package main

import "fmt"

func main() {
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 3},
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{
			Val:   9,
			Right: &TreeNode{Val: 7},
		},
	}

	fmt.Println(findTilt(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTilt(root *TreeNode) int {
	if root == nil {
		return 0
	}
	result := 0
	sum(root, &result)
	return result
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func sum(root *TreeNode, result *int) int {
	if root == nil {
		return 0
	}

	left := sum(root.Left, result)
	right := sum(root.Right, result)
	*result += abs(left, right)
	return root.Val + left + right
}
