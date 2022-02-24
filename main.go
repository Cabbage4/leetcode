package main

import "fmt"

func main() {
	fmt.Println(findPoisonedDuration([]int{1, 4}, 2))
}

func findPoisonedDuration(timeSeries []int, duration int) int {
	end := timeSeries[len(timeSeries)-1] + duration
	start := timeSeries[0]

	ti := 0
	result := 0
	for i := start; i <= end; i++ {
		for ti < len(timeSeries) && timeSeries[ti]+duration-1 < i {
			ti++
		}
		if ti == len(timeSeries)-1 {
			break
		}

		if i >= timeSeries[ti] {
			result++
		}
	}
	return result
}
