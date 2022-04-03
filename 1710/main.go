package main

import "container/heap"

func main() {
}

type Heap [][]int

func (h *Heap) Len() int {
	return len(*h)
}

func (h *Heap) Less(i, j int) bool {
	return (*h)[i][1] > (*h)[j][1]
}

func (h *Heap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

func (h *Heap) Pop() interface{} {
	r := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return r
}

func maximumUnits(boxTypes [][]int, truckSize int) int {
	h := Heap(boxTypes)
	heap.Init(&h)

	r := 0
	cs := 0
	for cs < truckSize && h.Len() > 0 {
		tmp := heap.Pop(&h).([]int)
		if tmp[0] <= truckSize-cs {
			r += tmp[0] * tmp[1]
			cs += tmp[0]
			continue
		}

		r += (truckSize - cs) * tmp[1]
		break
	}

	return r
}
