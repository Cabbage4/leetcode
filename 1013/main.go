package main

import "fmt"

func main() {
	fmt.Println(canThreePartsEqualSum([]int{0, 0, 0, 0}))
	fmt.Println(canThreePartsEqualSum([]int{12, -4, 16, -5, 9, -3, 3, 8, 0}))
	fmt.Println(canThreePartsEqualSum([]int{3, 3, 6, 5, -2, 2, 5, 1, -9, 4}))
	fmt.Println(canThreePartsEqualSum([]int{0, 2, 1, -6, 6, -7, 9, 1, 2, 0, 1}))
}

func canThreePartsEqualSum(arr []int) bool {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	if sum%3 != 0 {
		return false
	}

	avg := sum / 3

	tmp := 0
	count := 0
	for _, v := range arr {
		tmp += v
		if tmp == avg {
			count++
			tmp = 0
		}

		if count == 3 {
			return true
		}
	}

	return false
}
