package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findLongestWord("abpcplea", []string{"ale", "apple", "monkey", "plea", "abpcplaaa", "abpcllllll", "abccclllpppeeaaaa"}))
	//fmt.Println(findLongestWord("abpcplea", []string{"ale", "apple", "monkey", "plea"}))
}

func findLongestWord(s string, dictionary []string) string {
	sort.Slice(dictionary, func(i, j int) bool {
		if len(dictionary[i]) == len(dictionary[j]) {
			return dictionary[i] < dictionary[j]
		}
		return len(dictionary[j]) < len(dictionary[i])
	})

	mp := make([][26]int, len(s)+1)
	for i := range mp[len(s)] {
		mp[len(s)][i] = -1
	}
	for i := len(s) - 1; i >= 0; i-- {
		mp[i] = mp[i+1]
		mp[i][s[i]-'a'] = i
	}

	for _, d := range dictionary {
		if len(d) > len(s) {
			continue
		}

		mpI := 0
		for i := 0; i < len(d); i++ {
			nmpI := mp[mpI][d[i]-'a']
			if nmpI == -1 {
				break
			}

			mpI = nmpI + 1

			if i == len(d)-1 {
				return d
			}
		}
	}

	return ""
}
