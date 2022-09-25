package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(sortPeople([]string{"Alice", "Bob", "Bob"}, []int{155, 185, 150}))
}

type Node struct {
	name   string
	height int
}

func sortPeople(names []string, heights []int) []string {
	node := make([]Node, len(names))
	for i, n := range names {
		node[i] = Node{
			name:   n,
			height: heights[i],
		}
	}

	sort.Slice(node, func(i, j int) bool {
		return node[i].height > node[j].height
	})

	r := make([]string, len(names))
	for i := range node {
		r[i] = node[i].name
	}

	return r
}
