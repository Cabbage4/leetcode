package main

import "fmt"

func main() {
	fmt.Println(findAnagrams("aa", "bb"))
	fmt.Println(findAnagrams("cbaebabacd", "abc"))
}

func findAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return nil
	}

	var mp [26]int
	for i := 0; i < len(p); i++ {
		mp[p[i]-'a']++
		mp[s[i]-'a']--
	}

	diff := 0
	for _, v := range mp {
		if v != 0 {
			diff++
		}
	}

	var r []int
	if diff == 0 {
		r = append(r, 0)
	}

	for i := len(p); i < len(s); i++ {
		if mp[s[i-len(p)]-'a'] == 0 {
			diff++
		} else if mp[s[i-len(p)]-'a'] == -1 {
			diff--
		}
		mp[s[i-len(p)]-'a']++

		if mp[s[i]-'a'] == 1 {
			diff--
		} else if mp[s[i]-'a'] == 0 {
			diff++
		}
		mp[s[i]-'a']--

		if diff == 0 {
			r = append(r, i-len(p)+1)
		}
	}

	return r
}
