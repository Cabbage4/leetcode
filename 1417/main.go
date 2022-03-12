package main

import "fmt"

func main() {
	fmt.Println(reformat("a0b1c2"))
}

func reformat(s string) string {
	p1 := make([]byte, 0)
	p2 := make([]byte, 0)
	r := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		if '0' <= s[i] && s[i] <= '9' {
			p2 = append(p2, s[i])
		} else {
			p1 = append(p1, s[i])
		}
	}

	if len(p1) < len(p2) {
		p2, p1 = p1, p2
	}

	if len(p1)-len(p2) > 1 {
		return ""
	}

	p1i := 0
	p2i := 0
	for i := 0; i < len(r); i++ {
		if i%2 == 0 {
			r[i] = p1[p1i]
			p1i++
		} else {
			r[i] = p2[p2i]
			p2i++
		}
	}

	return string(r)
}
