package main

import "math"

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {
	list := make([]int, 0)
	dfs(root, &list)

	result := math.MaxInt32
	for i := 0; i < len(list)-1; i++ {
		if list[i+1]-list[i] < result {
			result = list[i+1] - list[i]
		}
	}
	return result
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
