package main

import (
	"fmt"
)

func main() {
	fmt.Println(isPossible([]int{1, 2, 3, 3, 4, 5}))
	fmt.Println(isPossible([]int{4, 5, 6, 7, 7, 8, 8, 9, 10, 11}))
}

func isPossible(nums []int) bool {
	mp := make(map[int]int)
	for _, v := range nums {
		mp[v]++
	}

	endMp := make(map[int]int)
	for _, v := range nums {
		if mp[v] == 0 {
			continue
		}

		if endMp[v-1] > 0 {
			mp[v]--
			endMp[v-1]--
			endMp[v]++
		} else if mp[v+1] > 0 && mp[v+2] > 0 {
			mp[v]--
			mp[v+1]--
			mp[v+2]--
			endMp[v+2]++
		} else {
			return false
		}
	}

	return true
}
