package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(ambiguousCoordinates("(00011)"))
}

func ambiguousCoordinates(s string) []string {
	gen := func(i, j int) []string {
		var r []string
		for d := 1; d <= j-i; d++ {
			a := s[i : d+i]
			b := s[d+i : j]
			if (a[0] != '0' || a == "0") && !strings.HasSuffix(b, "0") {
				if d < j-i {
					r = append(r, fmt.Sprintf("%s.%s", a, b))
				} else {
					r = append(r, fmt.Sprintf("%s%s", a, b))
				}
			}
		}
		return r
	}

	var r []string
	for i := 1; i < len(s)-1; i++ {
		for _, left := range gen(1, i) {
			for _, right := range gen(i, len(s)-1) {
				r = append(r, fmt.Sprintf("(%s, %s)", left, right))
			}
		}
	}

	return r
}
