package main

import "fmt"

func main() {
	fmt.Println(computeArea(-2, -2, 2, 2, 1, -3, 3, 3) == 24)
	fmt.Println(computeArea(-2, -2, 2, 2, -1, 4, 1, 6) == 20)
	fmt.Println(computeArea(-2, -2, 2, 2, 3, 3, 4, 4) == 17)
	fmt.Println(computeArea(-3, 0, 3, 4, 0, -1, 9, 2) == 45)
	fmt.Println(computeArea(-2, -2, 2, 2, -2, -2, 2, 2) == 16)
}

func computeArea(ax1 int, ay1 int, ax2 int, ay2 int, bx1 int, by1 int, bx2 int, by2 int) int {
	m1 := (ax1 - ax2) * (ay1 - ay2)
	m2 := (bx1 - bx2) * (by1 - by2)

	flag := (ax1 <= bx1 && bx1 <= ax2 && ay1 < by2 && by1 < ay2) ||
		(bx1 <= ax1 && ax1 <= bx2 && by1 < ay2 && ay1 < by2)

	if flag {
		cx1, cx2 := max(ax1, bx1), min(ax2, bx2)
		cy1, cy2 := max(ay1, by1), min(ay2, by2)
		m3 := (cx1 - cx2) * (cy1 - cy2)
		return m1 + m2 - m3
	}

	return m1 + m2
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
