package main

import (
	"fmt"
)

func main() {
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 2},
			Right: &TreeNode{Val: 4},
		},
		Right: &TreeNode{
			Val:   6,
			Right: &TreeNode{Val: 7},
		},
	}

	root = deleteNode(root, 5)
	fmt.Println("done")
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val < key {
		root.Right = deleteNode(root.Right, key)
		return root
	}

	if root.Val > key {
		root.Left = deleteNode(root.Left, key)
		return root
	}

	if root.Left == nil || root.Right == nil {
		if root.Left != nil {
			return root.Left
		}

		return root.Right
	}

	successor := root.Right
	for successor.Left != nil {
		successor = successor.Left
	}

	successor.Right = deleteNode(root.Right, successor.Val)
	successor.Left = root.Left

	return successor
}
