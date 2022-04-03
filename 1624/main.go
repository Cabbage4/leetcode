package main

import "fmt"

func main() {
	fmt.Println(maxLengthBetweenEqualCharacters("aa"))
	fmt.Println(maxLengthBetweenEqualCharacters("abca"))
}

func maxLengthBetweenEqualCharacters(s string) int {
	mp := make([]int, 26)
	for i := 0; i < 26; i++ {
		mp[i] = -1
	}

	r := -1
	for i := 0; i < len(s); i++ {
		if mp[s[i]-'a'] == -1 {
			mp[s[i]-'a'] = i
			continue
		}

		if i-mp[s[i]-'a']-1 > r {
			r = i - mp[s[i]-'a'] - 1
		}
	}

	return r
}
