package main

func main() {
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 3},
			Right: &TreeNode{Val: 1},
		},
		Right: &TreeNode{
			Val:  6,
			Left: &TreeNode{Val: 5},
		},
	}

	root = addOneRow(root, 1, 2)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	return dfs(root, true, val, depth)
}

func dfs(root *TreeNode, isLeft bool, val int, depth int) *TreeNode {
	if depth == 1 {
		if isLeft {
			return &TreeNode{
				Val:  val,
				Left: root,
			}
		}

		return &TreeNode{
			Val:   val,
			Right: root,
		}
	}

	if depth == 0 || root == nil {
		return root
	}

	root.Left = dfs(root.Left, true, val, depth-1)
	root.Right = dfs(root.Right, false, val, depth-1)
	return root
}
