package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(lastStoneWeight([]int{2, 7, 4, 1, 8, 1}))
}

type Heap []int

func (h *Heap) Len() int {
	return len(*h)
}
func (h *Heap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *Heap) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}

func (h *Heap) Pop() interface{} {
	r := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return r
}
func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func lastStoneWeight(stones []int) int {
	h := Heap(stones)
	heap.Init(&h)

	for h.Len() > 1 {
		y := heap.Pop(&h).(int)
		x := heap.Pop(&h).(int)
		if y-x != 0 {
			heap.Push(&h, y-x)
		}
	}

	if h.Len() == 0 {
		return 0
	}

	return h.Pop().(int)
}
