package main

import "fmt"

func main() {
	root := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
				},
			},
		},
	}

	result := swapPairs(root)
	fmt.Println(result.Val)
	fmt.Println(result.Next.Val)
	fmt.Println(result.Next.Next.Val)
	fmt.Println(result.Next.Next.Next.Val)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	t := head
	rp := new(ListNode)
	tr := rp

	for t != nil && t.Next != nil {
		k := t.Next.Next

		t1, t2 := t, t.Next
		t1.Next, t2.Next = nil, t1

		tr.Next = t2
		tr = t1
		t = k
	}

	if t != nil {
		tr.Next = t
	}

	return rp.Next
}
