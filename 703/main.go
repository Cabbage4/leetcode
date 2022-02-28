package main

import (
	"container/heap"
	"fmt"
)

func main() {
	c := Constructor(3, []int{4, 5, 8, 2})
	fmt.Println(c.Add(3) == 4)
	fmt.Println(c.Add(5) == 5)
	fmt.Println(c.Add(10) == 5)
	fmt.Println(c.Add(9) == 8)
	fmt.Println(c.Add(4) == 8)
}

type Heap []int

func (h *Heap) Len() int {
	return len(*h)
}
func (h *Heap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}
func (h *Heap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}
func (h *Heap) Pop() interface{} {
	r := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return r
}
func (h *Heap) Push(v interface{}) {
	*h = append(*h, v.(int))
}

type KthLargest struct {
	minHeap *Heap
	size    int
}

func Constructor(k int, nums []int) KthLargest {
	h := Heap(make([]int, 0))
	heap.Init(&h)

	for _, v := range nums {
		if h.Len() < k {
			heap.Push(&h, v)
			continue
		}

		if h[0] < v {
			heap.Pop(&h)
			heap.Push(&h, v)
			continue
		}
	}

	return KthLargest{
		minHeap: &h,
		size:    k,
	}
}

func (k *KthLargest) Add(val int) int {
	h := k.minHeap

	if h.Len() < k.size {
		heap.Push(h, val)
		return (*h)[0]
	}

	if val <= (*h)[0] {
		return (*h)[0]
	}

	(*h)[0] = val
	heap.Fix(h, 0)
	return (*h)[0]
}
