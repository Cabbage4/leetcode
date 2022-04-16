package main

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:  2,
			Left: &TreeNode{Val: 3},
		},
	}
	//root := &TreeNode{
	//	Val: 1,
	//	Left: &TreeNode{
	//		Val:   2,
	//		Left:  &TreeNode{Val: 3},
	//		Right: &TreeNode{Val: 4},
	//	},
	//	Right: &TreeNode{
	//		Val:   5,
	//		Right: &TreeNode{Val: 6},
	//	},
	//}
	flatten(root)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	h, _ := help(root)
	*root = *h
}

func help(root *TreeNode) (*TreeNode, *TreeNode) {
	if root == nil {
		return nil, nil
	}
	if root.Left == nil && root.Right == nil {
		return root, root
	}

	hl, tl := help(root.Left)
	hr, tr := help(root.Right)
	root.Left = nil

	if tl != nil {
		root.Right = hl
		tl.Right = hr
	} else {
		root.Right = hr
	}

	if tr != nil {
		return root, tr
	} else {
		return root, tl
	}
}
