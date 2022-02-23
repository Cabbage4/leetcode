package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(readBinaryWatch(1))
}

func readBinaryWatch(turnedOn int) []string {
	result := make([]string, 0)
	for i := 0; i < 12; i++ {
		for j := 0; j < 60; j++ {
			f := strconv.FormatInt(int64(i<<6|j), 2)
			if strings.Count(f, "1") == turnedOn {
				if j < 10 {
					result = append(result, fmt.Sprintf("%d:0%d", i, j))
				} else {
					result = append(result, fmt.Sprintf("%d:%d", i, j))
				}
			}
		}
	}

	return result
}
