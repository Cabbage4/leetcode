package main

import "fmt"

func main() {
	_1 := &Node{Val: 1}
	_2 := &Node{Val: 2}
	_3 := &Node{Val: 3}
	_4 := &Node{Val: 4}
	_5 := &Node{Val: 5}
	_6 := &Node{Val: 6}
	_7 := &Node{Val: 7}
	_8 := &Node{Val: 8}
	//_9 := &Node{Val: 9}
	//_10 := &Node{Val: 10}
	_11 := &Node{Val: 11}
	_12 := &Node{Val: 12}

	_1.Next = _2
	_2.Prev = _1

	_2.Next = _3
	_3.Prev = _2

	_3.Next = _4
	_4.Prev = _3

	_4.Next = _5
	_5.Prev = _4

	_5.Next = _6
	_6.Prev = _5

	//
	_7.Next = _8
	_8.Prev = _7

	//_8.Next = _9
	//_9.Prev = _8
	//
	//_9.Next = _10
	//_10.Prev = _9

	//
	_11.Next = _12
	_12.Prev = _11

	//*
	_3.Child = _7
	_8.Child = _11

	r := flatten(_1)
	tmp := r
	for tmp != nil {
		fmt.Printf("%d\t", tmp.Val)
		tmp = tmp.Next
	}
}

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func dfs(node *Node) (last *Node) {
	cur := node
	for cur != nil {
		next := cur.Next
		if cur.Child != nil {
			childLast := dfs(cur.Child)

			next = cur.Next
			cur.Next = cur.Child
			cur.Child.Prev = cur

			if next != nil {
				childLast.Next = next
				next.Prev = childLast
			}

			cur.Child = nil
			last = childLast
		} else {
			last = cur
		}
		cur = next
	}
	return
}

func flatten(root *Node) *Node {
	dfs(root)
	return root
}
