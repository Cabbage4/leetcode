package main

import (
	"fmt"
)

func main() {
	fmt.Println(kthSmallestPrimeFraction([]int{1, 2, 3, 5}, 3))
}

func kthSmallestPrimeFraction(arr []int, k int) []int {
	left := 0.
	right := 1.
	for {
		mid := (left + right) / 2
		x, y := 0, 1
		i, count := -1, 0
		for j := 1; j < len(arr); j++ {
			for float64(arr[i+1])/float64(arr[j]) < mid {
				i++
				if arr[i]*y > arr[j]*x {
					x, y = arr[i], arr[j]
				}
			}
			count += i + 1
		}

		if count == k {
			return []int{x, y}
		}
		if count < k {
			left = mid
		}
		if count > k {
			right = mid
		}
	}
}
