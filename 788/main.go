package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(rotatedDigits(10))
}

func rotatedDigits(n int) int {
	list := []int{0, 0, 1, -1, -1, 1, 1, -1, 0, 1}

	var r int
	for i := 1; i <= n; i++ {
		is := strconv.Itoa(i)
		valid, diff := true, false

		for j := 0; j < len(is); j++ {
			if list[is[j]-'0'] == 1 {
				diff = true
			} else if list[is[j]-'0'] == -1 {
				valid = false
				break
			}
		}

		if valid && diff {
			r++
		}
	}
	return r
}
