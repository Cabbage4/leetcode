package main

import (
	"math"
	"sort"
)

func main() {
}

func findRadius(houses []int, heaters []int) int {
	sort.Ints(heaters)

	var r int
	for _, h := range houses {
		index := sort.SearchInts(heaters, h)
		diff := math.MaxInt
		if index < len(heaters) {
			diff = heaters[index] - h
		}
		if index-1 >= 0 && h-heaters[index-1] < diff {
			diff = h - heaters[index-1]
		}
		if diff > r {
			r = diff
		}
	}

	return r
}
