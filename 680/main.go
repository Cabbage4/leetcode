package main

import "fmt"

func main() {
	fmt.Println(validPalindrome("deeee"))
	fmt.Println(validPalindrome("aba"))
	fmt.Println(validPalindrome("abca"))
	fmt.Println(validPalindrome("abc"))
}

func validPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return is(s[i+1:len(s)-i]) || is(s[i:len(s)-i-1])
		}
	}

	return true
}

func is(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}

	return true
}
