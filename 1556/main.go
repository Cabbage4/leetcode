package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(thousandSeparator(1234))
}

func thousandSeparator(n int) string {
	if n == 0 {
		return "0"
	}

	r := ""
	c := 0
	for n > 0 {
		t := n % 10
		r = strconv.Itoa(t) + r

		n = n / 10
		c++

		if c >= 3 && n != 0 {
			c = 0
			r = "." + r
		}
	}
	return r
}
