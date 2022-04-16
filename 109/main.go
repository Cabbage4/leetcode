package main

import "fmt"

func main() {
	root := &ListNode{
		Val: -10,
		Next: &ListNode{
			Val: -3,
			Next: &ListNode{
				Val: 0,
				Next: &ListNode{
					Val: 5,
					Next: &ListNode{
						Val:  9,
						Next: nil,
					},
				},
			},
		},
	}
	r := sortedListToBST(root)
	fmt.Println(r)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	return h(head)
}

func h(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}

	wrapH := &ListNode{Next: head}
	tmp := wrapH
	qTmp := wrapH

	for qTmp.Next != nil && qTmp.Next.Next != nil {
		tmp = tmp.Next
		qTmp = qTmp.Next.Next
	}

	root := tmp.Next
	tmp.Next = nil

	return &TreeNode{
		Val:   root.Val,
		Left:  h(wrapH.Next),
		Right: h(root.Next),
	}
}
