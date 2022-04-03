package main

import "fmt"

func main() {
	fmt.Println(decrypt([]int{5, 7, 1, 4}, 3))
	fmt.Println(decrypt([]int{2, 4, 9, 3}, -2))
}

func decrypt(code []int, k int) []int {
	ln := len(code)
	r := make([]int, ln)
	if k == 0 {
		return r
	}

	start := 1
	end := k + 1
	if k < 0 {
		k = -k
		start = ln - k
		end = ln
	}
	for i := start; i < end; i++ {
		r[0] += code[i]
	}

	for i := 1; i < ln; i++ {
		r[i] = r[i-1] - code[start%ln] + code[end%ln]
		start++
		end++
	}

	return r
}
