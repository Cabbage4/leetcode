package main

func main() {
}

func countConsistentStrings(allowed string, words []string) int {
	mp := make([]bool, 256)
	for i := range allowed {
		mp[allowed[i]] = true
	}

	r := 0
	for _, w := range words {
		flag := true
		for i := range w {
			if !mp[w[i]] {
				flag = false
				break
			}
		}

		if flag {
			r++
		}
	}
	return r
}
