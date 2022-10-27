package main

import "fmt"

func main() {
	fmt.Println(shiftingLetters("xrdofd", []int{70, 41, 71, 72, 73, 84}))
	fmt.Println(shiftingLetters("aaa", []int{1, 2, 3}))
	fmt.Println(shiftingLetters("abc", []int{3, 5, 9}))
}

func shiftingLetters(s string, shifts []int) string {
	sb := make([]byte, len(s))
	sum := 0
	for i := len(s) - 1; i >= 0; i-- {
		sum += shifts[i]
		sb[i] = 'a' + byte((int(s[i])-int('a')+sum)%26)
	}
	return string(sb)
}
