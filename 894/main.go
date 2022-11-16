package main

func main() {
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func allPossibleFBT(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	if n == 1 {
		return []*TreeNode{{}}
	}

	r := make([]*TreeNode, 0)
	for size := 1; size < n-1; size++ {
		leftList := allPossibleFBT(size)
		rightList := allPossibleFBT(n - 1 - size)

		for _, leftRoot := range leftList {
			for _, rightRoot := range rightList {
				r = append(r, &TreeNode{
					Left:  fork(leftRoot),
					Right: fork(rightRoot),
				})
			}
		}
	}

	return r
}

func fork(node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}

	return &TreeNode{
		Left:  fork(node.Left),
		Right: fork(node.Right),
	}
}
