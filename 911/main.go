package main

import (
	"fmt"
	"sort"
)

func main() {
	c := Constructor([]int{0, 1, 1, 0, 0, 1, 0}, []int{0, 5, 10, 15, 20, 25, 30})
	fmt.Println(c.Q(3))
	fmt.Println(c.Q(12))
	fmt.Println(c.Q(25))
	fmt.Println(c.Q(15))
	fmt.Println(c.Q(24))
	fmt.Println(c.Q(8))
}

type TopVotedCandidate struct {
	dp    []int
	times []int
}

func Constructor(persons []int, times []int) TopVotedCandidate {
	maxIndex := -1

	mp := make(map[int]int)
	dp := make([]int, len(persons))
	for i := range persons {
		mp[persons[i]]++
		if maxIndex == -1 || mp[persons[i]] >= mp[maxIndex] {
			maxIndex = persons[i]
		}
		dp[i] = maxIndex
	}
	return TopVotedCandidate{
		dp:    dp,
		times: times,
	}
}

func (to *TopVotedCandidate) Q(t int) int {
	index := sort.SearchInts(to.times, t)
	if index >= len(to.times) {
		return to.dp[len(to.dp)-1]
	}

	if to.times[index] == t {
		return to.dp[index]
	}

	if index == 0 {
		// unknown
		return -1
	}

	return to.dp[index-1]
}
