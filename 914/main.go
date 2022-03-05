package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(hasGroupsSizeX([]int{1, 2, 3, 4, 4, 3, 2, 1}))
	fmt.Println(hasGroupsSizeX([]int{1, 1, 1, 2, 2, 2, 3, 3}))
	fmt.Println(hasGroupsSizeX([]int{1, 1, 2, 2, 2, 2}))
	fmt.Println(hasGroupsSizeX([]int{1, 1, 1, 1, 2, 2, 2, 2, 2, 2}))
}

func hasGroupsSizeX(deck []int) bool {
	if len(deck) == 1 {
		return false
	}

	ln := int(math.Pow10(4)) + 1
	mp := make([]int, ln)
	for _, v := range deck {
		mp[v]++
	}

	min := math.MaxInt32
	for _, v := range mp {
		if v == 0 {
			continue
		}
		if v == 1 {
			return false
		}
		if v < min {
			min = v
		}
	}

	for i := 2; i <= min; i++ {
		if min%i != 0 {
			continue
		}

		flag := true
		for _, v := range mp {
			if v == 0 {
				continue
			}
			if v%i != 0 {
				flag = false
				break
			}
		}

		if flag {
			return true
		}
	}

	return false
}
