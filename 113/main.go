package main

func main() {
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return nil
	}

	r := make([][]int, 0)
	dfs(root, targetSum, &r, make([]int, 0))
	return r
}

func dfs(root *TreeNode, target int, r *[][]int, cr []int) {
	if root.Left == nil && root.Right == nil {
		if target == root.Val {
			t := make([]int, len(cr)+1)
			copy(t, cr)
			t[len(t)-1] = root.Val
			*r = append(*r, t)
		}
		return
	}

	if root.Left != nil {
		dfs(root.Left, target-root.Val, r, append(cr, root.Val))
	}
	if root.Right != nil {
		dfs(root.Right, target-root.Val, r, append(cr, root.Val))
	}
}
