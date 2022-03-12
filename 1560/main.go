package main

import (
	"fmt"
)

func main() {
	fmt.Println(mostVisited(4, []int{1, 3, 1, 2}))
	fmt.Println(mostVisited(2, []int{2, 1, 2, 1, 2, 1, 2, 1, 2}))
}

func mostVisited(n int, rounds []int) []int {
	r := make([]int, 0)
	for i := rounds[0]; i <= rounds[len(rounds)-1]; i++ {
		r = append(r, i)
	}
	if len(r) > 0 {
		return r
	}

	for i := rounds[0]; i <= n; i++ {
		r = append(r, i)
	}
	for i := 1; i <= rounds[len(rounds)-1]; i++ {
		r = append(r, i)
	}

	return r
}
