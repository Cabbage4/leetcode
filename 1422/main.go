package main

import "math"

func main() {
}

func maxScore(s string) int {
	r := math.MinInt32
	one := 0
	zero := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			one++
		} else {
			zero++
		}
		if i != len(s)-1 && zero-one > r {
			r = zero - one
		}
	}

	return r + one
}
