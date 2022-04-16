package main

import "fmt"

func main() {
	root := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val:  1,
			Next: nil,
		},
	}

	fmt.Println(partition(root, 2))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	hlr := new(ListNode)
	hhr := new(ListNode)

	tlr := hlr
	thr := hhr

	tmp := head
	for tmp != nil {
		next := tmp.Next
		tmp.Next = nil
		if tmp.Val < x {
			tlr.Next = tmp
			tlr = tlr.Next
		} else {
			thr.Next = tmp
			thr = thr.Next
		}
		tmp = next
	}

	tlr.Next = hhr.Next

	return hlr.Next
}
