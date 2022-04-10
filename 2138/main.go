package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(divideString("abcdefghi", 3, 'x'))
	fmt.Println(divideString("abcdefghij", 3, 'x'))
}

func divideString(s string, k int, fill byte) []string {
	r := make([]string, 0)

	for i := 0; i < len(s)-k+1; i = i + k {
		r = append(r, s[i:i+k])
	}

	if d := len(s) % k; d != 0 {
		r = append(r, s[len(s)-d:]+strings.Repeat(string(fill), k-d))
	}

	return r
}
