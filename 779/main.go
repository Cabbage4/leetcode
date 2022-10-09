package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(kthGrammar(2, 1))
	fmt.Println(kthGrammar(2, 2))
}

func kthGrammar(n int, k int) int {
	if n == 1 {
		return 0
	}

	m := int(math.Pow(2, float64(n-2)))
	if k <= m {
		return kthGrammar(n-1, k)
	}

	return (kthGrammar(n-1, k-m) + 1) % 2
}
