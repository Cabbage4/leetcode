package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(shortestToChar("aaba", 'b'))
	fmt.Println(shortestToChar("loveleetcode", 'e'))
}

func shortestToChar(s string, c byte) []int {
	result := make([]int, len(s))
	cLI := -1
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			cLI = i
			result[i] = 0
			continue
		}

		if cLI == -1 {
			result[i] = math.MaxInt32
			continue
		}
		result[i] = i - cLI
	}

	cRI := -1
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == c {
			cRI = i
			continue
		}

		if cRI == -1 {
			continue
		}

		result[i] = min(cRI-i, result[i])
	}

	return result
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
