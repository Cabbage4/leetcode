package main

import "fmt"

func main() {
	fmt.Println(minimumLengthEncoding([]string{"t"}))
	fmt.Println(minimumLengthEncoding([]string{"time", "me", "bell"}))
}

type node struct {
	c     byte
	flag  bool
	child map[byte]*node
}

type tie map[byte]*node

func (t tie) add(s string) {
	if len(s) == 0 {
		return
	}

	tmp := t[s[0]]
	if tmp == nil {
		tmp = &node{
			c:     s[0],
			child: make(map[byte]*node),
		}
		t[s[0]] = tmp
	}

	for i := 1; i < len(s); i++ {
		tTmp := tmp.child[s[i]]
		if tTmp == nil {
			tTmp = &node{
				c:     s[i],
				child: make(map[byte]*node),
			}
			tmp.child[s[i]] = tTmp
		}

		tmp = tTmp
	}

	tmp.flag = true
}

func (t tie) count() int {
	var r int
	var dfs func(*node, int)
	dfs = func(nd *node, step int) {
		if nd == nil {
			return
		}

		if nd.flag && len(nd.child) == 0 {
			r += step + 1
		}

		for _, v := range nd.child {
			dfs(v, step+1)
		}
	}

	for _, v := range t {
		dfs(v, 1)
	}
	return r
}

func minimumLengthEncoding(words []string) int {
	reverse := func(s string) string {
		r := []byte(s)
		for i := 0; i < len(r)/2; i++ {
			r[i], r[len(r)-1-i] = r[len(r)-1-i], r[i]
		}
		return string(r)
	}

	t := tie(make(map[byte]*node))
	for _, v := range words {
		t.add(reverse(v))
	}

	return t.count()
}
