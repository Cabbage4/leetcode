package main

import "fmt"

func main() {
	fmt.Println(numComponents(&ListNode{Val: 0, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2}}}, []int{0, 2}))
	fmt.Println(numComponents(&ListNode{Val: 0}, []int{0}))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func numComponents(head *ListNode, nums []int) int {
	if head == nil {
		return 0
	}

	set := make(map[int]bool)
	for _, v := range nums {
		set[v] = true
	}

	r := 0
	cur := 0
	tmp := head

	for tmp != nil {
		if set[tmp.Val] {
			if cur == 0 {
				r++
			}
			cur++
		} else {
			cur = 0
		}

		tmp = tmp.Next
	}

	return r
}
