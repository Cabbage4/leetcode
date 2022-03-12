package main

import "sort"

func main() {
}

func sortByBits(a []int) []int {
	sort.Slice(a, func(i, j int) bool {
		ic, jc := c(a[i]), c(a[j])
		if ic != jc {
			return ic < jc
		}

		return a[i] < a[j]
	})
	return a
}

func c(n int) int {
	r := 0
	for n > 0 {
		if n&1 == 1 {
			r++
		}
		n = n >> 1
	}
	return r
}
