package main

import (
	"bytes"
	"fmt"
	"sort"
)

func main() {
	fmt.Println(sortSentence("is2 sentence4 This1 a3"))
}

type Node struct {
	I byte
	V string
}

func sortSentence(s string) string {
	nl := make([]Node, 0)
	for i := 0; i < len(s); {
		if s[i] == ' ' {
			i++
			continue
		}
		j := i + 1
		for ; j < len(s); j++ {
			if s[j] == ' ' {
				break
			}
		}

		nl = append(nl, Node{
			I: s[j-1],
			V: s[i : j-1],
		})
		i = j
	}

	sort.Slice(nl, func(i, j int) bool {
		return nl[i].I < nl[j].I
	})

	r := bytes.NewBuffer([]byte{})
	for _, v := range nl {
		r.WriteString(v.V)
		r.WriteString(" ")
	}

	return r.String()[:r.Len()-1]
}
