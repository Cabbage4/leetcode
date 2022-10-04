package main

import "fmt"

func main() {
	m := Constructor()
	fmt.Println(m.Book(10, 20))
	fmt.Println(m.Book(50, 60))
	fmt.Println(m.Book(10, 40))
	fmt.Println(m.Book(5, 15))
	fmt.Println(m.Book(5, 10))
	fmt.Println(m.Book(25, 55))
}

type pair struct {
	first  int
	second int
}

type MyCalendarTwo map[int]pair

func Constructor() MyCalendarTwo {
	return MyCalendarTwo{}
}

func (m MyCalendarTwo) Book(s int, e int) bool {
	m.update(s, e-1, 0, 1e9, 1, 1)
	if m[1].first > 2 {
		m.update(s, e-1, 0, 1e9, -1, 1)
		return false
	}
	return true
}

func (m MyCalendarTwo) update(s, e, ls, le, value, label int) {
	if s > le || e < ls {
		return
	}

	if s <= ls && le <= e {
		p := m[label]
		p.second += value
		p.first += value
		m[label] = p
		return
	}
	mid := (ls + le) >> 1
	m.update(s, e, ls, mid, value, 2*label)
	m.update(s, e, mid+1, le, value, 2*label+1)
	p := m[label]
	p.first = p.second + max(m[2*label].first, m[2*label+1].first)
	m[label] = p
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
