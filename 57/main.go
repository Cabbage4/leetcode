package main

import (
	"fmt"
)

func main() {
	fmt.Println(insert([][]int{{1, 3}, {6, 9}}, []int{4, 5}))
	fmt.Println(insert([][]int{{1, 5}}, []int{0, 0}))
	fmt.Println(insert([][]int{{1, 5}}, []int{2, 7}))
}

func insert(ol [][]int, nl []int) [][]int {
	if len(ol) == 0 {
		return [][]int{nl}
	}

	r := make([][]int, 0)
	index := 0
	for index < len(ol) && ol[index][1] < nl[0] {
		r = append(r, ol[index])
		index++
	}

	for index < len(ol) && ol[index][0] <= nl[1] {
		if ol[index][0] < nl[0] {
			nl[0] = ol[index][0]
		}
		if ol[index][1] > nl[1] {
			nl[1] = ol[index][1]
		}
		index++
	}

	r = append(r, nl)
	for index < len(ol) {
		r = append(r, ol[index])
		index++
	}

	return r
}
