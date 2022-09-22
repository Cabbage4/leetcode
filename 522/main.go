package main

import "fmt"

func main() {
	fmt.Println(findLUSlength([]string{"aabbcc", "aabbcc", "c", "e", "aabbcd"}))
	fmt.Println(findLUSlength([]string{"aabbcc", "aabbcc", "cb"}))
	fmt.Println(findLUSlength([]string{"aaa", "aaa", "aa"}))
	fmt.Println(findLUSlength([]string{"aba", "cdc", "eae"}))
}

func findLUSlength(strs []string) int {
	r := -1
	for i := 0; i < len(strs); i++ {
		flag := false
		for j := 0; j < len(strs); j++ {
			if i == j {
				continue
			}

			if isSub(strs[i], strs[j]) {
				flag = true
				break
			}
		}

		if flag {
			continue
		}

		if len(strs[i]) > r {
			r = len(strs[i])
		}
	}
	return r
}

func isSub(s1, s2 string) bool {
	s1i := 0
	s2i := 0
	for s2i < len(s2) {
		if s1i == len(s1) {
			return true
		}

		if s1[s1i] == s2[s2i] {
			s1i++
			s2i++
			continue
		}

		s2i++
	}

	return s1i == len(s1)
}
