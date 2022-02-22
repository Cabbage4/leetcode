package main

import (
	"fmt"
	"strconv"
)

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	result := make([]string, 0)
	if root == nil {
		return result
	}
	if root.Right == nil && root.Left == nil {
		result = append(result, strconv.Itoa(root.Val))
		return result
	}

	dfs(root.Left, &result, fmt.Sprintf("%d", root.Val))
	dfs(root.Right, &result, fmt.Sprintf("%d", root.Val))
	return result
}

func dfs(root *TreeNode, result *[]string, cur string) {
	if root == nil {
		return
	}

	next := fmt.Sprintf("%s->%d", cur, root.Val)
	dfs(root.Right, result, next)
	dfs(root.Left, result, next)
	if root.Right == nil && root.Left == nil {
		*result = append(*result, next)
	}
}
