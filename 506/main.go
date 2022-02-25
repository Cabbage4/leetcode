package main

import (
	"sort"
	"strconv"
)

func main() {

}

func findRelativeRanks(score []int) []string {
	result := make([]string, len(score))

	imp := make(map[int]int)
	for i, v := range score {
		imp[v] = i
	}

	nList := make([]int, len(score))
	copy(nList, score)
	sort.Ints(nList)

	count := 1
	for i := len(nList) - 1; i >= 0; i-- {
		v := nList[i]
		s := ""
		if count == 1 {
			s = "Gold Medal"
		} else if count == 2 {
			s = "Silver Medal"
		} else if count == 3 {
			s = "Bronze Medal"
		} else {
			s = strconv.Itoa(count)
		}

		count++
		result[imp[v]] = s
	}

	return result
}
