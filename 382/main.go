package main

import (
	"fmt"
	"math/rand"
)

func main() {
	linkNode := &ListNode{
		Val: 10,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 10,
				Next: &ListNode{
					Val: 20,
					Next: &ListNode{
						Val:  100,
						Next: nil,
					},
				},
			},
		},
	}

	mp := make(map[int]int)
	s := Constructor(linkNode)
	for i := 0; i < 10001; i++ {
		mp[s.GetRandom()]++
	}
	fmt.Println(mp)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

type Solution struct {
	list []int
}

func Constructor(head *ListNode) Solution {
	p := head
	list := make([]int, 0)
	for p != nil {
		list = append(list, p.Val)
		p = p.Next
	}

	return Solution{list: list}
}

func (s *Solution) GetRandom() int {
	return s.list[rand.Intn(len(s.list))]
}
