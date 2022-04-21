package main

import "fmt"

func main() {
	w := Constructor()
	w.AddWord("bad")

	fmt.Println(w.Search("b.."))
}

type Node struct {
	c  byte
	is bool
	m  map[byte]*Node
}

func (n *Node) Search(w []byte) bool {
	if n.c == '/' {
		for _, v := range n.m {
			if v.Search(w) {
				return true
			}
		}
		return false
	}

	if len(w) == 1 {
		return (w[0] == n.c || w[0] == '.') && n.is
	}

	if w[0] != '.' && w[0] != n.c {
		return false
	}

	if w[0] == '.' || (w[0] == n.c && w[1] == '.') {
		for _, v := range n.m {
			if v.Search(w[1:]) {
				return true
			}
		}
		return false
	}

	v := n.m[w[1]]
	if v == nil {
		return false
	}

	return v.Search(w[1:])
}

func (n *Node) AddWord(w []byte) {
	v := n.m[w[0]]
	if v == nil {
		v = &Node{c: w[0], m: make(map[byte]*Node)}
		n.m[w[0]] = v
	}

	if len(w) == 1 {
		v.is = true
		return
	}

	v.AddWord(w[1:])
}

type WordDictionary struct {
	m *Node
}

func Constructor() WordDictionary {
	return WordDictionary{m: &Node{c: '/', m: make(map[byte]*Node)}}
}

func (w *WordDictionary) AddWord(word string) {
	w.m.AddWord([]byte(word))
}

func (w *WordDictionary) Search(word string) bool {
	return w.m.Search([]byte(word))
}
