package main

func main() {
}

func makeEqual(words []string) bool {
	mp := make([]int, 26)
	for _, w := range words {
		for _, c := range w {
			mp[c-'a']++
		}
	}

	ln := len(words)
	for i := 0; i < len(mp); i++ {
		if mp[i]%ln != 0 {
			return false
		}
	}

	return true
}
