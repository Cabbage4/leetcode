package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(topKFrequent([]string{"i", "love", "leetcode", "i", "love", "coding"}, 1))
	fmt.Println(topKFrequent([]string{"i", "love", "leetcode", "i", "love", "coding"}, 1))
}

type Node struct {
	count int
	value string
}

type Heap []Node

func (h *Heap) Len() int {
	return len(*h)
}

func (h *Heap) Less(i, j int) bool {
	if (*h)[i].count == (*h)[j].count {
		return (*h)[i].value < (*h)[j].value
	}
	return (*h)[i].count > (*h)[j].count
}

func (h *Heap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *Heap) Push(v interface{}) {
	*h = append(*h, v.(Node))
}

func (h *Heap) Pop() interface{} {
	r := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return r
}

func topKFrequent(words []string, k int) []string {
	mp := make(map[string]int)
	for _, w := range words {
		mp[w]++
	}

	h := Heap(make([]Node, 0))
	heap.Init(&h)
	for k, v := range mp {
		heap.Push(&h, Node{
			count: v,
			value: k,
		})
	}

	var r []string
	for k > 0 {
		k--
		v := heap.Pop(&h).(Node)
		r = append(r, v.value)
	}

	return r
}
