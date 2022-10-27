package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(carFleet(20, []int{6, 2, 17}, []int{3, 9, 2}))
	fmt.Println(carFleet(12, []int{10, 8, 0, 5, 3}, []int{2, 4, 1, 1, 3}))
}

func carFleet(target int, position []int, speed []int) int {
	type Node struct {
		p int
		t float64
	}

	list := make([]Node, len(position))
	for i := range position {
		list[i] = Node{
			p: position[i],
			t: float64(target-position[i]) / float64(speed[i]),
		}
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i].p < list[j].p
	})

	var r int
	for i := len(position) - 1; i > 0; i-- {
		if list[i].t < list[i-1].t {
			r++
		} else {
			list[i-1].p, list[i-1].t = list[i].p, list[i].t
		}
	}

	return r + 1
}
