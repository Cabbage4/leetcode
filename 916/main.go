package main

import "fmt"

func main() {
	fmt.Println(wordSubsets([]string{"amazon", "apple", "facebook", "google", "leetcode"}, []string{"e", "oo"}))
	fmt.Println(wordSubsets([]string{"amazon", "apple", "facebook", "google", "leetcode"}, []string{"l", "e"}))
}

func wordSubsets(words1 []string, words2 []string) []string {
	var w2List [26]int
	for _, w := range words2 {
		var tmpList [26]int
		for _, c := range w {
			tmpList[c-'a']++
		}

		for i := range w2List {
			if w2List[i] < tmpList[i] {
				w2List[i] = tmpList[i]
			}
		}
	}

	var r []string
	for i := range words1 {
		var list [26]int
		for _, w := range words1[i] {
			list[w-'a']++
		}

		flag := true
		for k := range w2List {
			if list[k] < w2List[k] {
				flag = false
				break
			}
		}

		if flag {
			r = append(r, words1[i])
		}
	}
	return r
}
