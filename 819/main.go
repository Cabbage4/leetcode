package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(mostCommonWord("Bob hit a ball, the hit BALL flew far after it was hit.", []string{"hit"}))
}

func mostCommonWord(paragraph string, banned []string) string {
	bmp := make(map[string]bool)
	for _, v := range banned {
		bmp[strings.ToLower(v)] = true
	}

	result := ""
	cmp := make(map[string]int)
	for i := 0; i < len(paragraph); {
		if !('a' <= paragraph[i] && paragraph[i] <= 'z') && !('A' <= paragraph[i] && paragraph[i] <= 'Z') {
			i++
			continue
		}

		j := i + 1
		for j < len(paragraph) {
			if ('a' <= paragraph[j] && paragraph[j] <= 'z') || ('A' <= paragraph[j] && paragraph[j] <= 'Z') {
				j++
				continue
			}
			break
		}

		k := strings.ToLower(paragraph[i:j])
		cmp[k]++
		if cmp[k] > cmp[result] && !bmp[k] {
			result = k
		}

		i = j
	}

	return result
}
