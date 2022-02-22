package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	res := &ListNode{}
	resp := res
	p := head.Next
	if head.Val != val {
		head.Next = nil

		resp.Next = head
		resp = resp.Next
	}

	for p != nil {
		if p.Val != val {
			resp.Next = p
			resp = resp.Next

			p = p.Next
			resp.Next = nil
			continue
		}

		p = p.Next
	}

	return res.Next
}
