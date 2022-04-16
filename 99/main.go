package main

import "math"

func main() {
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type H struct {
	pre *TreeNode
	p1  *TreeNode
	p2  *TreeNode
}

func recoverTree(root *TreeNode) {
	h := H{
		pre: &TreeNode{Val: math.MinInt32},
	}
	dfs(root, &h)
	h.p1.Val, h.p2.Val = h.p2.Val, h.p1.Val
}

func dfs(root *TreeNode, h *H) {
	if root.Left != nil {
		dfs(root.Left, h)
	}

	if h.p1 == nil && h.pre.Val > root.Val {
		h.p1 = h.pre
	}
	if h.p1 != nil && h.pre.Val > root.Val {
		h.p2 = root
	}
	h.pre = root

	if root.Right != nil {
		dfs(root.Right, h)
	}
}
