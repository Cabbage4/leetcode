package main

import "sort"

func main() {

}

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	gi := len(g) - 1
	si := len(s) - 1

	result := 0
	for gi >= 0 && si >= 0 {
		if s[si] >= g[gi] {
			si--
			gi--
			result++
			continue
		}
		gi--
	}

	return result
}
