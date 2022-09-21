package main

import "fmt"

func main() {
	fmt.Println(magicalString(7))
	fmt.Println(magicalString(6))
}

func magicalString(n int) int {
	if n <= 3 {
		return 1
	}

	s := "122"
	index := 2
	for len(s) < n {
		back := s[len(s)-1]

		if s[index] == '1' {
			if back == '1' {
				s += "2"
			} else {
				s += "1"
			}
		} else {
			if back == '1' {
				s += "22"
			} else {
				s += "11"
			}
		}

		index++
	}

	var r int
	for i := 0; i < n; i++ {
		if s[i] == '1' {
			r++
		}
	}

	return r
}
