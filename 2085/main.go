package main

func main() {
}

func countWords(words1 []string, words2 []string) int {
	mp := make(map[string]int)

	for _, w := range words1 {
		mp[w]++
	}
	for _, w := range words2 {
		if mp[w] < 2 {
			mp[w]--
		}
	}

	r := 0
	for k := range mp {
		if mp[k] == 0 {
			r++
		}
	}
	return r
}
