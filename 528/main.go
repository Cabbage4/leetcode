package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	c := Constructor([]int{1})
	for i := 0; i < 10; i++ {
		fmt.Println(c.PickIndex())
	}
}

type Solution struct {
	list []int
}

func Constructor(w []int) Solution {
	f := make([]int, len(w))
	f[0] = w[0]
	for i := 1; i < len(w); i++ {
		f[i] = w[i] + f[i-1]
	}
	return Solution{list: f}
}

func (s *Solution) PickIndex() int {
	index := rand.Intn(s.list[len(s.list)-1]) + 1
	return sort.SearchInts(s.list, index)
}
