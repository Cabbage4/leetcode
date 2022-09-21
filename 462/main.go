package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(minMoves2([]int{1, 10, 2, 9}))
	fmt.Println(minMoves2([]int{1, 2, 3}))
}

type Heap []int

func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *Heap) Pop() interface{} {
	r := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return r
}

func minMoves2(nums []int) int {
	heapSize := len(nums) / 2
	if len(nums)%2 != 0 {
		heapSize++
	}

	h := Heap([]int{})
	heap.Init(&h)
	for _, v := range nums {
		if h.Len() < heapSize {
			heap.Push(&h, v)
			continue
		}

		if h[0] < v {
			heap.Pop(&h)
			heap.Push(&h, v)
		}
	}

	target := heap.Pop(&h).(int)

	var r int
	for _, v := range nums {
		r += abs(v, target)
	}

	return r
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
