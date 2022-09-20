package main

import "fmt"

func main() {
	fmt.Println(longestContinuousSubstring("abcde"))
	//fmt.Println(longestContinuousSubstring("abacaba"))
}

func longestContinuousSubstring(s string) int {
	r := ""
	for i := 0; i < len(s); {
		j := i + 1
		for len(s) > j && s[j]-s[j-1] == 1 {
			j++
		}

		//fmt.Println(s[i:j])
		if len(s[i:j]) > len(r) {
			r = s[i:j]
		}
		i = j
	}
	return len(r)
}
