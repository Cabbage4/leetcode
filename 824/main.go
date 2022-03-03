package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(toGoatLatin("I speak Goat Latin"))
}

func toGoatLatin(sentence string) string {
	omp := make([]bool, 256)
	omp['a'] = true
	omp['e'] = true
	omp['i'] = true
	omp['o'] = true
	omp['u'] = true
	omp['A'] = true
	omp['E'] = true
	omp['I'] = true
	omp['O'] = true
	omp['U'] = true

	a := "a"
	r := bytes.NewBuffer(make([]byte, 0, len(sentence)*4))
	for i := 0; i < len(sentence); {
		for i < len(sentence) && !is(sentence[i]) {
			r.WriteByte(sentence[i])
			i++
		}

		j := i + 1
		for j < len(sentence) && is(sentence[j]) {
			j++
		}

		if omp[sentence[i]] {
			r.WriteString(sentence[i:j])
			r.WriteString("ma")
			r.WriteString(a)
		} else {
			r.WriteString(sentence[i+1 : j])
			r.WriteByte(sentence[i])
			r.WriteString("ma")
			r.WriteString(a)
		}

		a += "a"
		i = j
	}

	return r.String()
}

func is(a byte) bool {
	return ('a' <= a && a <= 'z') || ('A' <= a && a <= 'Z')
}
