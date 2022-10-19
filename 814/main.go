package main

func main() {
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pruneTree(root *TreeNode) *TreeNode {
	var dfs func(*TreeNode) bool
	dfs = func(root *TreeNode) bool {
		if root == nil {
			return false
		}

		if !dfs(root.Left) {
			root.Left = nil
		}
		if !dfs(root.Right) {
			root.Right = nil
		}

		if root.Left == nil && root.Right == nil && root.Val == 0 {
			return false
		}

		return true
	}

	if !dfs(root.Left) {
		root.Left = nil
	}

	if !dfs(root.Right) {
		root.Right = nil
	}

	if root.Left == nil && root.Right == nil && root.Val == 0 {
		return nil
	}

	return root
}
