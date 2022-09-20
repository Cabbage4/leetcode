package main

import "fmt"

func main() {
	root := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 8},
			Right: &TreeNode{Val: 13},
		},
		Right: &TreeNode{
			Val:   5,
			Left:  &TreeNode{Val: 21},
			Right: &TreeNode{Val: 34},
		},
	}

	result := reverseOddLevels(root)
	fmt.Println(result)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func reverseOddLevels(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	q := []*TreeNode{root}
	level := 0
	for len(q) != 0 {
		levelCount := len(q)

		for i := 0; i < levelCount; i++ {
			tmpNode := q[i]
			if tmpNode.Left != nil {
				q = append(q, tmpNode.Left)
			}
			if tmpNode.Right != nil {
				q = append(q, tmpNode.Right)
			}
		}

		if level%2 != 0 {
			for i := 0; i < levelCount/2; i++ {
				q[i].Val, q[levelCount-i-1].Val = q[levelCount-i-1].Val, q[i].Val
			}
		}

		q = q[levelCount:]
		level++
	}

	return root
}
