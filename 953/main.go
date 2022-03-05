package main

import "fmt"

func main() {
	fmt.Println(isAlienSorted([]string{"hello", "leetcode"}, "hlabcdefgijkmnopqrstuvwxyz"))
}

func isAlienSorted(words []string, order string) bool {
	omp := make([]int, 256)
	for i := range order {
		omp[order[i]] = i
	}

	for i := 1; i < len(words); i++ {
		if less(words[i-1], words[i], omp) {
			continue
		}
		return false
	}

	return true
}

func less(a, b string, omp []int) bool {
	ln := len(a)
	if len(a) > len(b) {
		ln = len(b)
	}

	for i := 0; i < ln; i++ {
		if omp[a[i]] == omp[b[i]] {
			continue
		}
		return omp[a[i]] < omp[b[i]]
	}

	return len(a) <= len(b)
}
