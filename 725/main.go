package main

import "fmt"

func main() {
	data1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}

	fmt.Println(splitListToParts(data1, 3))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func splitListToParts(head *ListNode, k int) []*ListNode {
	t := head
	ln := 0
	for t != nil {
		ln++
		t = t.Next
	}

	if ln < k {
		var r []*ListNode
		t = head
		for t != nil {
			r = append(r, t)
			t, t.Next = t.Next, nil
		}
		for i := 0; i < k-ln; i++ {
			r = append(r, nil)
		}
		return r
	}

	mod := ln % k
	limit := ln / k

	var r []*ListNode
	t = head
	for t != nil {
		count := 1
		node := t
		for count < limit {
			t = t.Next
			count++
		}

		if mod > 0 {
			t = t.Next
			mod--
		}

		r = append(r, node)
		t, t.Next = t.Next, nil
	}

	return r
}
