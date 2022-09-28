package main

import "fmt"

func main() {
	m := Constructor()
	m.Insert("apple", 3)
	fmt.Println(m.Sum("apple"))
	m.Insert("apple", 2)
	fmt.Println(m.Sum("ab"))
}

type Tire struct {
	key      byte
	value    int
	children map[byte]*Tire
}

type MapSum struct {
	t map[byte]*Tire
}

func Constructor() MapSum {
	return MapSum{t: make(map[byte]*Tire)}
}

func (m *MapSum) Insert(key string, val int) {
	t := m.t
	for i := 0; i < len(key); i++ {
		if t[key[i]] == nil {
			t[key[i]] = &Tire{
				key:      key[i],
				children: make(map[byte]*Tire),
			}
		}
		if i == len(key)-1 {
			t[key[i]].value = val
		}

		t = t[key[i]].children
	}
}

func (m *MapSum) Sum(prefix string) int {
	var r int
	t := m.t
	for i := 0; i < len(prefix); i++ {
		if t[prefix[i]] == nil {
			return 0
		}

		if i == len(prefix)-1 {
			r += t[prefix[i]].value
		}

		t = t[prefix[i]].children
	}

	var dfs func(map[byte]*Tire)
	dfs = func(mp map[byte]*Tire) {
		if mp == nil {
			return
		}

		for _, v := range mp {
			if v != nil {
				r += v.value
				dfs(v.children)
			}
		}
	}

	dfs(t)

	return r
}
