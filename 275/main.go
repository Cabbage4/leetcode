package main

import "fmt"

func main() {
	fmt.Println(hIndex([]int{0, 1, 4, 4, 5, 6}))
}

func hIndex(list []int) int {
	l := 0
	r := len(list)
	for l < r {
		mid := (l + r) / 2
		if list[mid] < len(list)-mid {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return len(list) - l
}
