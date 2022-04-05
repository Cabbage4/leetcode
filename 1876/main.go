package main

import "fmt"

func main() {
	fmt.Println(countGoodSubstrings("aababcabc"))
}

func countGoodSubstrings(s string) int {
	if len(s) < 3 {
		return 0
	}

	r := 0
	mp := make([]int, 256)
	for i := 0; i < 3; i++ {
		mp[s[i]]++
	}

	if mp[s[0]] == 1 && mp[s[1]] == 1 && mp[s[2]] == 1 {
		r++
	}

	for i := 1; i < len(s)-2; i++ {
		mp[s[i-1]]--
		mp[s[i+2]]++
		if mp[s[i]] == 1 && mp[s[i+1]] == 1 && mp[s[i+2]] == 1 {
			r++
		}
	}

	return r
}
