package main

func main() {
	r := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			//Next: &ListNode{
			//	Val: 3,
			//	Next: &ListNode{
			//		Val: 4,
			//		Next: &ListNode{
			//			Val:  5,
			//			Next: nil,
			//		},
			//	},
			//},
		},
	}

	reverseBetween(r, 1, 2)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == right {
		return head
	}

	var lnp *ListNode
	var rnn *ListNode

	var ln *ListNode
	var rn *ListNode
	cn := 1

	tmp := head
	for tmp != nil {
		if cn+1 == left {
			lnp = tmp
		}
		if cn == left {
			ln = tmp
		}
		if cn == right {
			rn = tmp
		}
		if cn == right+1 {
			rnn = tmp
		}

		tmp = tmp.Next
		cn++
	}

	flag := false
	if lnp == nil {
		flag = true
		lnp = &ListNode{}
	}

	lnp.Next = nil
	rn.Next = nil

	tmp2 := ln
	for tmp2 != nil {
		next := tmp2.Next

		n := tmp2
		n.Next = lnp.Next
		lnp.Next = n

		tmp2 = next
	}
	ln.Next = rnn

	if flag {
		return lnp.Next
	}

	return head
}
