package main

import (
	"fmt"
	"sort"
)

func main() {
	d := []int{2, 4, 6, 8, 10}
	p := []int{10, 20, 30, 40, 50}
	w := []int{4, 5, 6, 7}
	fmt.Println(maxProfitAssignment(d, p, w))
}
func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	list := make([][]int, len(difficulty))
	for i := range difficulty {
		list[i] = []int{difficulty[i], profit[i]}
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i][1] > list[j][1]
	})

	sort.Sort(sort.Reverse(sort.IntSlice(worker)))

	var r int
	for i := 0; i < len(list); i++ {
		if len(worker) == 0 {
			break
		}

		for len(worker) > 0 && worker[0] >= list[i][0] {
			r += list[i][1]
			worker = worker[1:]
		}
	}

	return r
}
