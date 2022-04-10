package main

import "fmt"

func main() {
	fmt.Println(checkAlmostEquivalent("aaaa", "bccb"))
}

func checkAlmostEquivalent(w1 string, w2 string) bool {
	set := make([]bool, 26)
	mp1 := make([]int, 26)
	mp2 := make([]int, 26)

	for i := 0; i < len(w1); i++ {
		set[w1[i]-'a'] = true
		mp1[w1[i]-'a']++
	}
	for i := 0; i < len(w2); i++ {
		set[w2[i]-'a'] = true
		mp2[w2[i]-'a']++
	}

	for k := range set {
		if t := abs(mp1[k], mp2[k]); t > 3 {
			return false
		}
	}

	return true
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
