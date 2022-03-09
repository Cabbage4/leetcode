package main

import (
	"fmt"
	"sort"
)

func main() {
	list := []int{6, 0, 7, 0, 7, 5, 7, 8, 3, 4, 0, 7, 8, 1, 6, 8, 1, 1, 2, 4, 8, 1, 9, 5, 4, 3, 8, 5, 10, 8, 6, 6, 1, 0, 6, 10, 8, 2, 3, 4}
	//list := []int{1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 3}
	fmt.Println(trimMean(list))
}

func trimMean(arr []int) float64 {
	sort.Ints(arr)
	l := len(arr) / 20
	r := len(arr) - l
	s := 0
	for i := l; i < r; i++ {
		s += arr[i]
	}
	return float64(s) / float64(r-l)
}
