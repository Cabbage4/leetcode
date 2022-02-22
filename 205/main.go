package main

import "fmt"

func main() {
	fmt.Println(isIsomorphic("badc", "baba"))
}

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	st := make([]bool, 256)
	sc := 0
	for i := range s {
		if !st[s[i]] {
			sc++
		}
		st[s[i]] = true
	}

	tt := make([]bool, 256)
	tc := 0
	for i := range t {
		if !tt[t[i]] {
			tc++
		}
		tt[t[i]] = true
	}

	if tc < sc {
		s, t = t, s
	}

	mp := make([]int, 256)
	for i := 0; i < 256; i++ {
		mp[i] = -1
	}

	for i := 0; i < len(s); i++ {
		if mp[s[i]] != -1 && mp[s[i]] != int(t[i]) {
			return false
		}
		mp[s[i]] = int(t[i])
	}

	return true
}
