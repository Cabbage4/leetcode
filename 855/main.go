package main

import (
	"fmt"
	"sort"
)

func main() {
	c := Constructor(10)
	for i := 0; i < 3; i++ {
		fmt.Print(c.Seat(), " ")
	}
	c.Leave(0)
	c.Leave(4)
	for i := 0; i < 9; i++ {
		fmt.Print(c.Seat(), " ")
	}
	//
	//for i := 0; i < 6; i++ {
	//	fmt.Print(c.Seat(), " ")
	//}
}

type ExamRoom struct {
	n    int
	list []int
}

func Constructor(n int) ExamRoom {
	return ExamRoom{
		n:    n,
		list: make([]int, 0),
	}
}

func (e *ExamRoom) Seat() int {
	if len(e.list) == 0 {
		e.list = append(e.list, 0)
		return 0
	}

	left := 0
	right := e.list[0]
	for i := 1; i < len(e.list); i++ {
		ln := e.list[i] - e.list[i-1] - 1
		mln := right - left
		if ln > mln && !(mln%2 != 0 && ln-mln == 1) {
			left = e.list[i-1] + 1
			right = e.list[i]
		}
	}

	if e.list[len(e.list)-1] != e.n-1 {
		ln := e.n - e.list[len(e.list)-1] - 1
		mln := right - left
		if ln > mln {
			e.list = append(e.list, e.n-1)
			sort.Ints(e.list)
			return e.n - 1
		}
	}

	if e.list[0] != 0 {
		ln := e.list[0]
		mln := right - left
		if ln >= mln {
			e.list = append(e.list, 0)
			sort.Ints(e.list)
			return 0
		}
	}

	index := (left + right - 1) / 2

	e.list = append(e.list, index)
	sort.Ints(e.list)
	return index
}

func (e *ExamRoom) Leave(p int) {
	i := sort.SearchInts(e.list, p)
	e.list = append(e.list[:i], e.list[i+1:]...)
}
