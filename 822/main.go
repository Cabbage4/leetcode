package main

import "math"

func main() {
}

func flipgame(fronts []int, backs []int) int {
	r := math.MaxInt
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	set := make(map[int]bool)
	for i := 0; i < len(fronts); i++ {
		if fronts[i] == backs[i] {
			set[fronts[i]] = true
		}
	}

	for i := 0; i < len(fronts); i++ {
		if fronts[i] == backs[i] {
			continue
		}

		if !set[fronts[i]] {
			r = min(r, fronts[i])
		}
		if !set[backs[i]] {
			r = min(r, backs[i])
		}
	}

	if r == math.MaxInt {
		return 0
	}

	return r
}
