package main

import (
	"strings"
)

func main() {
}

func stringMatching(words []string) []string {
	r := make([]string, 0)

	s := strings.Join(words, " ")
	for i := 0; i < len(words); i++ {
		if strings.Count(s, words[i]) >= 2 {
			r = append(r, words[i])
		}
	}

	return r
}
