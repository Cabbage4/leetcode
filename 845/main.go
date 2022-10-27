package main

import "fmt"

func main() {
	fmt.Println(longestMountain([]int{3, 4}))
	fmt.Println(longestMountain([]int{3, 2}))
	fmt.Println(longestMountain([]int{2, 2, 2}))
	fmt.Println(longestMountain([]int{2, 1, 4, 7, 3, 2, 5}))
}

func longestMountain(arr []int) int {
	left := 0
	lmp := make([]int, len(arr))
	lmp[0] = 0
	for i := 1; i < len(arr); i++ {
		if arr[i] > arr[i-1] {
			left++
			lmp[i] = left
		} else {
			left = 0
		}
	}

	var r int
	right := 0
	for i := len(arr) - 2; i >= 0; i-- {
		if arr[i] > arr[i+1] {
			right++
			if lmp[i] != 0 && r < lmp[i]+right+1 {
				r = lmp[i] + right + 1
			}
		} else {
			right = 0
		}
	}

	return r
}
