package main

import "fmt"

func main() {
	fmt.Println(reachNumber(3))
}

func reachNumber(target int) int {
	if target < 0 {
		target = -target
	}

	c := 0
	for target > 0 {
		c++
		target -= c
	}

	if target%2 == 0 {
		return c
	}

	return c + 1 + c%2
}
