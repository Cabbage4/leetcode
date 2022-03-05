package main

import "fmt"

func main() {
	fmt.Println(validMountainArray([]int{1, 3, 2}))
}

func validMountainArray(arr []int) bool {
	hasUp := false
	hasLow := false

	isUpper := true
	for i := 1; i < len(arr); i++ {
		if isUpper {
			if arr[i] > arr[i-1] {
				hasUp = true
				continue
			} else if arr[i] == arr[i-1] {
				return false
			}

			isUpper = false
			hasLow = true
			continue
		}
		if arr[i] >= arr[i-1] {
			return false
		}
		hasLow = true
	}

	return hasUp && hasLow
}
