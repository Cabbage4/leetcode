package main

import "fmt"

func main() {
	fmt.Println(removePalindromeSub("abbaaaab"))
	fmt.Println(removePalindromeSub("abb"))
	fmt.Println(removePalindromeSub("ababa"))
	fmt.Println(removePalindromeSub("baabb"))
	fmt.Println(removePalindromeSub("aab"))
}

func removePalindromeSub(s string) int {
	if len(s) == 0 {
		return 0
	}
	if is(s) {
		return 1
	}
	return 2
}

func is(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
