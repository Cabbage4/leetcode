package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(leastInterval([]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 2))
}

func leastInterval(tasks []byte, n int) int {
	mp := make(map[byte]int)
	for _, v := range tasks {
		mp[v]++
	}

	countList := make([]int, 0)
	timeList := make([]int, 0)
	for _, c := range mp {
		timeList = append(timeList, 1)
		countList = append(countList, c)
	}

	var r int
	for range tasks {
		r++
		minTime := math.MaxInt
		for i, t := range timeList {
			if countList[i] > 0 && t < minTime {
				minTime = t
			}
		}

		if minTime > r {
			r = minTime
		}

		taskIndex := -1
		for i, c := range countList {
			if c > 0 && minTime >= timeList[i] && (taskIndex == -1 || countList[taskIndex] < c) {
				taskIndex = i
			}
		}

		timeList[taskIndex] = minTime + n + 1
		countList[taskIndex]--
	}

	return r
}
