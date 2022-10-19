package main

import (
	"fmt"
)

func main() {
	r := "abcd"
	i := []int{0, 2}
	s := []string{"a", "cd"}
	t := []string{"eee", "ffff"}
	fmt.Println(findReplaceString(r, i, s, t))
}

func findReplaceString(s string, indices []int, sources []string, targets []string) string {
	match := make([]int, len(s))
	for i := range match {
		match[i] = -1
	}

	for i := range indices {
		ix := indices[i]

		if ix+len(sources[i]) < len(s)+1 && s[ix:ix+len(sources[i])] == sources[i] {
			match[indices[i]] = i
		}
	}

	r := ""
	for i := 0; i < len(s); {
		if match[i] >= 0 {
			r += targets[match[i]]
			i += len(sources[match[i]])
		} else {
			r += string(s[i])
			i++
		}
	}

	return r
}
