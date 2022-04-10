package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(sortEvenOdd([]int{5, 39, 33, 5, 12, 27, 20, 45, 14, 25, 32, 33, 30, 30, 9, 14, 44, 15, 21}))
	fmt.Println(sortEvenOdd([]int{4, 1, 2, 3}))
}

type List struct {
	isOdd bool
	data  []int
}

func (l List) Len() int {
	if l.isOdd || len(l.data)%2 == 0 {
		return len(l.data) / 2
	}
	return len(l.data)/2 + 1
}

func (l List) Less(i, j int) bool {
	if l.isOdd {
		return l.data[2*i+1] > l.data[2*j+1]
	}
	return l.data[2*i] < l.data[2*j]
}

func (l List) Swap(i, j int) {
	if l.isOdd {
		l.data[2*i+1], l.data[2*j+1] = l.data[2*j+1], l.data[2*i+1]
		return
	}
	l.data[2*i], l.data[2*j] = l.data[2*j], l.data[2*i]
}

func sortEvenOdd(nums []int) []int {
	l := List{
		isOdd: false,
		data:  nums,
	}

	sort.Sort(l)
	l.isOdd = true
	sort.Sort(l)

	return l.data
}
