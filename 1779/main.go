package main

import "math"

func main() {
}

func nearestValidPoint(x int, y int, points [][]int) int {
	r := math.MaxInt32
	m := math.MaxInt32
	for i, p := range points {
		if p[0] == x && p[1] == y {
			return i
		}

		if p[0] == x || p[1] == y {
			tm := abs(y, p[1]) + abs(x, p[0])

			if tm < m {
				m = tm
				r = i
			}
		}
	}

	if r == math.MaxInt32 {
		return -1
	}

	return r
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
