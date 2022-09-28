package main

import (
	"bytes"
	"strings"
)

func main() {
}

func repeatedStringMatch(a string, b string) int {
	buffer := bytes.NewBuffer([]byte{})
	max := 2*len(a) + len(b)

	var r int
	for buffer.Len() < max {
		buffer.WriteString(a)
		r++
		if strings.Contains(buffer.String(), b) {
			return r
		}
	}
	return -1
}
