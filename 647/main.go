package main

func main() {
}

func countSubstrings(s string) int {
	ln := len(s)
	var r int
	for i := 0; i < 2*ln-1; i++ {
		left := i / 2
		right := left + i%2
		for left >= 0 && right < ln && s[left] == s[right] {
			left--
			right++
			r++
		}
	}

	return r
}
