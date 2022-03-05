package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	l1 := make([]*TreeNode, 0)
	l2 := make([]*TreeNode, 0)
	dfs(root1, &l1)
	dfs(root2, &l2)

	if len(l1) != len(l2) {
		return false
	}

	for i := range l1 {
		if l1[i].Val != l2[i].Val {
			return false
		}
	}

	return true
}

func dfs(root *TreeNode, r *[]*TreeNode) {
	if root == nil {
		return
	}

	dfs(root.Left, r)
	if root.Left == nil && root.Right == nil {
		*r = append(*r, root)
	}
	dfs(root.Right, r)
}
