package main

import "bytes"

func main() {
}

func mergeAlternately(word1 string, word2 string) string {
	r := bytes.NewBufferString("")

	i1 := 0
	i2 := 0
	for i1 < len(word1) && i2 < len(word2) {
		if word1[i1] < word2[i2] {
			//r.WriteByte()
			i1++
		}
	}

	return r.String()
}
