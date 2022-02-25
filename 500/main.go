package main

import "fmt"

func main() {
	fmt.Println(findWords([]string{"Hello", "Alaska", "Dad", "Peace"}))
	fmt.Println(findWords([]string{"omk"}))
}

var oneMp = toMp("qwertyuiop")
var twoMp = toMp("asdfghjkl")
var threeMp = toMp("zxcvbnm")

func findWords(words []string) []string {
	result := make([]string, 0)

	for _, word := range words {
		fs := word[0]

		var mp map[uint8]bool
		if oneMp[fs] {
			mp = oneMp
		} else if twoMp[fs] {
			mp = twoMp
		} else {
			mp = threeMp
		}

		flag := true
		for i := 0; i < len(word); i++ {
			if !mp[word[i]] {
				flag = false
				break
			}
		}

		if flag {
			result = append(result, word)
		}
	}

	return result
}

func toMp(str string) map[uint8]bool {
	mp := make(map[uint8]bool)
	for _, v := range str {
		mp[uint8(v)] = true
		mp[uint8('A'+v-'a')] = true
	}

	return mp
}
