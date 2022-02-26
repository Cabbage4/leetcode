package main

import (
	"math"
)

func main() {
}

func findRestaurant(list1 []string, list2 []string) []string {
	mp := make(map[string]bool)
	mpI := make(map[string]int)

	for i, v := range list1 {
		mpI[v] = i
		mp[v] = true
	}

	minIndex := math.MaxInt32
	result := make([]string, 0)
	for i, v := range list2 {
		if !mp[v] {
			continue
		}

		if i+mpI[v] < minIndex {
			minIndex = i + mpI[v]
			result = []string{v}
		} else if i+mpI[v] == minIndex {
			result = append(result, v)
		}
	}

	return result
}
