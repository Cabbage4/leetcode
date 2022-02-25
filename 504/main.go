package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(convertToBase7(100))
	fmt.Println(convertToBase7(-7))
}

func convertToBase7(num int) string {
	result := ""

	isPos := false
	if num < 0 {
		num = -num
		isPos = true
	}

	if num == 0 {
		return "0"
	}

	for num > 0 {
		result = strconv.Itoa(num%7) + result
		num = num / 7
	}

	if isPos {
		return "-" + result
	}

	return result
}
