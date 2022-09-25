package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(replaceWords([]string{"cat", "bat", "rat"}, "the cattle was rattled by the battery"))
}

func replaceWords(dictionary []string, sentence string) string {
	mp := make(map[string]bool)
	for _, d := range dictionary {
		mp[d] = true
	}

	r := bytes.NewBuffer(make([]byte, 0, len(sentence)))
	for i := 0; i < len(sentence); {
		if sentence[i] == ' ' {
			r.Write([]byte(" "))
			i++
			continue
		}

		findJ := -1
		j := i + 1
		for j < len(sentence) && sentence[j] != ' ' {
			if findJ == -1 && mp[sentence[i:j]] {
				findJ = j
			}
			j++
		}

		if findJ != -1 {
			r.Write([]byte(sentence[i:findJ]))
		} else {
			r.Write([]byte(sentence[i:j]))
		}

		i = j
	}

	return r.String()
}
