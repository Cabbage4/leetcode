package main

import "fmt"

func main() {
	fmt.Println(truncateSentence("Hello how are you Contestant", 4))
}

func truncateSentence(s string, k int) string {
	c := 0
	for i := 0; i < len(s); {
		if s[i] == ' ' {
			i++
			continue
		}

		j := i + 1
		for ; j < len(s); j++ {
			if s[j] == ' ' {
				break
			}
		}

		c++
		i = j
		if c >= k {
			return s[:i]
		}
	}

	return s
}
