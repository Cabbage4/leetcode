package main

import (
	"fmt"
	"math/rand"
)

func main() {
	c := Constructor(10, 10)
	for i := 0; i < 10; i++ {
		fmt.Println(c.Flip())
	}
}

type Solution struct {
	n int
	m int

	mp    map[int]int
	total int
}

func Constructor(m, n int) Solution {
	return Solution{
		m: m,
		n: n,

		mp:    make(map[int]int),
		total: m * n,
	}
}

func (s *Solution) Flip() []int {
	r := make([]int, 2)

	x := rand.Intn(s.total)
	s.total--
	if v, ok := s.mp[x]; ok {
		r[0], r[1] = v/s.n, v%s.n
	} else {
		r[0], r[1] = x/s.n, x%s.n
	}

	if v, ok := s.mp[s.total]; ok {
		s.mp[x] = v
	} else {
		s.mp[x] = s.total
	}

	return r
}

func (s *Solution) Reset() {
	s.total = s.n * s.m
	s.mp = make(map[int]int)
}
