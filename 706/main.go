package main

import "fmt"

func main() {
	m := Constructor()
	m.Remove(27)
	m.Put(65, 65)
	m.Remove(19)
	m.Remove(0)
	fmt.Println(m.Get(18))
	m.Remove(3)
	m.Put(42, 0)
	fmt.Println(m.Get(19))
	m.Remove(42)
	m.Put(17, 90)
	m.Put(31, 76)
	m.Put(48, 71)
	m.Put(5, 50)
	m.Put(7, 68)
	m.Put(73, 74)
	m.Put(85, 18)
	m.Put(74, 95)
	m.Put(84, 82)
	m.Put(59, 29)
	m.Put(71, 71)
	m.Remove(42)
	m.Put(51, 40)
	m.Put(33, 76)
	fmt.Println(m.Get(17))
}

type Node struct {
	Key  int
	Val  int
	next *Node
}

type MyHashMap struct {
	table []*Node
	size  int
	cap   int
}

func Constructor() MyHashMap {
	return MyHashMap{
		table: make([]*Node, 16),
		size:  0,
		cap:   16,
	}
}

func (m *MyHashMap) Put(key int, value int) {
	if float64(m.size)/float64(m.cap) >= 0.75 {
		m.table = resize(m.table, m.cap*2)
		m.cap = m.cap * 2
	}

	ahash := hash(key, m.cap)
	tnd := m.table[ahash]
	if tnd == nil {
		m.table[ahash] = &Node{
			Key: key,
			Val: value,
		}
		m.size++
		return
	}

	if tnd.Key == key {
		tnd.Val = value
		return
	}

	for tnd.next != nil {
		if tnd.next.Key == key {
			tnd.next.Val = value
			return
		}
		tnd = tnd.next
	}

	tnd.next = &Node{
		Key: key,
		Val: value,
	}
	m.size++
}

func (m *MyHashMap) Get(key int) int {
	tnd := m.table[hash(key, m.cap)]
	if tnd == nil {
		return -1
	}

	for tnd != nil {
		if tnd.Key == key {
			return tnd.Val
		}
		tnd = tnd.next
	}

	return -1
}

func (m *MyHashMap) Remove(key int) {
	ahash := hash(key, m.cap)
	tnd := m.table[ahash]
	if tnd == nil {
		return
	}

	if tnd.Key == key {
		m.table[ahash] = tnd.next
		m.size--
		return
	}

	for tnd.next != nil {
		if tnd.next.Key == key {
			tnd.next = tnd.next.next
			m.size--
			return
		}

		tnd = tnd.next
	}
}

func hash(key int, mask int) int {
	return key & (mask - 1)
}

func resize(table []*Node, newCap int) []*Node {
	result := make([]*Node, newCap)
	for _, t := range table {
		if t == nil {
			continue
		}

		tnd := t
		for tnd != nil {
			aHash := hash(tnd.Key, newCap)
			y := result[aHash]
			if y == nil {
				result[aHash] = tnd
			} else {
				for y.next != nil {
					y = y.next
				}
				y.next = tnd
			}

			tmp := tnd.next
			tnd.next = nil
			tnd = tmp
		}
	}

	return result
}
