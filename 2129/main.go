package main

import "fmt"

func main() {
	fmt.Println(capitalizeTitle("first leTTeR of EACH Word"))
}

func capitalizeTitle(title string) string {
	r := []byte(title)

	for i := 0; i < len(title); {
		if title[i] == ' ' {
			i++
			continue
		}

		j := i + 1
		for j < len(title) {
			if title[j] == ' ' {
				break
			}
			j++
		}

		if j-i <= 2 {
			for o := i; o < j; o++ {
				if 'A' <= r[o] && r[o] <= 'Z' {
					r[o] = r[o] - 'A' + 'a'
				}
			}
		} else {
			if 'a' <= r[i] && r[i] <= 'z' {
				r[i] = r[i] - 'a' + 'A'
			}
			for o := i + 1; o < j; o++ {
				if 'A' <= r[o] && r[o] <= 'Z' {
					r[o] = r[o] - 'A' + 'a'
				}
			}
		}

		i = j
	}

	return string(r)
}
