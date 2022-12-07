package main

import (
	"fmt"
)

func main() {
	fmt.Println(getPermutation(3, 1))
	fmt.Println(getPermutation(3, 2))
	fmt.Println(getPermutation(3, 3))
}

var fca = []int{1, 1, 2, 6, 24, 120, 720, 5040, 40320, 362880}

func getPermutation(n int, k int) string {
	k--

	list := make([]int, 0)
	for i := 1; i <= n; i++ {
		list = append(list, i)
	}

	var r string
	for i := n; i >= 1; i-- {
		tr := k / fca[i-1]
		r += string('0' + list[tr])

		k = k % fca[i-1]
		list = append(list[:tr], list[tr+1:]...)
	}
	return r
}
