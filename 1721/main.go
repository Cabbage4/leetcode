package main

import "fmt"

func main() {
	//[7,9,6,6,7,8,3,0,9,5]
	//	5
	head := &ListNode{
		Val: 7,
		Next: &ListNode{
			Val: 9,
			Next: &ListNode{
				Val: 6,
				Next: &ListNode{
					Val: 6,
					Next: &ListNode{
						Val: 7,
						Next: &ListNode{
							Val: 8,
							Next: &ListNode{
								Val: 3,
								Next: &ListNode{
									Val: 0,
									Next: &ListNode{
										Val:  9,
										Next: &ListNode{Val: 5},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	head = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  2,
			Next: &ListNode{Val: 3},
		},
	}

	head = swapNodes(head, 2)
	fmt.Println(head)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapNodes(head *ListNode, k int) *ListNode {
	tmp := head
	for i := 0; i < k-1; i++ {
		tmp = tmp.Next
	}

	first := tmp
	second := head
	for tmp.Next != nil {
		second = second.Next
		tmp = tmp.Next
	}

	first.Val, second.Val = second.Val, first.Val

	return head
}
