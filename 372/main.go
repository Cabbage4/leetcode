package main

import "fmt"

func main() {
	fmt.Println(superPow(2, []int{3}))
	fmt.Println(superPow(2, []int{1, 0}))
}

func superPow(a int, b []int) int {
	r := 1
	for i := len(b) - 1; i >= 0; i-- {
		r = (r * pow(a, b[i])) % 1337
		a = pow(a, 10)
	}
	return r
}

func pow(a int, b int) int {
	r := 1
	h := a % 1337
	for b > 0 {
		if b%2 == 1 {
			r = (r * h) % 1337
		}

		h = (h * h) % 1337
		b = b >> 1
	}
	return r
}
