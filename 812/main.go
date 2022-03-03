package main

import (
	"fmt"
)

func main() {
	test1 := [][]int{{0, 0}, {0, 1}, {1, 0}, {0, 2}, {2, 0}}
	fmt.Println(largestTriangleArea(test1))
}

func largestTriangleArea(points [][]int) float64 {
	var result float64 = 0
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			for k := j + 1; k < len(points); k++ {
				s := gs(points[i], points[j], points[k])
				if s > result {
					result = s
				}
			}
		}
	}
	return result
}

func gs(a, b, c []int) float64 {
	x1, x2, x3 := a[0], b[0], c[0]
	y1, y2, y3 := a[1], b[1], c[1]
	r := 0.5 * float64(x1*(y2-y3)+x2*(y3-y1)+x3*(y1-y2))
	if r < 0 {
		return -r
	}
	return r
}
