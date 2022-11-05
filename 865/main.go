package main

import "fmt"

func main() {
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val:  5,
			Left: &TreeNode{Val: 6},
			Right: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 7},
				Right: &TreeNode{Val: 4},
			},
		},
		Right: &TreeNode{
			Val:   1,
			Left:  &TreeNode{Val: 0},
			Right: &TreeNode{Val: 8},
		},
	}

	result := subtreeWithAllDeepest(root)
	fmt.Println(result)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
	list := make([]*TreeNode, 0)
	pmp := make(map[*TreeNode]*TreeNode)
	dmp := make(map[*TreeNode]int)

	var dfs func(*TreeNode, *TreeNode, int)
	dfs = func(root *TreeNode, parent *TreeNode, depth int) {
		if root == nil {
			return
		}

		pmp[root] = parent
		dmp[root] = depth

		if len(list) == 0 {
			list = []*TreeNode{root}
		} else {
			d := dmp[root]
			md := dmp[list[0]]

			if md < d {
				list = []*TreeNode{root}
			} else if md == d {
				list = append(list, root)
			}
		}

		dfs(root.Left, root, depth+1)
		dfs(root.Right, root, depth+1)
	}

	dfs(root, nil, 0)

	for len(list) != 1 {
		tmp := make(map[*TreeNode]bool)
		for _, v := range list {
			tmp[pmp[v]] = true
		}

		tmpList := make([]*TreeNode, 0)
		for k := range tmp {
			tmpList = append(tmpList, k)
		}

		list = tmpList
	}

	return list[0]
}
