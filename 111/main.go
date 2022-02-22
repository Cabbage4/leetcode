package main

import "fmt"

func main() {
	//root := &TreeNode{
	//	Val: 1,
	//	Left: &TreeNode{
	//		Val: 2,
	//		Left: &TreeNode{
	//			Val: 4,
	//		},
	//		Right: &TreeNode{
	//			Val: 5,
	//		},
	//	},
	//	Right: &TreeNode{
	//		Val: 3,
	//	},
	//}

	//root := &TreeNode{
	//	Val: 2,
	//	Right: &TreeNode{
	//		Val: 3,
	//		Right: &TreeNode{
	//			Val: 4,
	//			Right: &TreeNode{
	//				Val: 5,
	//				Right: &TreeNode{
	//					Val: 6,
	//				},
	//			},
	//		},
	//	},
	//}
	root := &TreeNode{
		Val: 0,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 5},
				Right: &TreeNode{Val: 1},
			},
		},
		Right: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 3,
				Right: &TreeNode{
					Val:   6,
					Right: &TreeNode{Val: 6},
				},
			},
			Right: &TreeNode{
				Val:   -1,
				Right: &TreeNode{Val: 8},
			},
		},
	}
	fmt.Println(minDepth(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepthv2(root *TreeNode) int {
	if root == nil {
		return 0
	}

	stack := make([]*TreeNode, 0)
	stack = append(stack, root)

	res := 1
	curBucket := 1
	nextBucket := 0

	for len(stack) > 0 {
		peekNode := stack[0]

		if len(stack) > 1 {
			stack = stack[1:]
		} else {
			stack = stack[:0]
		}

		curBucket--

		if peekNode.Left == nil && peekNode.Right == nil {
			break
		}

		if peekNode.Left != nil {
			nextBucket++
			stack = append(stack, peekNode.Left)
		}

		if peekNode.Right != nil {
			nextBucket++
			stack = append(stack, peekNode.Right)
		}

		if curBucket == 0 {
			curBucket = nextBucket
			nextBucket = 0
			res++
		}
	}

	return res
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := minDepth(root.Left)
	right := minDepth(root.Right)

	if left == 0 || right == 0 {
		return left + right + 1
	}
	if left < right {
		return left + 1
	}
	return right + 1
}
