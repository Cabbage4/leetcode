package main

import "fmt"

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Right: &TreeNode{Val: 4},
		},
		Right: &TreeNode{Val: 3},
	}

	fmt.Println(isCousins(root, 2, 3))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isCousins(root *TreeNode, x int, y int) bool {
	dx := 0
	dy := 0
	var rx *TreeNode = nil
	var ry *TreeNode = nil
	dfs(root, 0, x, y, &dx, &dy, &rx, &ry)

	if dx != dy {
		return false
	}

	return is(root, rx, ry)
}

func is(root *TreeNode, a, b *TreeNode) bool {
	if root == nil {
		return true
	}

	if root.Left == a && root.Right == b ||
		root.Right == a && root.Left == b {
		return false
	}

	return is(root.Left, a, b) && is(root.Right, a, b)
}

func dfs(root *TreeNode, d int, x, y int, dx, dy *int, rx, ry **TreeNode) {
	if root == nil {
		return
	}

	if root.Val == x {
		*rx = root
		*dx = d
	}
	if root.Val == y {
		*ry = root
		*dy = d
	}

	if *rx != nil && *ry != nil {
		return
	}

	dfs(root.Left, d+1, x, y, dx, dy, rx, ry)

	if *rx != nil && *ry != nil {
		return
	}

	dfs(root.Right, d+1, x, y, dx, dy, rx, ry)
}
