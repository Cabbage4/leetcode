package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(maxRepeating("ababc", "ab"))
}

func maxRepeating(sequence string, word string) int {
	mc := len(sequence) / len(word)
	for i := mc; i >= 1; i-- {
		if strings.Contains(sequence, strings.Repeat(word, i)) {
			return i
		}
	}
	return 0
}
