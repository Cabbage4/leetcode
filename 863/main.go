package main

func main() {
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	mp := make(map[*TreeNode]*TreeNode)
	var dfs func(*TreeNode, *TreeNode)
	dfs = func(root *TreeNode, parent *TreeNode) {
		if root == nil {
			return
		}

		mp[root] = parent

		dfs(root.Left, root)
		dfs(root.Right, root)
	}
	dfs(root, nil)

	var r []int

	book := make(map[*TreeNode]bool)
	var dfs2 func(*TreeNode, int)
	dfs2 = func(root *TreeNode, count int) {
		if root == nil || book[root] {
			return
		}

		book[root] = true
		if count == 0 {
			r = append(r, root.Val)
			return
		}

		dfs2(root.Left, count-1)
		dfs2(root.Right, count-1)
		dfs2(mp[root], count-1)
	}

	dfs2(target, k)

	return r
}
