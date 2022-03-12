package main

import "fmt"

func main() {
	fmt.Println(containsPattern([]int{2, 2}, 1, 2))
	fmt.Println(containsPattern([]int{1, 2, 4, 4, 4, 4}, 1, 3))
	fmt.Println(containsPattern([]int{1, 2, 1, 2, 1, 1, 1, 3}, 2, 2))
	fmt.Println(containsPattern([]int{1, 2, 1, 2, 1, 3}, 2, 3))
}

func containsPattern(arr []int, m int, k int) bool {
	ln := len(arr)
	c := 0
	for i := 0; i < ln-m; i++ {
		if arr[i] != arr[i+m] {
			c = 0
			continue
		}
		c++
		if c == (k-1)*m {
			return true
		}
	}

	return false
}
