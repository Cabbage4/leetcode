package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println(sumOfNumberAndReverse(0))
	fmt.Println(sumOfNumberAndReverse(11))
	fmt.Println(sumOfNumberAndReverse(10))
	fmt.Println(sumOfNumberAndReverse(181))
	fmt.Println(sumOfNumberAndReverse(63))
	fmt.Println(sumOfNumberAndReverse(443))
}

func sumOfNumberAndReverse(num int) bool {
	if num == 0 {
		return true
	}

	ln := len(strconv.Itoa(num))
	min := int(math.Pow10(ln - 2))

	reverse := func(v int) int {
		var r int
		for v != 0 {
			r = r*10 + v%10
			v = v / 10
		}
		return r
	}

	for i := min; i <= num; i++ {
		if i+reverse(i) == num {
			return true
		}
	}
	return false
}
