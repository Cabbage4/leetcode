package main

import "fmt"

func main() {
	fmt.Println(isCovered([][]int{{1, 2}, {3, 4}, {5, 6}}, 2, 5))
}

func isCovered(ranges [][]int, left int, right int) bool {
	bits := make([]int, 52)
	for _, r := range ranges {
		bits[r[0]]++
		bits[r[1]+1]--
	}

	c := 0
	for i := 1; i <= right; i++ {
		c += bits[i]
		if i >= left && c == 0 {
			return false
		}
	}

	return true
}
