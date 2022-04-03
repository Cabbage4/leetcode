package main

import "fmt"

func main() {
	fmt.Println(sumOfUnique([]int{1, 1, 1, 1, 1}))
}

func sumOfUnique(nums []int) int {
	r := 0
	mp := make([]int, 101)
	for _, v := range nums {
		mp[v]++

		if mp[v] == 1 {
			r += v
		} else if mp[v] == 2 {
			r -= v
		}
	}
	return r
}
