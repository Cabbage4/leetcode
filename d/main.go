package main

import (
	"fmt"
)

func main() {
	fmt.Println(sumPrefixScores([]string{"abc", "ab", "bc", "b"}))
}

type Tie struct {
	val      byte
	count    int
	children map[byte]*Tie
}

func sumPrefixScores(words []string) []int {
	t := make(map[byte]*Tie)
	for i := 0; i < len(words); i++ {
		s := words[i]

		tmpT, ok := t[s[0]]
		if !ok {
			tmpT = &Tie{
				val:      s[0],
				children: make(map[byte]*Tie),
			}
			t[s[0]] = tmpT
		}

		tmpT.count++

		for j := 1; j < len(s); j++ {
			v, ok := tmpT.children[s[j]]
			if !ok {
				v = &Tie{
					val:      s[j],
					children: make(map[byte]*Tie),
				}
				tmpT.children[s[j]] = v
			}
			v.count++
			tmpT = v
		}
	}

	r := make([]int, len(words))
	for i := 0; i < len(words); i++ {
		s := words[i]
		tmpT := t[s[0]]
		r[i] += tmpT.count

		for j := 1; j < len(s); j++ {
			tmpT = tmpT.children[s[j]]
			r[i] += tmpT.count
		}
	}
	return r
}
