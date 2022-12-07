package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(appendCharacters("aa", "aaa"))
	fmt.Println(appendCharacters("coaching", "coding"))
	fmt.Println(appendCharacters("abcde", "a"))
	fmt.Println(appendCharacters("z", "abcde"))
}

func appendCharacters(s string, t string) int {
	mp := make(map[byte][]int)
	for i := range s {
		mp[s[i]] = append(mp[s[i]], i)
	}

	var r int
	si := 0
	for i := range t {
		list, ok := mp[t[i]]
		if !ok {
			r += len(t) - i
			break
		}

		nsi := sort.SearchInts(list, si)
		if nsi == len(list) {
			r += len(t) - i
			break
		}

		si = list[nsi] + 1
	}

	return r
}
