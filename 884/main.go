package main

import "fmt"

func main() {
	fmt.Println(uncommonFromSentences("this apple is sweet", "this apple is sour"))
}

func uncommonFromSentences(s1 string, s2 string) []string {
	r := make([]string, 0)
	mp := make(map[string]int)
	gmp(s1, mp)
	gmp(s2, mp)

	for k, v := range mp {
		if v == 1 {
			r = append(r, k)
		}
	}

	return r
}

func gmp(s string, r map[string]int) {
	for i := 0; i < len(s); {
		if s[i] == ' ' {
			i++
			continue
		}

		j := i + 1
		for j < len(s) {
			if s[j] == ' ' {
				break
			}
			j++
		}

		r[s[i:j]]++
		i = j
	}
}
