package main

func main() {

}

type MyStack struct {
	data []int
}

func Constructor() MyStack {
	return MyStack{data: make([]int, 0)}
}

func (m *MyStack) Push(x int) {
	m.data = append(m.data, x)
}

func (m *MyStack) Pop() int {
	res := m.data[len(m.data)-1]
	m.data = m.data[:len(m.data)-1]
	return res
}

func (m *MyStack) Top() int {
	return m.data[len(m.data)-1]
}

func (m *MyStack) Empty() bool {
	return len(m.data) == 0
}
