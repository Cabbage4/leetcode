package main

import "fmt"

func main() {
	fmt.Println(findSubstring("barfoofoobarthefoobarman", []string{"bar", "foo", "the"}))
	fmt.Println(findSubstring("barfoothefoobarman", []string{"foo", "bar"}))
}

func findSubstring(s string, words []string) []int {
	sl := len(s)
	wl := len(words)
	n := len(words[0])
	var r []int

	if sl < wl*n {
		return r
	}

	for i := 0; i < n && i <= sl-wl*n; i++ {
		mp := make(map[string]int)
		for j := range words {
			mp[words[j]]--
			if mp[words[j]] == 0 {
				delete(mp, words[j])
			}

			mp[s[i+n*j:i+n*(j+1)]]++
			if mp[s[i+n*j:i+n*(j+1)]] == 0 {
				delete(mp, s[i+n*j:i+n*(j+1)])
			}
		}

		for left := i; left < sl-wl*n+1; left += n {
			if left != i {
				mp[s[left-n:left]]--
				if mp[s[left-n:left]] == 0 {
					delete(mp, s[left-n:left])
				}

				mp[s[left+n*(wl-1):left+n*wl]]++
				if mp[s[left+n*(wl-1):left+n*wl]] == 0 {
					delete(mp, s[left+n*(wl-1):left+n*wl])
				}
			}

			if len(mp) == 0 {
				r = append(r, left)
			}
		}
	}
	return r
}
