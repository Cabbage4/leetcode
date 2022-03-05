package main

import "fmt"

func main() {
	fmt.Println(backspaceCompare("ab#c", "ad#c"))
}

func backspaceCompare(s string, t string) bool {
	s = trim(s)
	t = trim(t)
	if len(s) != len(t) {
		return false
	}

	for i := 0; i < len(s); i++ {
		if s[i] != t[i] {
			return false
		}
	}

	return true
}

func trim(s string) string {
	r := ""
	for i := 0; i < len(s); i++ {
		if s[i] == '#' {
			if len(r) > 0 {
				r = r[:len(r)-1]
			}
			continue
		}

		r = r + string(s[i])
	}

	return r
}
