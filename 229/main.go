package main

import "fmt"

func main() {
	fmt.Println(majorityElement([]int{1, 2}))
	fmt.Println(majorityElement([]int{3, 2, 3}))
}

func majorityElement(nums []int) []int {
	a, cna := 0, 0
	b, cnb := 0, 0
	for _, v := range nums {
		if v == a {
			cna++
		} else if v == b {
			cnb++
		} else if cna == 0 {
			a = v
			cna = 1
		} else if cnb == 0 {
			b = v
			cnb = 1
		} else {
			cnb--
			cna--
		}
	}

	cna, cnb = 0, 0
	for _, v := range nums {
		if v == a {
			cna++
		} else if v == b {
			cnb++
		}
	}

	r := make([]int, 0)
	if cna > len(nums)/3 {
		r = append(r, a)
	}
	if cnb > len(nums)/3 {
		r = append(r, b)
	}

	return r
}
