package main

import "fmt"

func main() {
	fmt.Println(minAddToMakeValid("()))(("))
	fmt.Println(minAddToMakeValid("())"))
	fmt.Println(minAddToMakeValid("((("))
}

func minAddToMakeValid(s string) int {
	cr := 0
	cl := 0
	for i := range s {
		if s[i] == '(' {
			cl++
		} else {
			if cl == 0 {
				cr++
			} else {
				cl--
			}
		}
	}

	return cl + cr
}
