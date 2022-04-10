package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maxSubsequence([]int{2, 1, 3, 3}, 2))
	fmt.Println(maxSubsequence([]int{-1, -2, 3, 4}, 3))
}

type Heap [][]int

func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Less(i, j int) bool {
	return h[i][1] < h[j][1]
}

func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

func (h *Heap) Pop() interface{} {
	r := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return r
}

func maxSubsequence(nums []int, k int) []int {
	h := Heap(make([][]int, 0))
	heap.Init(&h)

	for i, v := range nums {
		if h.Len() < k {
			heap.Push(&h, []int{i, v})
			continue
		}

		if v > h[0][1] {
			heap.Pop(&h)
			heap.Push(&h, []int{i, v})
		}
	}

	r := make([]int, 0)
	sort.Slice(h, func(i, j int) bool {
		return h[i][0] < h[j][0]
	})

	for _, hh := range h {
		r = append(r, hh[1])
	}

	return r
}
