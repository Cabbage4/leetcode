package main

import "fmt"

func main() {
	fmt.Println(longestNiceSubstring("YazaAay"))
}

func longestNiceSubstring(s string) string {
	r := ""

	for i := 0; i < len(s); i++ {
		mp := make([]int, 256)
		mp[s[i]]++

		flag := false
		for j := i + 1; j < len(s); j++ {
			mp[s[j]]++
			if mp[s[j]] > 1 {
				if flag && j-i+1 > len(r) {
					r = s[i : j+1]
				}
				continue
			}

			if flag {
				flag = false
				continue
			}

			k := s[j]
			if 'A' <= s[j] && s[j] <= 'Z' {
				k = s[j] - 'A' + 'a'
			}

			if !((mp[k] == 0 && mp[k-'a'+'A'] == 0) || (mp[k] > 0 && mp[k-'a'+'A'] > 0)) {
				continue
			}

			flag = true
			for i := 'a'; i <= 'z'; i++ {
				if (mp[i] == 0 && mp[i-'a'+'A'] == 0) || (mp[i] > 0 && mp[i-'a'+'A'] > 0) {
					continue
				}
				flag = false
				break
			}

			if flag && j-i+1 > len(r) {
				r = s[i : j+1]
			}
		}
	}

	return r
}
