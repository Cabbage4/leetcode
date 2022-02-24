package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(constructRectangle(4))
	fmt.Println(constructRectangle(37))
	fmt.Println(constructRectangle(122122))
}

func constructRectangle(area int) []int {
	x := int(math.Sqrt(float64(area)))
	for x >= 0 {
		if m := area % x; m == 0 {
			return []int{area / x, x}
		}
		x--
	}
	return nil
}
