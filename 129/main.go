package main

func main() {
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}

	r := 0
	dfs(root, &r, 0)
	return r
}

func dfs(root *TreeNode, r *int, cr int) {
	cv := cr*10 + root.Val
	if root.Left == nil && root.Right == nil {
		*r += cv
	}
	if root.Left != nil {
		dfs(root.Left, r, cv)
	}
	if root.Right != nil {
		dfs(root.Right, r, cv)
	}
}
