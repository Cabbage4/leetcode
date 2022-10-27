package main

import "fmt"

func main() {
	fmt.Println(mirrorReflection(2, 1))
}

func mirrorReflection(p int, q int) int {
	s := (p * q) / gcd(p, q)
	k := s / q

	if ((k*q)/p)%2 == 0 {
		return 0
	}

	if k%2 == 0 {
		return 2
	}
	return 1
}

func gcd(a, b int) int {
	if a%b == 0 {
		return b
	}
	return gcd(b, a%b)
}
