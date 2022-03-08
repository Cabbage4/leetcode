package main

import "fmt"

func main() {
	fmt.Println(addToArrayForm([]int{1, 2, 0, 0}, 34))
	fmt.Println(addToArrayForm([]int{2, 7, 4}, 181))
}

func addToArrayForm(r []int, k int) []int {
	offset := 0
	ir := len(r) - 1
	for k > 0 && ir >= 0 {
		v := r[ir] + k%10 + offset
		offset = v / 10
		r[ir] = v % 10
		ir--
		k = k / 10
	}

	for offset != 0 && ir >= 0 {
		v := r[ir] + offset
		offset = v / 10
		r[ir] = v % 10
		ir--
	}
	for k > 0 {
		v := k%10 + offset
		offset = v / 10
		r = append([]int{v % 10}, r...)
		k = k / 10
	}
	if offset != 0 {
		r = append([]int{offset}, r...)
	}

	return r
}
