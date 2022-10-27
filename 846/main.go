package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(isNStraightHand([]int{1, 1, 2, 2, 3, 3}, 2))
	fmt.Println(isNStraightHand([]int{1, 2, 3, 6, 2, 3, 4, 7, 8}, 3))
}

func isNStraightHand(hand []int, groupSize int) bool {
	if len(hand)%groupSize != 0 {
		return false
	}

	sort.Ints(hand)
	mp := make(map[int]int)
	for _, h := range hand {
		mp[h]++
	}

	for _, h := range hand {
		if mp[h] == 0 {
			continue
		}

		for i := 0; i < groupSize; i++ {
			if mp[h+i] == 0 {
				return false
			}

			mp[h+i]--
		}
	}

	return true
}
