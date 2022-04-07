package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(isThree(1))
}

func isThree(n int) bool {
	mp := make(map[int]bool)
	list := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101}
	for _, v := range list {
		mp[v] = true
	}

	s := int(math.Sqrt(float64(n)))

	return s*s == n && mp[s]
}
