package main

import "fmt"

func main() {
	fmt.Println((1) | 1<<2)
}

func bitwiseComplement(n int) int {
	if n == 0 {
		return 1
	}

	r := 0
	k := 0
	for n > 0 {
		t := n & 1
		r = r | ((t ^ 1) << k)
		k++
		n = n >> 1
	}

	return r
}
