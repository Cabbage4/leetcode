package main

func main() {
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	var dfs func(*TreeNode)
	dfs = func(r *TreeNode) {
		if r.Val > val {
			if r.Left == nil {
				r.Left = &TreeNode{Val: val}
			} else {
				dfs(r.Left)
			}
		} else {
			if r.Right == nil {
				r.Right = &TreeNode{Val: val}
			} else {
				dfs(r.Right)
			}
		}
	}

	dfs(root)
	return root
}
