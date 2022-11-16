package main

import (
	"fmt"
)

func main() {
	fmt.Println(totalFruit([]int{3, 3, 3, 1, 4}))
	fmt.Println(totalFruit([]int{0, 1, 0, 2}))
	fmt.Println(totalFruit([]int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4}))
	fmt.Println(totalFruit([]int{1, 2, 1}))
	fmt.Println(totalFruit([]int{0, 1, 2, 2}))
	fmt.Println(totalFruit([]int{1, 2, 3, 2, 2}))
}

func totalFruit(fruits []int) int {
	var r int
	left := 0
	mp := make(map[int]int)
	for right := range fruits {
		mp[fruits[right]]++

		for len(mp) > 2 {
			mp[fruits[left]]--
			if mp[fruits[left]] == 0 {
				delete(mp, fruits[left])
			}
			left++
		}

		r = max(r, right-left+1)
	}

	return r
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
