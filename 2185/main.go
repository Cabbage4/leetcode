package main

import "strings"

func main() {
}

func prefixCount(words []string, pref string) int {
	r := 0
	for _, w := range words {
		if strings.HasPrefix(w, pref) {
			r++
		}
	}
	return r
}
