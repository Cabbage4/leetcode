package main

func main() {
	root := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			//Next: &ListNode{
			//	Val: 3,
			//	Next: &ListNode{
			//		Val: 4,
			//		Next: &ListNode{
			//			Val: 5,
			//			Next: &ListNode{
			//				Val:  6,
			//				Next: nil,
			//			},
			//		},
			//	},
			//},
		},
	}

	reorderList(root)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	mid := head
	qTmp := head
	for qTmp.Next != nil && qTmp.Next.Next != nil {
		mid = mid.Next
		qTmp = qTmp.Next.Next
	}

	midHead := new(ListNode)
	tmp := mid.Next
	mid.Next = nil
	for tmp != nil {
		next := tmp.Next

		tmp.Next = midHead.Next
		midHead.Next = tmp

		tmp = next
	}

	midHead = midHead.Next
	tmp2 := head
	for tmp2 != nil && midHead != nil {
		next := tmp2.Next
		nextMidHead := midHead.Next

		tmp2.Next = midHead
		midHead.Next = next

		tmp2 = next
		midHead = nextMidHead
	}

	return
}
