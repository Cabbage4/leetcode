package main

import "fmt"

func main() {
	fmt.Println(letterCasePermutation("3z4"))
}

func letterCasePermutation(s string) []string {
	r := []string{s}
	sbyte := []byte(s)

	var dfs func(int)
	dfs = func(index int) {
		if index == len(s) {
			return
		}

		for i := index; i < len(s); i++ {
			c := sbyte[i]
			if !('a' <= c && c <= 'z' || 'A' <= c && c <= 'Z') {
				continue
			}

			if 'a' <= c && c <= 'z' {
				sbyte[i] = c - 'a' + 'A'
				r = append(r, string(sbyte))
				dfs(i + 1)
				sbyte[i] = c
			}

			if 'A' <= c && c <= 'Z' {
				sbyte[i] = c - 'A' + 'a'
				r = append(r, string(sbyte))
				dfs(i + 1)
				sbyte[i] = c
			}
		}
	}

	dfs(0)
	return r
}
