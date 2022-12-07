package main

import (
	"fmt"
	"strings"
)

func main() {
	for _, sen := range fullJustify([]string{"This", "is", "an", "example", "of", "text", "justification."}, 16) {
		fmt.Println(sen)
	}

	fmt.Println()
	for _, sen := range fullJustify([]string{"What", "must", "be", "acknowledgment", "shall", "be"}, 16) {
		fmt.Println(sen)
	}

	fmt.Println()
	for _, sen := range fullJustify([]string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"}, 20) {
		fmt.Println(sen)
	}
}

func fullJustify(words []string, maxWidth int) []string {
	var r []string
	for i := 0; i < len(words); {
		j := i + 1
		ln := len(words[i])
		for j < len(words) && j-i+ln+len(words[j]) <= maxWidth {
			ln += len(words[j])
			j++
		}

		if j == len(words) {
			str := strings.Join(words[i:j], " ")
			r = append(r, str+strings.Repeat(" ", maxWidth-len(str)))
			break
		}

		if j-i == 1 {
			r = append(r, words[i]+strings.Repeat(" ", maxWidth-len(words[i])))
			i = j
			continue
		}

		var tr string
		wln := j - i
		sp := (maxWidth - ln) / (wln - 1)
		lsp := (maxWidth - ln) % (wln - 1)

		for k := i; k < j-1; k++ {
			tr += words[k] + strings.Repeat(" ", sp)
			if lsp != 0 {
				tr += " "
				lsp--
			}
		}
		tr += words[j-1]

		r = append(r, tr)

		i = j
	}

	return r
}
