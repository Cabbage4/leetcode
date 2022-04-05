package main

import "fmt"

func main() {
	fmt.Println(sumBase(34, 6))
}

func sumBase(n int, k int) int {
	r := 0
	for n > 0 {
		r += n % k
		n = n / k
	}
	return r
}
