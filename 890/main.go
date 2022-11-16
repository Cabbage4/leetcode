package main

import "fmt"

func main() {
	fmt.Println(findAndReplacePattern([]string{"deq", "mee", "aqq", "dkd", "ccc"}, "abb"))
}

func findAndReplacePattern(words []string, p string) []string {
	var r []string
	for _, w := range words {
		flag := false
		mp := make([]byte, 26)
		mpR := make([]byte, 26)
		wb := []byte(w)
		for i := range p {
			if v := mp[w[i]-'a']; v != 0 && v != p[i] {
				flag = true
				break
			}
			if v := mpR[p[i]-'a']; v != 0 && v != w[i] {
				flag = true
				break
			}

			mp[w[i]-'a'] = p[i]
			mpR[p[i]-'a'] = w[i]

			wb[i] = p[i]
		}

		if flag || string(wb) != p {
			continue
		}

		r = append(r, w)
	}

	return r
}
