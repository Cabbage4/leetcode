package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	result := head
	fast := head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		result = result.Next
	}

	return result
}
