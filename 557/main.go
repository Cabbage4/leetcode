package main

import "fmt"

func main() {
	fmt.Println(reverseWords("Let's take LeetCode contest"))
}

func reverseWords(s string) string {
	b := []uint8(s)

	for i := 0; i < len(s); {
		for i < len(s) && s[i] == ' ' {
			i++
		}

		if i >= len(s) {
			break
		}

		j := i + 1
		for j < len(s) && s[j] != ' ' {
			j++
		}

		for k := i; k < i+(j-i)/2; k++ {
			b[k], b[j-1-(k-i)] = b[j-1-(k-i)], b[k]
		}

		i = j
	}

	return string(b)
}
