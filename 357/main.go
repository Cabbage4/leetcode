package main

import "fmt"

func main() {
	fmt.Println(countNumbersWithUniqueDigits(3))
}

func countNumbersWithUniqueDigits(n int) int {
	if n == 0 {
		return 1
	}

	r := 10
	tr := 9
	k := 9
	for i := 2; i <= n; i++ {
		tr *= k
		r += tr
		k--
	}
	return r
}
