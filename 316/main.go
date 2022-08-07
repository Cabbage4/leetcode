package main

import (
	"fmt"
)

func main() {
	fmt.Println(removeDuplicateLetters("thesqtitxyetpxloeevdeqifkz"))
	fmt.Println(removeDuplicateLetters("cdadabcc"))
	fmt.Println(removeDuplicateLetters("bcabc"))
}

func removeDuplicateLetters(s string) string {
	lastIndex := make([]int, 26)
	for i := range s {
		lastIndex[s[i]-'a'] = i
	}

	stack := make([]byte, 0)
	book := make([]bool, 26)
	for i := range s {
		if book[s[i]-'a'] {
			continue
		}

		for len(stack) != 0 && stack[len(stack)-1] > s[i] && lastIndex[stack[len(stack)-1]-'a'] > i {
			book[stack[len(stack)-1]-'a'] = false
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, s[i])

		book[s[i]-'a'] = true
	}

	return string(stack)
}
