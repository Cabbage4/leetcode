package main

import "fmt"

func main() {
	fmt.Println(findPoisonedDuration([]int{1, 4}, 2))
	fmt.Println(findPoisonedDuration([]int{1, 2}, 2))
}

func findPoisonedDuration(timeSeries []int, duration int) int {
	if len(timeSeries) == 0 {
		return 0
	}

	result := 0
	for i := 0; i < len(timeSeries)-1; i++ {
		if timeSeries[i+1]-timeSeries[i] > duration {
			result += timeSeries[i+1] - timeSeries[i]
		} else {
			result += duration
		}
	}

	return result + duration
}
