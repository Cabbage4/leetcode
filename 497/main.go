package main

import (
	"math/rand"
	"sort"
)

func main() {
}

type Solution struct {
	rects [][]int
	sum   []int
}

func Constructor(rects [][]int) Solution {
	sum := make([]int, len(rects)+1)
	for i := 0; i < len(rects); i++ {
		x1, y1, x2, y2 := rects[i][0], rects[i][1], rects[i][2], rects[i][3]
		sum[i+1] = sum[i] + (x2-x1+1)*(y2-y1+1)
	}

	return Solution{rects: rects, sum: sum}
}

func (s *Solution) Pick() []int {
	x := rand.Intn(s.sum[len(s.sum)-1])
	index := sort.SearchInts(s.sum, x+1) - 1
	sum := s.sum[index]
	x1, y1, y2 := s.rects[index][0], s.rects[index][1], s.rects[index][3]
	return []int{x1 + (x-sum)/(y2-y1+1), y1 + (x-sum)%(y2-y1+1)}
}
