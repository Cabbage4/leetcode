package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	fmt.Println(totalCost([]int{57, 33, 26, 76, 14, 67, 24, 90, 72, 37, 30}, 11, 2))
	fmt.Println(totalCost([]int{28, 35, 21, 13, 21, 72, 35, 52, 74, 92, 25, 65, 77, 1, 73, 32, 43, 68, 8, 100, 84, 80, 14, 88, 42, 53, 98, 69, 64, 40, 60, 23, 99, 83, 5, 21, 76, 34}, 32, 12))
	fmt.Println(totalCost([]int{1, 2, 4, 1}, 3, 3))
	fmt.Println(totalCost([]int{17, 12, 10, 2, 7, 2, 11, 20, 8}, 3, 8))
}

type Heap struct {
	sort.IntSlice
}

func (h *Heap) Pop() interface{} {
	r := h.IntSlice[len(h.IntSlice)-1]
	h.IntSlice = h.IntSlice[:len(h.IntSlice)-1]
	return r
}

func (h *Heap) Push(v interface{}) {
	h.IntSlice = append(h.IntSlice, v.(int))
}

func totalCost(costs []int, k int, candidates int) int64 {
	var r int64

	if len(costs) <= 2*candidates {
		sort.Ints(costs)
		for i := 0; i < k; i++ {
			r += int64(costs[i])
		}
		return r
	}

	leftHeap := new(Heap)
	rightHeap := new(Heap)
	for i := 0; i < candidates; i++ {
		heap.Push(leftHeap, costs[i])
		heap.Push(rightHeap, costs[len(costs)-1-i])
	}

	leftIndex := candidates - 1
	rightIndex := len(costs) - candidates
	for k > 0 {
		if leftIndex >= rightIndex-1 {
			if leftHeap.Len() == 0 {
				r += int64(heap.Pop(rightHeap).(int))
			} else if rightHeap.Len() == 0 {
				r += int64(heap.Pop(leftHeap).(int))
			} else if leftHeap.IntSlice[0] <= rightHeap.IntSlice[0] {
				r += int64(heap.Pop(leftHeap).(int))
			} else {
				r += int64(heap.Pop(rightHeap).(int))
			}

			k--
			continue
		}

		if leftHeap.IntSlice[0] <= rightHeap.IntSlice[0] {
			r += int64(heap.Pop(leftHeap).(int))
			heap.Push(leftHeap, costs[leftIndex+1])
			leftIndex++
		} else {
			r += int64(heap.Pop(rightHeap).(int))
			heap.Push(rightHeap, costs[rightIndex-1])
			rightIndex--
		}

		k--
	}

	return r
}
