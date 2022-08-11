package main

import "fmt"

func main() {
	n := 2
	for i := 1; i <= n; i++ {
		fmt.Printf("%d\t", i)
	}
	fmt.Println()
	fmt.Println(lastRemaining(n))
}

func lastRemaining(n int) int {
	if n == 1 {
		return 1
	}

	if n%2 == 0 {
		n++
	}
	return 2 * h(n/2, false)
}

func h(n int, isFormLeft bool) int {
	if n == 1 {
		return 1
	}

	if n%2 != 0 {
		return 2 * h(n/2, !isFormLeft)
	}

	if isFormLeft {
		return 2 * h(n/2, !isFormLeft)
	}

	return 2*h(n/2, !isFormLeft) - 1
}
