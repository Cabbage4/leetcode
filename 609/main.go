package main

import (
	"strings"
)

func main() {
}

func findDuplicate(paths []string) [][]string {
	mp := make(map[string][]string)
	for _, p := range paths {
		info := strings.Split(p, " ")
		dir := info[0]

		for _, f := range info[1:] {
			left := 0
			for left < len(f) && f[left] != '(' {
				left++
			}

			right := left
			for right < len(f) && f[right] != ')' {
				right++
			}

			content := f[left+1 : right]

			mp[content] = append(mp[content], dir+"/"+f[:left])
		}
	}

	var r [][]string
	for _, v := range mp {
		if len(v) <= 1 {
			continue
		}

		r = append(r, v)
	}

	return r
}
