package main

import (
	"fmt"
	"math"
)

func main() {
	for i := 0; i < 100; i++ {
		fmt.Println(i, bulbSwitch(i))
	}
}

func bulbSwitch(n int) int {
	return int(math.Sqrt(float64(n)))
}
