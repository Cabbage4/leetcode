package main

func main() {
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	return dfs(root)
}

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return dfs(root.Left) + dfs(root.Right) + 1
}

func bfs(root *TreeNode) int {
	if root == nil {
		return 0
	}

	r := 1
	q := make([]*TreeNode, 0)
	q = append(q, root)
	for len(q) != 0 {
		t := q[0]
		q = q[1:]
		if t.Left != nil {
			r++
			q = append(q, t.Left)
		}
		if t.Right != nil {
			r++
			q = append(q, t.Right)
		}
	}

	return r
}
