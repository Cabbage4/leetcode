package main

import "fmt"

func main() {
	fmt.Println(isRectangleOverlap([]int{4, 4, 14, 7}, []int{4, 3, 8, 8}))
}

func isRectangleOverlap(a []int, b []int) bool {
	if a[0] == a[2] ||
		a[1] == a[3] ||
		b[0] == b[2] ||
		b[1] == b[3] {
		return false
	}

	flag := b[2] <= a[0] ||
		b[3] <= a[1] ||
		b[0] >= a[2] ||
		b[1] >= a[3]

	return !flag
}
