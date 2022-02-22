package main

import (
	"fmt"
)

func main() {
	fmt.Println(lengthOfLastWord("Today is a nice day"))
	fmt.Println(lengthOfLastWord("a"))
	fmt.Println(lengthOfLastWord(" a"))
	fmt.Println(lengthOfLastWord("Hello World"))
	fmt.Println(lengthOfLastWord("   fly me   to   the moon  "))
	fmt.Println(lengthOfLastWord("luffy is still joyboy"))
}

func lengthOfLastWord(s string) int {
	ns := s
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			continue
		}

		ns = s[:i+1]
		break
	}

	res := 0
	l := len(ns)
	isNext := false
	for i := len(ns) - 1; i >= 0; i-- {
		if ns[i] == ' ' {
			if isNext {
				break
			}
			continue
		}
		res = l - i
		isNext = true
	}

	return res
}
