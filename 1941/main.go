package main

import "fmt"

func main() {
	fmt.Println(areOccurrencesEqual("abacbc"))
}

func areOccurrencesEqual(s string) bool {
	mp := make([]int, 26)
	for i := 0; i < len(s); i++ {
		mp[s[i]-'a']++
	}

	key := mp[s[0]-'a']
	for _, v := range mp {
		if v != 0 && key != v {
			return false
		}
	}
	return true
}
