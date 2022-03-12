package main

import "fmt"

func main() {
	fmt.Println(getNoZeroIntegers(11))
}

func getNoZeroIntegers(n int) []int {
	a, b := 0, 0
	s := 1
	for n > 0 {
		d := n % 10
		n = n / 10

		if (d == 0 || d == 1) && n > 0 {
			a += s * (1 + d)
			b += s * (9)
			n--
		} else {
			a += s * 1
			b += s * (d - 1)
		}

		s *= 10
	}
	return []int{a, b}
}
