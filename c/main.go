package main

func main() {
	//5,2,13,3,8
	head := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 13,
				Next: &ListNode{
					Val:  3,
					Next: &ListNode{Val: 8},
				},
			},
		},
	}

	removeNodes(head)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNodes(head *ListNode) *ListNode {
	list := toList(head)
	dp := make([]int, len(list))

	dp[len(list)-1] = -1
	for i := len(list) - 2; i >= 0; i-- {
		k := i + 1

		for dp[k] != -1 && list[i] >= list[k] {
			k = dp[k]
		}

		if list[i] >= list[k] {
			dp[i] = -1
		} else {
			dp[i] = k
		}
	}

	r := new(ListNode)
	t := r
	for i := range dp {
		if dp[i] != -1 {
			continue
		}

		t.Next = &ListNode{
			Val: list[i],
		}
		t = t.Next
	}

	return r.Next
}

func toList(node *ListNode) []int {
	var r []int
	for node != nil {
		r = append(r, node.Val)
		node = node.Next
	}
	return r
}
