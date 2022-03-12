package main

import "math"

func main() {
}

func average(salary []int) float64 {
	mx := math.MinInt32
	mn := math.MaxInt32
	sum := 0

	for _, v := range salary {
		sum += v
		if v > mx {
			mx = v
		}

		if v < mn {
			mn = v
		}
	}

	return float64(sum-mx-mn) / float64(len(salary)-2)
}
