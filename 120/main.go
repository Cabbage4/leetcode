package main

import "math"

func main() {
}

func minimumTotal(t [][]int) int {
	r := math.MaxInt32
	for i := 1; i < len(t); i++ {
		t[i][0] += t[i-1][0]
		t[i][len(t[i])-1] += t[i-1][len(t[i-1])-1]

		for j := 1; j < len(t[i])-1; j++ {
			t[i][j] += min(t[i-1][j], t[i-1][j-1])
		}
	}

	for _, v := range t[len(t)-1] {
		r = min(r, v)
	}

	return r
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
