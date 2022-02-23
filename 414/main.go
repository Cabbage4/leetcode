package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(thirdMax([]int{1, 2}))
	fmt.Println(thirdMax([]int{4, 2, 3, 1}))
}

type Heap []int

func (h *Heap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *Heap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *Heap) Len() int {
	return len(*h)
}

func (h *Heap) Push(v interface{}) {
	*h = append(*h, v.(int))
}

func (h *Heap) Pop() interface{} {
	v := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return v
}

func thirdMax(nums []int) int {
	mp := make(map[int]bool)
	max := -1
	for _, v := range nums {
		if v > max {
			max = v
		}
		mp[v] = true
	}

	if len(mp) < 3 {
		return max
	}

	h := new(Heap)
	heap.Init(h)

	for k := range mp {
		if mp[k] {
			if h.Len() < 3 {
				heap.Push(h, k)
			} else {
				v := heap.Pop(h).(int)
				if k > v {
					heap.Push(h, k)
				} else {
					heap.Push(h, v)
				}
			}
		}
	}

	return heap.Pop(h).(int)
}
