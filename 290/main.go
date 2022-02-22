package main

import (
	"strings"
)

func main() {

}

func wordPattern(pattern string, s string) bool {
	mp := make(map[string]string)
	mp2 := make(map[string]string)

	list := strings.Split(s, " ")
	if len(list) != len(pattern) {
		return false
	}

	for i, v := range list {
		p := string(pattern[i])
		mv, has := mp[p]
		if has && mv != v {
			return false
		}
		mp[p] = v

		mv2, has2 := mp2[v]
		if has2 && mv2 != p {
			return false
		}
		mp2[v] = p
	}

	return true
}
