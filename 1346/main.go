package main

import (
	"fmt"
)

func main() {
	//fmt.Println(checkIfExist([]int{10, 2, 5, 3}))
	fmt.Println(checkIfExist([]int{-2, 0, 10, -19, 4, 6, -8}))
}

func checkIfExist(arr []int) bool {
	mp := make(map[int]bool)
	for _, v := range arr {
		if v == 0 && mp[v] {
			return true
		}
		mp[v] = true
	}

	for k, v := range mp {
		if k != 0 && v && mp[2*k] {
			return true
		}
	}

	return false
}
