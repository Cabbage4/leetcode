package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(findMinDifference([]string{"23:59", "00:00"}))
}

func findMinDifference(timePoints []string) int {
	if len(timePoints) > 1440 {
		return 0
	}

	timeIntList := make([]int, len(timePoints))
	for i, v := range timePoints {
		timeIntList[i] = parse(v)
	}
	sort.Ints(timeIntList)
	ln := len(timeIntList)

	r := (timeIntList[0] - timeIntList[ln-1] + 1440) % 1440
	for i := 1; i < ln; i++ {
		tmp := timeIntList[i] - timeIntList[i-1]
		r = min(tmp, r)
	}

	return r
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func parse(t string) int {
	list := strings.Split(t, ":")
	a, _ := strconv.Atoi(list[0])
	b, _ := strconv.Atoi(list[1])
	return a*60 + b
}
