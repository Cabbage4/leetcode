package main

import (
	"fmt"
)

func main() {
	fmt.Println(grayCode(3))
}

func grayCode(n int) []int {
	if n == 1 {
		return []int{0, 1}
	}

	pr := grayCode(n - 1)
	tr := make([]int, 0)
	for i := 0; i < len(pr); i++ {
		tr = append(tr, pr[len(pr)-1-i]|1<<(n-1))
	}

	return append(pr, tr...)
}
