package main

import "fmt"

func main() {
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val:   4,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 1},
		},
		Right: &TreeNode{
			Val:   5,
			Right: &TreeNode{Val: 5},
		},
	}
	fmt.Println(longestUnivaluePath(root))

	root2 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:  2,
				Left: &TreeNode{Val: 2},
			},
			Right: &TreeNode{Val: 2},
		},
		Right: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 2},
			Right: &TreeNode{Val: 2},
		},
	}

	fmt.Println(longestUnivaluePath(root2))

	root3 := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 1},
			},
			Right: &TreeNode{
				Val:  1,
				Left: &TreeNode{Val: 1},
			},
		},
	}

	fmt.Println(longestUnivaluePath(root3))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func longestUnivaluePath(root *TreeNode) int {
	var r int
	dfs(root, &r)
	return r
}

func dfs(root *TreeNode, r *int) int {
	if root == nil {
		return 0
	}

	left := dfs(root.Left, r)
	right := dfs(root.Right, r)
	var arrowLeft int
	var arrowRight int
	if root.Left != nil && root.Left.Val == root.Val {
		arrowLeft += left + 1
	}
	if root.Right != nil && root.Right.Val == root.Val {
		arrowRight += right + 1
	}

	if arrowLeft+arrowRight > *r {
		*r = arrowLeft + arrowRight
	}

	if arrowLeft > arrowRight {
		return arrowLeft
	}
	return arrowRight
}
