package main

import "fmt"

func main() {
	root := constructFromPrePost([]int{1, 2, 4, 5, 3, 6, 7}, []int{4, 5, 2, 6, 7, 3, 1})
	fmt.Println(root)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}

	preMp := make(map[int]int)
	postMp := make(map[int]int)
	for i := range preorder {
		preMp[preorder[i]] = i
		postMp[postorder[i]] = i
	}

	var dfs func(int, int, int, int) *TreeNode
	dfs = func(preL, preR, postL, postR int) *TreeNode {
		if preR-preL == 0 {
			return nil
		}

		r := &TreeNode{Val: postorder[postR-1]}
		if postR-postL == 1 {
			return r
		}

		splitIndex := postMp[preorder[preL+1]]

		ln := splitIndex - postL + 1

		r.Left = dfs(preL+1, preL+ln+1, postL, splitIndex+1)
		r.Right = dfs(preL+ln+1, preR, splitIndex+1, postR-1)

		return r
	}

	return dfs(0, len(preorder), 0, len(postorder))
}
