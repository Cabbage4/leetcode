package main

import "fmt"

func main() {
	fmt.Println(canBeTypedWords("hello world", "ad"))
}

func canBeTypedWords(s string, brokenLetters string) int {
	mp := make([]bool, 256)
	for _, v := range brokenLetters {
		mp[v] = true
	}

	r := 0
	for i := 0; i < len(s); {
		if s[i] == ' ' {
			i++
			continue
		}

		flag := !mp[s[i]]
		j := i + 1
		for j < len(s) {
			if mp[s[j]] {
				flag = false
			}

			if s[j] != ' ' {
				j++
				continue
			}
			break
		}

		if flag {
			r++
		}

		i = j
	}
	return r
}
