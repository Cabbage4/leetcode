package main

import (
	"fmt"
)

func main() {
	var list []int
	list = []int{11, 15}
	fmt.Println(hIndex(list))

	list = []int{1}
	fmt.Println(hIndex(list))

	list = []int{3, 0, 6, 1, 5}
	fmt.Println(hIndex(list))

	list = []int{1, 3, 1}
	fmt.Println(hIndex(list))
}

func hIndex(list []int) int {
	ln := len(list)
	tmp := make([]int, ln+1)
	for _, v := range list {
		if v > ln {
			tmp[ln]++
		} else {
			tmp[v]++
		}
	}

	sum := 0
	for i := ln; i >= 1; i-- {
		sum += tmp[i]
		if sum >= i {
			return i
		}
	}

	return 0
}
