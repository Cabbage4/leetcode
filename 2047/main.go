package main

import "fmt"

func main() {
	fmt.Println(countValidWords("a-!b"))
	fmt.Println(countValidWords("!"))
	fmt.Println(countValidWords("he bought 2 pencils, 3 erasers, and 1  pencil-sharpener."))
	fmt.Println(countValidWords("!this  1-s b8d!"))
}

func countValidWords(s string) int {
	r := 0
	for i := 0; i < len(s); {
		if s[i] == ' ' {
			i++
			continue
		}

		pi := -1
		li := -1
		is := false

		if s[i] == '-' {
			is = false
			li = i
		}
		if s[i] == '.' || s[i] == ',' || s[i] == '!' {
			pi = i
			is = true
		}
		if 'a' <= s[i] && s[i] <= 'z' {
			is = true
		}

		j := i + 1
		for ; j < len(s) && s[j] != ' '; j++ {
			if !is {
				continue
			}

			if '0' <= s[j] && s[j] <= '9' {
				is = false
			} else if s[j] == '-' {
				if li != -1 {
					is = false
					continue
				}

				li = j
			} else if s[j] == '.' || s[j] == ',' || s[j] == '!' {
				if pi != -1 {
					is = false
					continue
				}

				pi = j
			}
		}

		if pi != -1 && pi != j-1 {
			is = false
		}
		if li != -1 && (li == j-1 || li == i) {
			is = false
		}
		if is && pi != -1 && li != -1 && (pi-li == 1 || li-pi == 1) {
			is = false
		}

		if is {
			r++
		}

		i = j
	}
	return r
}
