package main

func main() {
}

func mostWordsFound(sentences []string) int {
	r := 0
	for _, s := range sentences {
		c := 0
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

			c++
			i = j
		}

		if c > r {
			r = c
		}
	}
	return r
}
