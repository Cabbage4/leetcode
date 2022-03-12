package main

import "fmt"

func main() {
	fmt.Println(freqAlphabets("10#11#12"))
}

func freqAlphabets(s string) string {
	mp := make(map[string]string)
	for i := 0; i < 9; i++ {
		mp[string('1'+i)] = string('a' + i)
	}
	for i := 10; i <= 26; i++ {
		mp[fmt.Sprintf("%d#", i)] = string('a' + i - 1)
	}

	r := ""

	for i := 0; i < len(s); {
		if i+2 < len(s) && s[i+2] == '#' {
			r += mp[s[i:i+3]]
			i = i + 3
			continue
		}
		r += mp[string(s[i])]
		i++
	}

	return r
}
