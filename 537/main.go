package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(complexNumberMultiply("78+-76i", "-86+72i"))
	fmt.Println(complexNumberMultiply("1+1i", "1+1i"))
	fmt.Println(complexNumberMultiply("1+-1i", "1+-1i"))
}

func complexNumberMultiply(num1 string, num2 string) string {
	first1, second1 := parse(num1)
	first2, second2 := parse(num2)

	a := first2*first1 - second2*second1
	b := first2*second1 + second2*first1
	return fmt.Sprintf("%d+%di", a, b)
}

func parse(s string) (int, int) {
	i := strings.Index(s, "+")
	a, _ := strconv.Atoi(s[:i])
	b, _ := strconv.Atoi(s[i+1 : len(s)-1])
	return a, b
}
