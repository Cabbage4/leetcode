package main

import "fmt"

func main() {
	fmt.Println(numWaterBottles(9, 3))
}

func numWaterBottles(numBottles int, numExchange int) int {
	r := numBottles
	c := numBottles

	for c >= numExchange {
		k := c / numExchange
		c = c%numExchange + k
		r += k
	}

	return r
}
