package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	root := &TreeNode{
		Val:  1,
		Left: &TreeNode{Val: 2},
	}

	fmt.Println(printTree(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func printTree(root *TreeNode) [][]string {
	if root == nil {
		return nil
	}

	h := height(root)
	result := make([][]string, h+1)
	n := int(math.Pow(2, float64(h+1))) - 1
	for i := range result {
		result[i] = make([]string, n)
	}

	var dfs func(*TreeNode, int, int)
	dfs = func(root *TreeNode, r int, c int) {
		if root == nil {
			return
		}

		g := int(math.Pow(2, float64(h-r-1)))

		if root.Left != nil {
			result[r+1][c-g] = strconv.Itoa(root.Left.Val)
			dfs(root.Left, r+1, c-g)
		}

		if root.Right != nil {
			result[r+1][c+g] = strconv.Itoa(root.Right.Val)
			dfs(root.Right, r+1, c+g)
		}
	}

	result[0][(n-1)/2] = strconv.Itoa(root.Val)
	dfs(root, 0, (n-1)/2)

	return result
}

func height(root *TreeNode) int {
	if root == nil {
		return -1
	}

	lh := height(root.Left)
	rh := height(root.Right)

	if lh > rh {
		return lh + 1
	}

	return rh + 1
}
