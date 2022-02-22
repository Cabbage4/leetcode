package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	newValue := targetSum - root.Val
	if root.Right == nil && root.Left == nil && newValue == 0 {
		return true
	}

	if root.Left != nil && hasPathSum(root.Left, newValue) {
		return true
	}
	if root.Right != nil && hasPathSum(root.Right, newValue) {
		return true
	}

	return false
}
