package main

import "fmt"

func main() {
	fmt.Println(commonFactors(43, 945))
}

func commonFactors(a int, b int) int {
	max := gcd(a, b)
	set := make(map[int]bool)
	for i := 1; i <= max; i++ {
		if a%i == 0 && b%i == 0 {
			set[i] = true
		}
	}
	fmt.Println(set)
	return len(set)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
