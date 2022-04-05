package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(mergeAlternately("abc", "pqr"))
}

func mergeAlternately(word1 string, word2 string) string {
	r := bytes.NewBufferString("")
	ln := len(word2)
	if len(word1) < len(word2) {
		ln = len(word1)
	}
	for i := 0; i < ln; i++ {
		r.WriteByte(word1[i])
		r.WriteByte(word2[i])
	}

	if len(word1) < len(word2) {
		r.WriteString(word2[ln:])
	} else if len(word1) > len(word2) {
		r.WriteString(word1[ln:])
	}

	return r.String()
}
