package main

import "fmt"

func main() {
	fmt.Println(countOdds(3, 7))
	fmt.Println(countOdds(8, 10))
}

func countOdds(low int, high int) int {
	if low%2 == 0 {
		low++
	}
	if high%2 == 0 {
		high--
	}
	return (high - low + 2) / 2
}
