package main

import "fmt"

func main() {
	fmt.Println(maxDistToClosest([]int{1, 0, 1}))
	fmt.Println(maxDistToClosest([]int{0, 1}))
	fmt.Println(maxDistToClosest([]int{1, 0, 0, 0}))
}

func maxDistToClosest(seats []int) int {
	left := 0
	for seats[left] == 0 {
		left++
	}

	r := left

	cur := 0
	for i := left; i < len(seats); i++ {
		if seats[i] == 0 {
			cur++
		} else {
			if cur%2 == 0 && cur/2 > r {
				r = cur / 2
			}
			if cur%2 != 0 && cur/2+1 > r {
				r = cur/2 + 1
			}

			cur = 0
		}
	}

	if cur != 0 && cur > r {
		r = cur
	}

	return r
}
