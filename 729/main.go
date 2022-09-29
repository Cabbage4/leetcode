package main

import "fmt"

func main() {
	m := Constructor()
	//fmt.Println(m.Book(97, 100))
	//fmt.Println(m.Book(33, 51))
	fmt.Println(m.Book(89, 99))
	fmt.Println(m.Book(90, 91))
}

type MyCalendar struct {
	hasMp map[int]bool
	lazy  map[int]bool
}

func Constructor() MyCalendar {
	return MyCalendar{lazy: make(map[int]bool), hasMp: make(map[int]bool)}
}

func (m *MyCalendar) Book(s int, e int) bool {
	if m.query(s, e-1, 0, 1e9, 1) {
		return false
	}

	m.update(s, e-1, 0, 1e9, 1)
	return true
}

func (m *MyCalendar) update(s, e, ls, le, label int) {
	if s > le || e < ls {
		return
	}

	m.hasMp[label] = true

	if s <= ls && le <= e {
		m.lazy[label] = true
		return
	}

	mid := (ls + le) >> 1
	m.update(s, e, ls, mid, label*2)
	m.update(s, e, mid+1, le, label*2+1)
	if m.lazy[2*label] && m.lazy[2*label+1] {
		m.lazy[label] = true
	}
}

func (m *MyCalendar) query(s, e, ls, le, label int) bool {
	if s > le || e < ls {
		return false
	}

	if m.lazy[label] {
		return true
	}

	if s <= ls && le <= e {
		return m.hasMp[label]
	}

	mid := (ls + le) >> 1
	return m.query(s, e, ls, mid, label*2) || m.query(s, e, mid+1, le, label*2+1)
}
