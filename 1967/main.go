package main

import "strings"

func main() {
}

func numOfStrings(patterns []string, word string) int {
	r := 0
	for _, p := range patterns {
		if strings.Contains(word, p) {
			r++
		}
	}
	return r
}
