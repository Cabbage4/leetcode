package main

import "math"

func main() {
	root := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  1,
				Next: &ListNode{Val: 3},
			},
		},
	}
	root = insertionSortList(root)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	pr := &ListNode{Val: math.MinInt32}

	tmp := head
	for tmp != nil {
		next := tmp.Next
		tmp.Next = nil

		t := pr
		for t.Next != nil && t.Next.Val < tmp.Val {
			t = t.Next
		}

		tmp.Next = t.Next
		t.Next = tmp
		tmp = next
	}

	return pr.Next
}
