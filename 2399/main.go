package main

import "fmt"

func main() {
	fmt.Println(checkDistances("abaccb", []int{1, 3, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}))
	fmt.Println(checkDistances("aa", []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}))
}

func checkDistances(s string, distance []int) bool {
	mp := make(map[byte]int)
	for i, v := range distance {
		mp[byte('a'+i)] = v
	}

	mp2 := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		v, ok := mp2[s[i]]
		if !ok {
			mp2[s[i]] = i
			continue
		}

		if i-v-1 != mp[s[i]] {
			return false
		}
	}

	return true
}
