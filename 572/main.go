package main

import "fmt"

func main() {
	//one := &TreeNode{
	//	Val: 3,
	//	Left: &TreeNode{
	//		Val:  4,
	//		Left: &TreeNode{Val: 1},
	//	},
	//	Right: &TreeNode{
	//		Val:  5,
	//		Left: &TreeNode{Val: 2},
	//	},
	//}
	//two := &TreeNode{
	//	Val:   3,
	//	Left:  &TreeNode{Val: 1},
	//	Right: &TreeNode{Val: 2},
	//}

	one := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 6},
				Right: &TreeNode{Val: 7},
			},
		},
	}
	two := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   1,
			Left:  &TreeNode{Val: 6},
			Right: &TreeNode{Val: 7},
		},
	}
	fmt.Println(isSubtree(one, two))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}

	if same(root, subRoot) {
		return true
	}

	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)

}

func same(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil && subRoot == nil {
		return true
	} else if root == nil || subRoot == nil {
		return false
	}

	if root.Val != subRoot.Val {
		return false
	}

	return same(root.Left, subRoot.Left) && same(root.Right, subRoot.Right)
}
