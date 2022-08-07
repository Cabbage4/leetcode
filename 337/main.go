package main

func main() {
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	v := h(root)
	return max(v[0], v[1])
}

func h(root *TreeNode) []int {
	if root == nil {
		return []int{0, 0}
	}

	l := h(root.Left)
	r := h(root.Right)
	return []int{max(l[0], l[1]) + max(r[0], r[1]), l[0] + r[0] + root.Val}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
