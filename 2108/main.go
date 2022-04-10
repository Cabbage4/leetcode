package main

func main() {
}

func firstPalindrome(words []string) string {
	for _, w := range words {
		if is(w) {
			return w
		}
	}
	return ""
}

func is(s string) bool {
	ln := len(s)
	for i := 0; i < ln/2; i++ {
		if s[i] != s[ln-1-i] {
			return false
		}
	}
	return true
}
