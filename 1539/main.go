package main

import "fmt"

func main() {
	fmt.Println(findKthPositive([]int{1, 3}, 1))
}

func findKthPositive(arr []int, k int) int {
	left := 0
	right := len(arr)
	for left < right {
		mid := (left + right) / 2
		if arr[mid]-mid-1 < k {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left + k
}
