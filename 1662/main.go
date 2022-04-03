package main

import "fmt"

func main() {
	fmt.Println(arrayStringsAreEqual([]string{"ab", "c"}, []string{"a", "bc"}))
	fmt.Println(arrayStringsAreEqual([]string{"ab", "c"}, []string{"a", "bcd"}))
}

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	i, j := 0, 0
	flag := false

	for _, w := range word1 {
		for ci := range w {
			if flag {
				return false
			}

			if word2[i][j] != w[ci] {
				return false
			}

			j++

			if j >= len(word2[i]) {
				i++
				j = 0
			}

			if i >= len(word2) {
				flag = true
			}
		}
	}

	return len(word2) == i
}
