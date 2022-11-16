package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(numRescueBoats([]int{3, 2, 2, 1}, 3))
	fmt.Println(numRescueBoats([]int{3, 5, 3, 4}, 5))
}

func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	left := 0
	right := len(people) - 1
	var r int
	for left <= right {
		if people[left] >= limit {
			r += right - left + 1
			break
		}

		if (people[right] >= limit) || (people[left]+people[right] > limit) {
			r++
			right--
			continue
		}

		left++
		right--
		r++
	}
	return r
}
