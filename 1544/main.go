package main

import (
	"fmt"
)

func main() {
	fmt.Println(makeGood("abBAcC"))
	fmt.Println(makeGood("leEeetcode"))
}

func makeGood(s string) string {
	stack := make([]byte, 0)
	for i := range s {
		if len(stack) != 0 && abs(stack[len(stack)-1], s[i]) == 32 {
			stack = stack[:len(stack)-1]
			continue
		}
		stack = append(stack, s[i])
	}

	return string(stack)
}

func abs(a, b byte) byte {
	if a > b {
		return a - b
	}
	return b - a
}
