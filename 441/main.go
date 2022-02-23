package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(arrangeCoins(8))
	fmt.Println(arrangeCoins(5))
}

func arrangeCoins(n int) int {
	return int(math.Sqrt(float64(2*n)+0.25) - 0.5)
}
