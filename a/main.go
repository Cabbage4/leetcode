package main

import "fmt"

func main() {
	fmt.Println(lcm(2, 5))
}

func smallestEvenMultiple(n int) int {
	if n%2 == 0 {
		return n
	}
	return n * 2
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(a, b int) int {
	return a * (b / gcd(a, b))
}
