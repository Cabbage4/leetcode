package main

import "fmt"

func main() {
	fmt.Println(numDifferentIntegers("a1b01c001"))
	fmt.Println(numDifferentIntegers("a123bc34d8ef34"))
}

func numDifferentIntegers(word string) int {
	mp := make(map[string]int)
	for i := 0; i < len(word); {
		if word[i] == '0' && (i+1 == len(word) || !('0' <= word[i+1] && word[i+1] <= '9')) {
			mp["0"]++
			i++
			continue
		}

		if !('1' <= word[i] && word[i] <= '9') {
			i++
			continue
		}

		j := i + 1
		for ; j < len(word); j++ {
			if !('0' <= word[j] && word[j] <= '9') {
				break
			}
		}

		mp[word[i:j]]++
		i = j
	}

	return len(mp)
}
