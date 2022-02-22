package main

import "fmt"

func main() {
	fmt.Println(isUgly(6))
	fmt.Println(isUgly(14))
}

func isUgly(n int) bool {
	if n == 0 {
		return false
	}

	if n == 1 {
		return true
	}

	for n%2 == 0 && n != 0 {
		n = n / 2
	}
	for n%3 == 0 && n != 0 {
		n = n / 3
	}
	for n%5 == 0 && n != 0 {
		n = n / 5
	}

	if n != 0 && n != 1 {
		return false
	}
	return true
}
