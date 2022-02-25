package main

import "fmt"

func main() {
	fmt.Println(findLUSlength("aba", "cdc") == 3)
	fmt.Println(findLUSlength("aaa", "bbb") == 3)
	fmt.Println(findLUSlength("aaa", "aaa") == -1)
}

func findLUSlength(a string, b string) int {
	if a == b {
		return -1
	}
	return max(len(a), len(b))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
