package main

import "fmt"

func main() {
	fmt.Println(detectCapitalUse("USA"))
	fmt.Println(detectCapitalUse("FFFFFFFFFFFFFFFFFFFFf"))
	fmt.Println(detectCapitalUse("FlaG"))
}

func detectCapitalUse(word string) bool {
	if len(word) < 1 {
		return true
	}

	if 'a' <= word[0] && 'z' >= word[0] {
		for i := 1; i < len(word); i++ {
			if !('a' <= word[i] && 'z' >= word[i]) {
				return false
			}
		}
		return true
	}

	if len(word) < 2 {
		return true
	}

	if 'a' <= word[1] && 'z' >= word[1] {
		for i := 1; i < len(word); i++ {
			if !('a' <= word[i] && 'z' >= word[i]) {
				return false
			}
		}
		return true
	}

	for i := 1; i < len(word); i++ {
		if !('A' <= word[i] && 'Z' >= word[i]) {
			return false
		}
	}
	return true
}
