package main

import "container/heap"

func main() {
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
	r := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return r
}

func minimumSum(num int) int {
	h := Heap(make([]int, 0))
	heap.Init(&h)

	for num > 0 {
		t := num % 10
		if t != 0 {
			heap.Push(&h, t)
		}
		num = num / 10
	}

	t1 := 0
	t2 := 0
	flag := true
	for h.Len() != 0 {
		d := heap.Pop(&h).(int)
		if flag {
			t1 = t1*10 + d
		} else {
			t2 = t2*10 + d
		}

		flag = !flag
	}
	return t1 + t2
}
