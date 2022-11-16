package main

import "fmt"

func main() {
	fmt.Println(numSpecialEquivGroups([]string{"abcd", "cdab", "cbad", "xyzz", "zzxy", "zzyx"}))
}

func numSpecialEquivGroups(words []string) int {
	mp := make(map[[26 * 2]int]int)
	for _, w := range words {
		var list [26 * 2]int
		for i := range w {
			if i%2 == 0 {
				list[w[i]-'a']++
			} else {
				list[w[i]-'a'+26]++
			}
		}

		mp[list]++
	}
	return len(mp)
}
