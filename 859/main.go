package main

import "fmt"

func main() {
	fmt.Println(buddyStrings("ab", "ba"))
	fmt.Println(buddyStrings("ab", "ab"))
	fmt.Println(buddyStrings("aa", "aa"))
}

func buddyStrings(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}

	dList := make([][]byte, 0)
	cmp := make([]int, 256)
	hasTwoDef := false

	for i := 0; i < len(s); i++ {
		if s[i] == goal[i] {
			cmp[s[i]]++
			if cmp[s[i]] >= 2 {
				hasTwoDef = true
			}
		} else {
			if len(dList) >= 2 {
				return false
			}
			dList = append(dList, []byte{s[i], goal[i]})
		}
	}

	if len(dList) == 1 {
		return false
	} else if len(dList) == 0 {
		return hasTwoDef
	}

	return dList[0][1] == dList[1][0] && dList[0][0] == dList[1][1]
}
