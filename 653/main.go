package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {
	list := make([]int, 0)
	dfs(root, &list)
	left := 0
	right := len(list) - 1

	for left < right {
		v := list[left] + list[right]
		if v < k {
			left++
		} else if v > k {
			right--
		} else {
			return true
		}
	}

	return false
}

func dfs(root *TreeNode, list *[]int) {
	if root.Left != nil {
		dfs(root.Left, list)
	}
	*list = append(*list, root.Val)
	if root.Right != nil {
		dfs(root.Right, list)
	}
}
