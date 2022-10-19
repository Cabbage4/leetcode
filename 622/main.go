package main

func main() {
}

type MyCircularQueue struct {
	cap  int
	len  int
	data []int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		cap: k,
		len: 0,
	}
}

func (m *MyCircularQueue) EnQueue(value int) bool {
	if m.cap == m.len {
		return false
	}

	m.data = append(m.data, value)
	m.len++
	return true
}

func (m *MyCircularQueue) DeQueue() bool {
	if m.len == 0 {
		return false
	}

	m.data = m.data[1:]
	m.len--
	return true
}

func (m *MyCircularQueue) Front() int {
	if m.len == 0 {
		return -1
	}
	return m.data[0]
}

func (m *MyCircularQueue) Rear() int {
	if m.len == 0 {
		return -1
	}
	return m.data[len(m.data)-1]
}

func (m *MyCircularQueue) IsEmpty() bool {
	return m.len == 0
}

func (m *MyCircularQueue) IsFull() bool {
	return m.len == m.cap
}
