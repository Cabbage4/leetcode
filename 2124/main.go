package main

import "fmt"

func main() {
	fmt.Println(checkString("aaabbb"))
}

func checkString(s string) bool {
	hb := false
	for i := 0; i < len(s); i++ {
		if hb && s[i] == 'a' {
			return false
		}

		if s[i] == 'b' {
			hb = true
		}
	}

	return true
}
