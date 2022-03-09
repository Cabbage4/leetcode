package main

import "fmt"

func main() {
	fmt.Println(distributeCandies(7, 4))
	fmt.Println(distributeCandies(10, 3))
}

func distributeCandies(c int, n int) []int {
	count := 1
	tc := c

	left := 1
	right := n
	sum := ((left + right) * right) / 2
	for sum <= tc {

		tc = tc - sum
		left = right + 1
		right = right + n
		count++

		sum = ((left + right) * n) / 2
	}

	r := make([]int, n)

	if count > 1 {
		cc := count - 1
		for i := 0; i < len(r); i++ {
			r[i] = ((2*i + 2 + n*(cc-1)) * cc) / 2
		}
	}

	for i := 0; i < len(r); i++ {
		t := n*(count-1) + i + 1
		if t < tc {
			r[i] += t
		} else {
			r[i] += tc
			break
		}
		tc = tc - t
	}

	return r
}
