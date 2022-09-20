package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(frequencySort("loveleetcode"))
}

func frequencySort(s string) string {
	var mp [128]int
	for i := 0; i < len(s); i++ {
		mp[s[i]]++
	}

	sList := []byte(s)
	sort.Slice(sList, func(i, j int) bool {
		if mp[sList[i]] == mp[sList[j]] {
			return sList[i] < sList[j]
		}

		return mp[sList[i]] > mp[sList[j]]
	})

	return string(sList)
}
