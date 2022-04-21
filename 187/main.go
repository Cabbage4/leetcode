package main

import "fmt"

func main() {
	fmt.Println(findRepeatedDnaSequences("AAAAAAAAAAA"))
	fmt.Println(findRepeatedDnaSequences("AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"))
}

func findRepeatedDnaSequences(s string) []string {
	r := make([]string, 0)
	if len(s) < 10 {
		return r
	}

	mp := make(map[string]int)
	for i := 0; i < len(s)-9; i++ {
		mp[s[i:i+10]]++
	}

	for k, v := range mp {
		if v > 1 {
			r = append(r, k)
		}
	}

	return r
}
