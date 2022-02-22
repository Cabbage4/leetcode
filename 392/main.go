package main

import "fmt"

func main() {
	//fmt.Println(isSubsequence("abc", "ahbgdc"))
	fmt.Println(isSubsequence("b", "abc"))
}

func isSubsequence(s string, t string) bool {
	if s == "" {
		return true
	}
	if t == "" {
		return false
	}

	si := 0
	ti := 0

	for ti < len(t) {
		for ti < len(t) && t[ti] != s[si] {
			ti++
		}

		if ti < len(t) {
			si++
			ti++

			if si >= len(s) {
				return true
			}

			continue
		}

		return false
	}

	return si == len(s)
}
