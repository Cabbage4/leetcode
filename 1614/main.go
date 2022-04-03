package main

import "fmt"

func main() {
	fmt.Println(maxDepth("(1+(2*3)+((8)/4))+1"))
}

func maxDepth(s string) int {
	if s == "" {
		return 0
	}

	r := 0
	c := 0
	for i := 0; i < len(s); i++ {
		if s[i] != '(' && s[i] != ')' {
			continue
		}

		if s[i] == '(' {
			c++
			if c > r {
				r = c
			}
			continue
		}

		c--
	}

	return r
}
