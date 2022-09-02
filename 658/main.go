package main

import (
	"fmt"
	"sort"
)

func main() {
	//fmt.Println(search([]int{1, 2}, 4))
	fmt.Println(findClosestElements([]int{1, 2, 3, 4, 5}, 4, 3))
	fmt.Println(findClosestElements([]int{-2, -1, 1, 2, 3, 4, 5}, 7, 3))
	fmt.Println(findClosestElements([]int{0, 0, 1, 2, 3, 3, 4, 7, 7, 8}, 3, 5))
}

func findClosestElements(arr []int, k int, x int) []int {
	right := sort.SearchInts(arr, x)
	left := right - 1
	for ; k > 0; k-- {
		if left < 0 {
			right++
		} else if right >= len(arr) || (x-arr[left] <= arr[right]-x) {
			left--
		} else {
			right++
		}
	}
	return arr[left+1 : right]
}
