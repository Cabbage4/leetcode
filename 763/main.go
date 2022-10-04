package main

import (
	"fmt"
)

func main() {
	fmt.Println(partitionLabels("ababcbacadefegdehijhklij"))
}

func partitionLabels(s string) []int {
	var mp [26]int
	for i := 0; i < len(s); i++ {
		mp[s[i]-'a'] = i
	}

	var r []int
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		if mp[s[i]-'a'] > end {
			end = mp[s[i]-'a']
		}

		if end == i {
			r = append(r, end-start+1)
			start = end + 1
		}
	}
	return r
}
