package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maximum69Number(9669))
}

func maximum69Number(n int) int {
	ln := 0
	tn := n
	for tn > 0 {
		tn = tn / 10
		ln++
	}

	r := 0
	c := 1
	for ln > 0 {
		tp := int(math.Pow10(ln - 1))

		t := n / tp
		if t == 6 && c == 1 {
			r = r*10 + 9
			c = 0
		} else {
			r = r*10 + t
		}

		n = n % tp
		ln--
	}

	return r
}
