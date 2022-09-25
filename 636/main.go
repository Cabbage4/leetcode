package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(exclusiveTime(1, []string{"0:start:0", "0:start:2", "0:end:5", "0:start:6", "0:end:6", "0:end:7"}))
	fmt.Println(exclusiveTime(2, []string{"0:start:0", "1:start:2", "1:end:5", "0:end:6"}))
}

func exclusiveTime(n int, logs []string) []int {
	r := make([]int, n)
	s := make([][]int, 0)
	for _, l := range logs {
		info := strings.Split(l, ":")
		id, _ := strconv.Atoi(info[0])
		timestamp, _ := strconv.Atoi(info[2])

		if info[1] == "start" {
			if len(s) > 0 {
				r[s[len(s)-1][0]] += timestamp - s[len(s)-1][1]
				s[len(s)-1][1] = timestamp
			}

			s = append(s, []int{id, timestamp})
			continue
		}

		last := s[len(s)-1]
		s = s[:len(s)-1]

		r[id] += timestamp - last[1] + 1

		if len(s) > 0 {
			s[len(s)-1][1] = timestamp + 1
		}
	}

	return r
}
