package main

import "fmt"

func main() {
	c := Constructor()

	c.AddAtIndex(1, 0)
	fmt.Println(c.Get(0))
}

type MyLinkedList struct {
	Len  int
	Head *Node
	Tail *Node
}

type Node struct {
	Val  int
	Next *Node
	Pre  *Node
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}

func (m *MyLinkedList) Get(index int) int {
	if index >= m.Len {
		return -1
	}

	node := m.Index(index)
	return node.Val
}

func (m *MyLinkedList) AddAtHead(val int) {
	node := &Node{
		Val:  val,
		Next: m.Head,
		Pre:  m.Tail,
	}

	if m.Len == 0 {
		m.Len++
		m.Head, m.Tail = node, node
		m.Head.Pre, m.Tail.Next = node, node
		return
	}

	m.Len++
	m.Head.Pre, m.Tail.Next = node, node
	m.Head = node
}

func (m *MyLinkedList) AddAtTail(val int) {
	node := &Node{
		Val:  val,
		Next: m.Head,
		Pre:  m.Tail,
	}

	if m.Len == 0 {
		m.Len++
		m.Head, m.Tail = node, node
		m.Head.Pre, m.Tail.Next = node, node
		return
	}

	m.Len++
	m.Head.Pre, m.Tail.Next = node, node
	m.Tail = node
}

func (m *MyLinkedList) AddAtIndex(index int, val int) {
	if index > m.Len {
		return
	}

	if index == 0 {
		m.AddAtHead(val)
		return
	} else if index == m.Len {
		m.AddAtTail(val)
		return
	}

	indexNode := m.Index(index)
	node := &Node{
		Val:  val,
		Next: indexNode,
		Pre:  indexNode.Pre,
	}
	indexNode.Pre.Next, indexNode.Pre = node, node

	m.Len++
}

func (m *MyLinkedList) DeleteAtIndex(index int) {
	if index >= m.Len {
		return
	}

	if m.Len == 1 {
		m.Len = 0
		m.Head = nil
		m.Tail = nil
		return
	}

	if index == 0 {
		m.Head = m.Head.Next
	} else if index == m.Len-1 {
		m.Tail = m.Tail.Pre
	}

	indexNode := m.Index(index)
	indexNode.Pre.Next, indexNode.Next.Pre = indexNode.Next, indexNode.Pre

	m.Len--
}

func (m *MyLinkedList) Index(index int) *Node {
	r := m.Head
	for i := 0; i < index; i++ {
		r = r.Next
	}
	return r

}
