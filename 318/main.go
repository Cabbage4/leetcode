package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxProduct([]string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"}))
}

func maxProduct(words []string) int {
	list := make([]int, len(words))
	for i := range list {
		tmp := make([]int, 26)
		for j := range words[i] {
			tmp[words[i][j]-'a'] = 1
		}

		for j := 0; j < 26; j++ {
			list[i] += tmp[j] << (25 - j)
		}
	}

	r := 0
	for i := range words {
		for j := i + 1; j < len(words); j++ {
			if (list[i] & list[j]) != 0 {
				continue
			}

			t := len(words[i]) * len(words[j])
			if t > r {
				r = t
			}
		}
	}
	return r
}
