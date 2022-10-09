package main

import "fmt"

func main() {
	fmt.Println(numRabbits([]int{1, 0, 1, 0, 0}))
	fmt.Println(numRabbits([]int{0, 0, 1, 1, 1}))
}

func numRabbits(answers []int) int {
	var mp [1001]int
	for _, a := range answers {
		mp[a]++
	}

	var r int
	for c, cc := range mp {
		if cc == 0 {
			continue
		}

		for cc > 0 {
			r += c + 1
			cc = cc - c - 1
		}
	}
	return r
}
