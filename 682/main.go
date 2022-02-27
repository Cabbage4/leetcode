package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(calPoints([]string{"5", "2", "C", "D", "+"}))
}

func calPoints(ops []string) int {
	list := make([]int, 0)
	for _, op := range ops {
		if op == "C" {
			list = list[:len(list)-1]
		} else if op == "D" {
			list = append(list, list[len(list)-1]*2)
		} else if op == "+" {
			list = append(list, list[len(list)-1]+list[len(list)-2])
		} else {
			v, _ := strconv.Atoi(op)
			list = append(list, v)
		}
	}

	result := 0
	for _, v := range list {
		result += v
	}

	return result
}
