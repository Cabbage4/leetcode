package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(hammingDistance(1, 4))
	fmt.Println(hammingDistance(1, 3))
}

func hammingDistance(x int, y int) int {
	z := strconv.FormatInt(int64(x^y), 2)

	result := 0
	for _, v := range z {
		if v == '1' {
			result++
		}
	}

	return result
}
