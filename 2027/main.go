package main

import "fmt"

func main() {
	fmt.Println(minimumMoves("XXOX"))
	fmt.Println(minimumMoves("XXX"))
}

func minimumMoves(s string) int {
	r := 0
	for i := 0; i < len(s); {
		if s[i] == 'O' {
			i++
			continue
		}

		r++
		i = i + 3
	}
	return r
}
