package main

import "fmt"

func main() {
	fmt.Println(countVowelSubstrings("cuaieuouac"))
	fmt.Println(countVowelSubstrings("aeiouu"))
}

func countVowelSubstrings(w string) int {
	mp := make([]bool, 256)
	mp['a'] = true
	mp['e'] = true
	mp['i'] = true
	mp['o'] = true
	mp['u'] = true

	ans := 0
	ln := len(w)
	for c := ln; c >= 5; c-- {
		for l := 0; l < ln-c+1; l++ {
			if !mp[w[l]] {
				continue
			}

			cc := 0
			flag := true
			bl := make([]int, 256)
			for r := l; r < l+c; r++ {
				if !mp[w[r]] {
					flag = false
					break
				}
				bl[w[r]]++
				if bl[w[r]] == 1 {
					cc++
				}
			}

			if !flag {
				continue
			}

			if cc == 5 {
				ans++
			}
		}
	}
	return ans
}
