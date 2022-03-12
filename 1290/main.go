package main

func main() {
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	r := 0
	t := head
	for t != nil {
		r = r<<1 | t.Val
		t = t.Next
	}
	return r
}
