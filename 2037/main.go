package main

import "sort"

func main() {
}

func minMovesToSeat(seats []int, students []int) int {
	sort.Ints(seats)
	sort.Ints(students)

	r := 0
	for i := 0; i < len(seats); i++ {
		r += abs(seats[i], students[i])
	}
	return r
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
