package main

func main() {

}

type Node struct {
	val  int
	next *Node
}

type MyHashSet struct {
	table []*Node
	size  int
	cap   int
}

func Constructor() MyHashSet {
	m := MyHashSet{
		table: make([]*Node, 0),
		size:  0,
		cap:   16,
	}

	return m
}

func (m *MyHashSet) Add(val int) {
	key := hash(val, m.cap-1)
	nd := m.table[key]

	if nd == nil {
		m.table[key] = &Node{val: val}
		if float64(len(m.table))/float64(m.cap) >= 0.75 {
			m.cap = m.cap * 2
			m.table = resize(m.table, m.cap)
		}
		return
	}

	if nd.val == val {
		return
	}

	tnd := nd
	for tnd.next != nil {
		if tnd.next.val == val {
			return
		}

		tnd = tnd.next
	}

	tnd.next = &Node{val: val}
	m.size++
}

func (m *MyHashSet) Remove(val int) {
	key := hash(val, m.cap-1)
	nd := m.table[key]
	if nd == nil {
		return
	}

	if nd.val == val {
		m.table[key] = nd.next
		return
	}

	tnd := nd
	for tnd.next != nil {
		if tnd.next.val == val {
			tnd.next = tnd.next.next
			return
		}

		tnd = tnd.next
	}
}

func (m *MyHashSet) Contains(val int) bool {
	key := hash(val, m.cap-1)
	nd := m.table[key]
	if nd == nil {
		return false
	}

	if nd.val == val {
		return true
	}

	tnd := nd.next
	for tnd != nil {
		if tnd.val == val {
			return true
		}

		tnd = tnd.next
	}
	return false
}

func hash(key, size int) int {
	return key & size
}

func resize(table []*Node, cap int) []*Node {
	nt := make([]*Node, cap)

	for _, n := range table {
		if n == nil {
			continue
		}

		tnd := n
		for tnd != nil {
			tnd.next = nil
			tmp := tnd.next
		}
	}

	return nt
}
