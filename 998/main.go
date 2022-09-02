package main

func main() {
	//root := &TreeNode{
	//	Val:  4,
	//	Left: &TreeNode{Val: 1},
	//	Right: &TreeNode{
	//		Val:  3,
	//		Left: &TreeNode{Val: 2},
	//	},
	//}

	root2 := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val:   2,
			Right: &TreeNode{Val: 1},
		},
		Right: &TreeNode{Val: 4},
	}

	//insertIntoMaxTree(root, 5)
	insertIntoMaxTree(root2, 3)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val < val {
		tmp := *root
		*root = TreeNode{
			Val:   val,
			Left:  &tmp,
			Right: nil,
		}
		return root
	}

	if root.Right != nil {
		insertIntoMaxTree(root.Right, val)
	} else {
		root.Right = &TreeNode{
			Val: val,
		}
	}

	return root
}
