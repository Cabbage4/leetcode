package main

func main() {
}

func isPrefixString(s string, words []string) bool {
	wi, wj := 0, 0
	for i := 0; i < len(s); i++ {
		if wj >= len(words[wi]) {
			wi++
			wj = 0
		}

		if wi >= len(words) {
			return false
		}

		w := words[wi]

		if w[wj] != s[i] {
			return false
		}

		wj++
	}

	return len(words[wi]) == wj
}
