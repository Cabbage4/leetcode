package main

import (
	"fmt"
)

func main() {
	fmt.Println(minSteps(100))
}

func minSteps(n int) int {
	var r int
	for i := 2; i*i <= n; i++ {
		for n%i == 0 {
			n /= i
			r += i
		}
	}

	if n > 1 {
		r += n
	}

	return r
}
