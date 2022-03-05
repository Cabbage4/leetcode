package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isUnivalTree(root *TreeNode) bool {
	return dfs(root) != -1
}

func dfs(root *TreeNode) int {
	if root.Right == nil && root.Left == nil {
		return root.Val
	}

	if root.Left != nil {
		v := dfs(root.Left)
		if v == -1 {
			return -1
		}
		if v != root.Val {
			return -1
		}
	}

	if root.Right != nil {
		v := dfs(root.Right)
		if v == -1 {
			return -1
		}
		if v != root.Val {
			return -1
		}
	}

	return root.Val
}
