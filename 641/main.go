package main

func main() {
}

type node struct {
	val  int
	next *node
	pre  *node
}

type MyCircularDeque struct {
	cap int
	len int

	head *node
	tail *node
}

func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		cap: k,
		len: 0,
	}
}

func (m *MyCircularDeque) InsertFront(value int) bool {
	if m.cap == m.len {
		return false
	}

	ne := &node{
		val:  value,
		next: m.head,
	}

	if m.len == 0 {
		m.head = ne
		m.tail = ne
		m.len++
		return true
	}

	m.head.pre, m.head = ne, ne
	m.len++
	return true
}

func (m *MyCircularDeque) InsertLast(value int) bool {
	if m.cap == m.len {
		return false
	}

	ne := &node{
		val: value,
		pre: m.tail,
	}

	if m.len == 0 {
		m.head = ne
		m.tail = ne
		m.len++
		return true
	}

	m.tail.next, m.tail = ne, ne
	m.len++
	return true
}

func (m *MyCircularDeque) DeleteFront() bool {
	if m.len == 0 {
		return false
	}

	if m.len == 1 {
		m.head, m.tail = nil, nil
	} else {
		m.head, m.head.pre = m.head.next, nil
	}

	m.len--
	return true
}

func (m *MyCircularDeque) DeleteLast() bool {
	if m.len == 0 {
		return false
	}

	if m.len == 1 {
		m.head, m.tail = nil, nil
	} else {
		m.tail, m.tail.next = m.tail.pre, nil
	}

	m.len--
	return true
}

func (m *MyCircularDeque) GetFront() int {
	if m.len == 0 {
		return -1
	}
	return m.head.val
}

func (m *MyCircularDeque) GetRear() int {
	if m.len == 0 {
		return -1
	}
	return m.tail.val
}

func (m *MyCircularDeque) IsEmpty() bool {
	return m.len == 0
}

func (m *MyCircularDeque) IsFull() bool {
	return m.len == m.cap
}
