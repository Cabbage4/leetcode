package main

import (
	"math"
)

func main() {

}

var pmp = func() []bool {
	r := make([]bool, 101)
	for i := 1; i <= 100; i++ {
		if is(i) {
			r[i] = true
		}
	}
	return r
}()

func numPrimeArrangements(n int) int {
	ln := int(math.Pow10(9)) + 7

	pr := 0
	for i := 2; i <= n; i++ {
		if pmp[i] {
			pr++
		}
	}

	r := 1
	for i := 1; i <= n-pr; i++ {
		r = (r * i) % ln
	}

	for i := 1; i <= pr; i++ {
		r = (r * i) % ln
	}

	return r
}

func is(n int) bool {
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
