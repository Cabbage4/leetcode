package main

import (
	"fmt"
)

func main() {
	fmt.Println(specialArray([]int{2, 2}))
}

func specialArray(nums []int) int {
	mp := make([]int, 102)
	for _, i := range nums {
		if i > 100 {
			mp[101]++
		} else {
			mp[i]++
		}
	}

	for i := 100; i > 0; i-- {
		mp[i] += mp[i+1]
		if i == mp[i] {
			return i
		}
	}

	return -1
}
