package main

func main() {
	root := &TreeNode{
		Val:   5,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 7},
	}
	increasingBST(root)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func increasingBST(root *TreeNode) *TreeNode {
	list := make([]int, 0)
	dfs(root, &list)

	if len(list) == 0 {
		return nil
	}

	r := &TreeNode{Val: list[0]}
	t := r
	for _, v := range list[1:] {
		t.Right = &TreeNode{Val: v}
		t = t.Right
	}

	return r
}

func dfs(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}

	dfs(root.Left, result)
	*result = append(*result, root.Val)
	dfs(root.Right, result)
}
