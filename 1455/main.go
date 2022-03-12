package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isPrefixOfWord("this problem is an easy problem", "pro"))
}

func isPrefixOfWord(s string, w string) int {
	c := 1
	for i := 0; i < len(s); {
		if s[i] == ' ' {
			i++
			continue
		}

		j := i + 1
		for j < len(s) && s[j] != ' ' {
			j++
		}

		if strings.HasPrefix(s[i:j], w) {
			return c
		}

		i = j
		c++
	}

	return -1
}
