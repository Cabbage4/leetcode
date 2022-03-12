package main

import "sort"

func main() {
}

func canMakeArithmeticProgression(a []int) bool {
	sort.Ints(a)
	d := a[1] - a[0]

	for i := 2; i < len(a); i++ {
		if a[i]-a[i-1] != d {
			return false
		}
	}

	return true
}
