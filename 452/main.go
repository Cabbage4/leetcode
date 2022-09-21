package main

import "sort"

func main() {
}

func findMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})

	r := 1
	maxRight := points[0][1]
	for _, p := range points {
		if p[0] > maxRight {
			maxRight = p[1]
			r++
		}
	}
	return r
}
