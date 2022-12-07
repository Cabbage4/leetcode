package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(pivotInteger(15))
	fmt.Println(pivotInteger(4))
	fmt.Println(pivotInteger(8))
}

func pivotInteger(n int) int {
	target := (n*n + n) / 2
	r := math.Sqrt(float64(target))

	if r != math.Floor(r) {
		return -1
	}

	return int(r)
}
